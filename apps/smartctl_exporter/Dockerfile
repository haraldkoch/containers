FROM ghcr.io/haraldkoch/alpine:rolling@sha256:7e4f8a44a49759ffb2a5e296fceb36f3ad186e6fc1b9f2d42a930daedc666959
ARG TARGETPLATFORM
ARG TARGETARCH
ARG TARGETOS
ARG VERSION
ARG CHANNEL

RUN apk add smartmontools
RUN wget "https://github.com/prometheus-community/smartctl_exporter/releases/download/v$VERSION/smartctl_exporter-$VERSION.$TARGETOS-$TARGETARCH.tar.gz" \
  && tar xvzf smartctl_exporter-$VERSION.$TARGETOS-$TARGETARCH.tar.gz \
  && ls -l smartctl_exporter-$VERSION.$TARGETOS-$TARGETARCH.tar.gz \
  && mv smartctl_exporter-$VERSION.$TARGETOS-$TARGETARCH/smartctl_exporter /bin/smartctl_exporter \
  && chmod +x /bin/smartctl_exporter \
  && rm -rf smartctl_exporter*

EXPOSE 9633
USER nobody
CMD [ "/bin/smartctl_exporter" ]

LABEL org.opencontainers.image.source="https://github.com/prometheus-community/smartctl_exporter"
