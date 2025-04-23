ARG VERSION
FROM ghcr.io/actions/actions-runner:${VERSION}
ARG TARGETARCH

ENV DEBCONF_NONINTERACTIVE_SEEN=true \
    DEBIAN_FRONTEND="noninteractive" \
    APT_KEY_DONT_WARN_ON_DANGEROUS_USAGE=DontWarn

USER root

RUN \
    apt-get -qq update \
    && \
    apt-get -qq install -y --no-install-recommends --no-install-suggests \
        ca-certificates \
        curl \
        gcc \
        git \
        jo \
        jq \
        moreutils \
        unrar \
        unzip \
        wget \
        zip \
        zstd \
    && rm -rf /var/lib/apt/lists/*

COPY --from=ghcr.io/fluxcd/flux-cli:v2.5.1 /usr/local/bin/flux /usr/local/bin/flux
COPY --from=registry.k8s.io/kubectl:v1.33.0 /bin/kubectl /usr/local/bin/kubectl
COPY --from=docker.io/mikefarah/yq:4.45.1 /usr/bin/yq /usr/local/bin/yq

USER runner
