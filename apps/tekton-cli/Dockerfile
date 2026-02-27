# syntax=docker/dockerfile:1

FROM docker.io/library/alpine:3.23
ARG TARGETARCH
ARG TKNARCH=${TARGETARCH/amd64/x86_64}
ARG VERSION

USER root

RUN \
    apk add --no-cache \
        bash \
        ca-certificates \
        curl \
        tzdata \
    && \
    mkdir -p /app/bin \
    && curl -fsSL "https://github.com/tektoncd/cli/releases/download/${VERSION}/tkn_${VERSION#*v}_Linux_${TKNARCH}.tar.gz" \
      | tar xzf - -C /app/bin \
    && chown -R root:root /app \
    && chmod -R 755 /app \
    && rm -rf /tmp/*


USER nobody:nogroup
ENTRYPOINT ["/app/bin/tkn"]
