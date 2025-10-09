FROM docker.io/library/alpine:3.22.2

ARG TARGETARCH
ARG VERSION

#hadolint ignore=DL3018
RUN \
    apk add --no-cache \
        bash \
        ca-certificates \
        coreutils \
        curl \
        openssl

#hadolint ignore=DL3018,DL3059
RUN \
    curl -fsSL -o /usr/local/bin/kubectl "https://dl.k8s.io/release/${VERSION}/bin/linux/${TARGETARCH}/kubectl" \
    && chmod +x /usr/local/bin/kubectl

USER nobody
ENTRYPOINT ["/usr/local/bin/kubectl"]
