FROM ghcr.io/haraldkoch/alpine:rolling@sha256:68be29ba0fff59cebd8b36e62a50b769aaecb4706103ce3237213c3312617f28

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
