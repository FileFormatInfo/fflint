#syntax=docker/dockerfile-upstream:master-experimental
FROM golang:1.15-buster AS builder
LABEL org.opencontainers.image.source https://github.com/fileformat/badger

# Install upx to compress the compiled action
RUN apt-get update \
    && apt-get -y install upx

ENV CGO_ENABLED=0

# Copy all the files from the host into the container
COPY . .

# build a standalone binary
RUN GOPATH= go build \
    -a \
    -trimpath \
    -ldflags "-s -w -extldflags '-static'" \
    -installsuffix cgo \
    -tags netgo \
    -o /bin/badger-gha \
    .

# Strip any symbols - this is not a library
RUN strip /bin/badger-gha

# Compress the compiled action
RUN upx -q -9 /bin/badger-gha

#
# need a shell to expand parameters
#
FROM debian:buster

COPY ./.github/actions/entrypoint.sh /bin/entrypoint.sh
COPY --from=builder /bin/badger-gha /bin/badger-gha

# Specify the container's entrypoint as the action
ENTRYPOINT ["/bin/entrypoint.sh"]
