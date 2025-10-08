# syntax=docker/dockerfile:1

FROM docker.io/library/alpine:3.22.2
ARG TARGETARCH
ARG TARGETOS
ARG VERSION

RUN \
    apk add --no-cache \
        bash \
        ca-certificates \
        catatonit

RUN apk add --no-cache bash catatonit smartmontools \
    && wget "https://github.com/prometheus-community/smartctl_exporter/releases/download/v$VERSION/smartctl_exporter-$VERSION.$TARGETOS-$TARGETARCH.tar.gz" \
    && tar xvzf smartctl_exporter-$VERSION.$TARGETOS-$TARGETARCH.tar.gz \
    && ls -l smartctl_exporter-$VERSION.$TARGETOS-$TARGETARCH.tar.gz \
    && mv smartctl_exporter-$VERSION.$TARGETOS-$TARGETARCH/smartctl_exporter /bin/smartctl_exporter \
    && chmod +x /bin/smartctl_exporter \
    && rm -rf smartctl_exporter*

EXPOSE 9633
ENTRYPOINT ["/usr/bin/catatonit", "--", "/bin/smartctl_exporter"]
