FROM ghcr.io/haraldkoch/alpine:rolling@sha256:dcb1622b0d71ea9cdf720ec0eb54dafeb5f644484fada446966c5c1261b6d95c

ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

#hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates postgresql16-client~"${VERSION}"

USER kah

COPY ./apps/postgres-init/entrypoint.sh /entrypoint.sh

CMD ["/entrypoint.sh"]

LABEL org.opencontainers.image.source="https://github.com/postgres/postgres"
