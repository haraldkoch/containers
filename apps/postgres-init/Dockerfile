FROM ghcr.io/haraldkoch/alpine:rolling@sha256:38fdc5c23295f389bea86278d1888b7f19a13b6875e9c112a557f3099a3dd1c0

ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

#hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates postgresql15-client~"${VERSION}"

USER kah

COPY ./apps/postgres-init/entrypoint.sh /entrypoint.sh

CMD ["/entrypoint.sh"]

LABEL org.opencontainers.image.source="https://github.com/postgres/postgres"
