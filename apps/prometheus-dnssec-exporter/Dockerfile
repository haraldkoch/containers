# syntax=docker/dockerfile:1

FROM docker.io/library/golang:1.25-alpine AS builder
ARG VERSION
ENV CGO_ENABLED=0
RUN go install -ldflags="-s -w" "github.com/haraldkoch/prometheus-dnssec-exporter@v${VERSION}"

FROM docker.io/library/alpine:3.23.3
RUN \
    apk add --no-cache \
        bash \
        ca-certificates \
        catatonit

COPY --from=builder /go/bin/prometheus-dnssec-exporter /app/prometheus-dnssec-exporter

# this hack is required for testing; the container will not start without a configuration file.
COPY ./resources/dnssec-checks /etc/dnssec-checks

EXPOSE 9204
ENTRYPOINT ["/usr/bin/catatonit", "--", "/app/prometheus-dnssec-exporter"]
