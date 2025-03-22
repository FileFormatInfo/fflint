# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS builder
RUN apk update && \
    apk upgrade && \
    apk --no-cache add git

RUN mkdir /build
ADD . /build/
WORKDIR /build
ARG COMMIT
ARG LASTMOD
ARG VERSION
RUN echo "INFO: building for $COMMIT on $LASTMOD"
RUN \
    CGO_ENABLED=0 GOOS=linux go build \
    -a \
    -installsuffix cgo \
    -ldflags "-X internal.server.COMMIT=$COMMIT -X internal.server.LASTMOD=$LASTMOD -extldflags '-static'" \
    -o fflint-server cmd/server/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/fflint-server /app/
WORKDIR /app
ENV PORT 4000
ENTRYPOINT ["./fflint-server"]
