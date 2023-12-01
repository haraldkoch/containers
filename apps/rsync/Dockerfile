FROM ghcr.io/haraldkoch/alpine:rolling@sha256:0fd96a8419acc589d13de7931c438683de3b3903f006e700399da9779d7612d8
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
