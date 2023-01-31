FROM ghcr.io/haraldkoch/alpine:rolling@sha256:711ec1ad15c9ee11f695be6577f49728ab89360b7b57b2a2c9025cd5b4e64c0b

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
