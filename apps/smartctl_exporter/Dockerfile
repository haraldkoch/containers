FROM docker.io/library/alpine:3.19
ARG TARGETPLATFORM
ARG TARGETARCH
ARG TARGETOS
ARG VERSION
ARG CHANNEL

RUN apk add --no-cache smartmontools \
    && wget "https://github.com/prometheus-community/smartctl_exporter/releases/download/v$VERSION/smartctl_exporter-$VERSION.$TARGETOS-$TARGETARCH.tar.gz" \
    && tar xvzf smartctl_exporter-$VERSION.$TARGETOS-$TARGETARCH.tar.gz \
    && ls -l smartctl_exporter-$VERSION.$TARGETOS-$TARGETARCH.tar.gz \
    && mv smartctl_exporter-$VERSION.$TARGETOS-$TARGETARCH/smartctl_exporter /bin/smartctl_exporter \
    && chmod +x /bin/smartctl_exporter \
    && rm -rf smartctl_exporter*

EXPOSE 9633
ENTRYPOINT [ "/bin/smartctl_exporter" ]

LABEL org.opencontainers.image.source="https://github.com/prometheus-community/smartctl_exporter"
