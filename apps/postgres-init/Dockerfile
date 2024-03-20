FROM ghcr.io/haraldkoch/alpine:rolling@sha256:3be4b72d7cb5e21db7fe2de3e76dd36e9d1e805fc0ee0c162121985aa45070aa

ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

#hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates postgresql16-client~"${VERSION}"

USER kah

COPY ./apps/postgres-init/entrypoint.sh /entrypoint.sh

CMD ["/entrypoint.sh"]

LABEL org.opencontainers.image.source="https://github.com/postgres/postgres"
