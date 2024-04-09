FROM ghcr.io/haraldkoch/alpine:rolling@sha256:e0575093e874528ec64354c0a8e8a84d757699d01fa6aa05c1d72912d2e1f685

ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

#hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates postgresql16-client~"${VERSION}"

USER kah

COPY ./apps/postgres-init/entrypoint.sh /entrypoint.sh

CMD ["/entrypoint.sh"]

LABEL org.opencontainers.image.source="https://github.com/postgres/postgres"
