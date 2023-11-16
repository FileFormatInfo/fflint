# syntax=docker/dockerfile:1
FROM golang:1.18-alpine as builder
RUN apk update && \
    apk upgrade && \
    apk --no-cache add git
RUN mkdir /build
ADD . /build/
WORKDIR /build
ARG COMMIT
ARG LASTMOD
ARG VERSION
ARG BUILTBY
RUN echo "INFO: building for $COMMIT on $LASTMOD"
RUN \
    CGO_ENABLED=0 GOOS=linux go build \
    -a \
    -installsuffix cgo \
    -ldflags "-X main.commit=$COMMIT -X main.date=$LASTMOD -X main.version=$VERSION -X main.builtBy=$BUILTBY -extldflags '-static'" \
    -o fflint-online cmd/online/main.go

RUN \
    CGO_ENABLED=0 GOOS=linux go build \
    -a \
    -installsuffix cgo \
    -ldflags "-X main.commit=$COMMIT -X main.date=$LASTMOD -X main.version=$VERSION -X main.builtBy=$BUILTBY -extldflags '-static'" \
    -o fflint cmd/fflint/main.go

FROM scratch
COPY --from=builder /build/fflint /app/
COPY --from=builder /build/fflint-online /app/
WORKDIR /app
ENV PORT 4000
ENTRYPOINT ["./fflint-online"]
