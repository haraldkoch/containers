FROM docker.io/library/alpine:3.24

ARG TARGETARCH
ARG VERSION

RUN \
    apk add --no-cache \
        bash \
        ca-certificates \
        coreutils \
        curl \
        openssl

RUN \
    curl -fsSL -o /usr/local/bin/kubectl "https://dl.k8s.io/release/${VERSION}/bin/linux/${TARGETARCH}/kubectl" \
    && chmod +x /usr/local/bin/kubectl

USER nobody
ENTRYPOINT ["/usr/local/bin/kubectl"]
