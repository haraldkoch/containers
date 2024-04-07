FROM ghcr.io/haraldkoch/alpine:rolling@sha256:b5cbe5d3df61f6f916b3fac01e73677811aa8236d3962d8c6078e98a9772f555
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
