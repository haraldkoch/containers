FROM ghcr.io/haraldkoch/alpine:rolling@sha256:e0575093e874528ec64354c0a8e8a84d757699d01fa6aa05c1d72912d2e1f685
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
ENTRYPOINT [ "/bin/smartctl_exporter" ]

LABEL org.opencontainers.image.source="https://github.com/prometheus-community/smartctl_exporter"
