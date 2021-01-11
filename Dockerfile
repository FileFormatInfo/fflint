#syntax=docker/dockerfile-upstream:master-experimental
FROM golang:1.15-buster AS builder
LABEL org.opencontainers.image.source https://github.com/fileformat/badger

# Install upx to compress the compiled action
RUN apt-get update \
    && apt-get -y install upx

WORKDIR /build

# Copy all the files from the host into the container
COPY . .

# build a standalone binary
RUN CGO_ENABLED=0 GOPATH= .github/actions/build.sh

# Strip any symbols - this is not a library
#RUN strip .github/actions/badger-gha

# Compress the compiled action
RUN upx -q -9 .github/actions/badger-gha

#
# need a shell to expand parameters
#
FROM debian:buster

COPY .github/actions/entrypoint.sh /bin/entrypoint.sh
COPY --from=builder /build/.github/actions/badger-gha /bin/badger-gha

# Specify the container's entrypoint as the action
ENTRYPOINT ["/bin/entrypoint.sh"]
