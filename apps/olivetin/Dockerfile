# syntax=docker/dockerfile:1

ARG VERSION
FROM docker.io/jamesread/olivetin:${VERSION}

ENV \
    CRYPTOGRAPHY_DONT_BUILD_RUST=1 \
    PIP_BREAK_SYSTEM_PACKAGES=1 \
    PIP_DISABLE_PIP_VERSION_CHECK=1 \
    PIP_NO_CACHE_DIR=1 \
    PIP_ROOT_USER_ACTION=ignore \
    PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    UV_NO_CACHE=true \
    UV_SYSTEM_PYTHON=true

USER root

ARG HTTPIE_VERSION

RUN \
    microdnf install -y --nodocs --noplugins --setopt=keepcache=0 --setopt=install_weak_deps=0 \
      python3-pip \
    && pip install uv \
    && mkdir -p /app \
    && \
    uv pip install --system 'urllib3<2' wheel httpie=="${HTTPIE_VERSION}" \
    && chown -R root:root /app && chmod -R 755 /app \
    && pip uninstall --yes uv \
    && microdnf clean all

USER olivetin
