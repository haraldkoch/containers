ARG VERSION
FROM public.ecr.aws/docker/library/alpine:3.17 as builder
ARG VERSION
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT=""
ARG TARGETPLATFORM
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH} \
    GOARM=${TARGETVARIANT} \
    GOPATH=/go
ENV PATH $GOPATH/bin:$PATH
#hadolint ignore=DL3018
RUN \
    apk add --no-cache git go upx \
    && go install -ldflags="-s -w" "github.com/haraldkoch/prometheus-dnssec-exporter@v${VERSION}" \
    && upx --ultra-brute /go/bin/prometheus-dnssec-exporter

FROM ghcr.io/haraldkoch/alpine:rolling@sha256:ef3c36e19bf0277aa75ce05d61b2f6eec1661dc7e9c99a45fe5ca276757077c4

ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

COPY --from=builder /go/bin/prometheus-dnssec-exporter /app/prometheus-dnssec-exporter

# this hack is required for testing; the container will not start without a configuration file.
COPY ./apps/prometheus-dnssec-exporter/dnssec-checks /etc/dnssec-checks

#hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates
#hadolint ignore=DL3059
RUN \
    chown -R root:root /app \
    && chmod -R 755 /app \
    && rm -rf /tmp/*

EXPOSE 9204
USER kah
CMD ["/app/prometheus-dnssec-exporter"]

LABEL org.opencontainers.image.source="https://github.com/haraldkoch/prometheus-dnssec-exporter"
