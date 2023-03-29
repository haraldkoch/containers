FROM ghcr.io/haraldkoch/alpine:rolling@sha256:f8b4a0dd65db5971a38ebfc5ec31f1132269b8bde43870d08477c0f39ee5a01d

ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

#hadolint ignore=DL3018
RUN apk add --no-cache openssh rsync~"${VERSION}"
#hadolint ignore=DL3059
RUN chown -R root:root /usr/bin/rsync \
    && chmod -R 755 /usr/bin/rsync \
    && ln -s /usr/bin/rsync /app/rsync \
    && rm -rf /tmp/*

CMD ["/app/rsync"]

LABEL org.opencontainers.image.source="https://github.com/WayneD/rsync"
