ARG VERSION
FROM docker.io/library/alpine:3.19 as builder
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
    && go install -ldflags="-s -w" "github.com/haraldkoch/openweathermap-exporter@v${VERSION}" \
    && upx /go/bin/openweathermap-exporter

FROM ghcr.io/haraldkoch/alpine:rolling@sha256:dcb1622b0d71ea9cdf720ec0eb54dafeb5f644484fada446966c5c1261b6d95c

ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

COPY --from=builder /go/bin/openweathermap-exporter /app/openweathermap-exporter

#hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates
#hadolint ignore=DL3059
RUN \
    chown -R root:root /app \
    && chmod -R 755 /app \
    && rm -rf /tmp/*

EXPOSE 2112
USER kah
CMD ["/app/openweathermap-exporter"]

LABEL org.opencontainers.image.source="https://github.com/haraldkoch/openweathermap-exporter"
