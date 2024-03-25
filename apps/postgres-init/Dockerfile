FROM ghcr.io/haraldkoch/alpine:rolling@sha256:b5cbe5d3df61f6f916b3fac01e73677811aa8236d3962d8c6078e98a9772f555

ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

#hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates postgresql16-client~"${VERSION}"

USER kah

COPY ./apps/postgres-init/entrypoint.sh /entrypoint.sh

CMD ["/entrypoint.sh"]

LABEL org.opencontainers.image.source="https://github.com/postgres/postgres"
