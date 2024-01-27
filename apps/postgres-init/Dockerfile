FROM ghcr.io/haraldkoch/alpine:rolling@sha256:7e4f8a44a49759ffb2a5e296fceb36f3ad186e6fc1b9f2d42a930daedc666959

ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

#hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates postgresql16-client~"${VERSION}"

USER kah

COPY ./apps/postgres-init/entrypoint.sh /entrypoint.sh

CMD ["/entrypoint.sh"]

LABEL org.opencontainers.image.source="https://github.com/postgres/postgres"
