FROM ghcr.io/haraldkoch/alpine:rolling@sha256:e2ea02abb6d94a68a26a26addb79811cad083f998371a0fd6169e5603d9d1409

ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

ENV POSTGRES_SUPER_USER="postgres" \
    POSTGRES_RESET="false" \
    POSTGRES_USER_RESET="false" \
    POSTGRES_RESET_CONFIRM="NO"

#hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates postgresql15-client~"${VERSION}"

USER kah

COPY ./apps/postgres-initdb/entrypoint.sh /entrypoint.sh

CMD ["/entrypoint.sh"]

LABEL org.opencontainers.image.source="https://github.com/postgres/postgres"
