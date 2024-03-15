FROM ghcr.io/haraldkoch/alpine:rolling@sha256:58688e0e653bdf3aa6d854a1849545ba6a8b484ebe268001e8fed4b8abff2ff8

ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

#hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates postgresql16-client~"${VERSION}"

USER kah

COPY ./apps/postgres-init/entrypoint.sh /entrypoint.sh

CMD ["/entrypoint.sh"]

LABEL org.opencontainers.image.source="https://github.com/postgres/postgres"
