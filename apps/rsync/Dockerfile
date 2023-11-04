FROM ghcr.io/haraldkoch/alpine:rolling@sha256:38fdc5c23295f389bea86278d1888b7f19a13b6875e9c112a557f3099a3dd1c0
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
