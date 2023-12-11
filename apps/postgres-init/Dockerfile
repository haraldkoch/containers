FROM ghcr.io/haraldkoch/alpine:rolling@sha256:2c9b74b4e800bd130cff5919dfba9e1ceaf33f23e0a7a66833e01ef7d377e4a7

ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

#hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates postgresql16-client~"${VERSION}"

USER kah

COPY ./apps/postgres-init/entrypoint.sh /entrypoint.sh

CMD ["/entrypoint.sh"]

LABEL org.opencontainers.image.source="https://github.com/postgres/postgres"
