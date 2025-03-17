FROM docker.io/library/alpine:3.21

ARG TARGETPLATFORM
ARG VERSION

USER root
WORKDIR /app

#hadolint ignore=DL3018
RUN \
    apk add --no-cache \
        bash \
        ca-certificates \
        catatonit \
        coreutils \
        curl \
        openssl

#hadolint ignore=DL3018,DL3059
RUN \
    curl -fsSL -o /usr/local/bin/kubectl "https://dl.k8s.io/release/${VERSION}/bin/linux/${TARGETARCH}/kubectl" \
    && chmod +x /usr/local/bin/kubectl

ENTRYPOINT ["/usr/bin/catatonit", "--", "/usr/local/bin/kubectl"]

LABEL org.opencontainers.image.source="https://github.com/kubernetes/kubectl"
