FROM docker.io/library/alpine:3.21
ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

#hadolint ignore=DL3018
RUN apk add --no-cache bash ca-certificates catatonit openssh rsync~"${VERSION}"

ENTRYPOINT ["/usr/bin/catatonit", "--", "/usr/bin/rsync"]

LABEL org.opencontainers.image.source="https://github.com/WayneD/rsync"
