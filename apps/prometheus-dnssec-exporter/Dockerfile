ARG VERSION
FROM docker.io/library/golang:1.24-alpine AS dnssec-exporter
ARG VERSION
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT=""
ARG TARGETPLATFORM
ENV CGO_ENABLED=0 \
    GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH} \
    GOARM=${TARGETVARIANT}
RUN go install -ldflags="-s -w" "github.com/haraldkoch/prometheus-dnssec-exporter@v${VERSION}"

FROM docker.io/library/alpine:3.21.3
ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

RUN \
    apk add --no-cache \
        bash \
        ca-certificates \
        catatonit

COPY --from=dnssec-exporter /go/bin/prometheus-dnssec-exporter /app/prometheus-dnssec-exporter

# this hack is required for testing; the container will not start without a configuration file.
COPY ./apps/prometheus-dnssec-exporter/dnssec-checks /etc/dnssec-checks

EXPOSE 9204
ENTRYPOINT ["/usr/bin/catatonit", "--", "/app/prometheus-dnssec-exporter"]

LABEL org.opencontainers.image.source="https://github.com/haraldkoch/prometheus-dnssec-exporter"
