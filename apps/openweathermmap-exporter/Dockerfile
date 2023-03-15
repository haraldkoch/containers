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
    && go install -ldflags="-s -w" "github.com/haraldkoch/openweathermap-exporter@v${VERSION}" \
    && upx --ultra-brute /go/bin/openweathermap-exporter

FROM ghcr.io/haraldkoch/alpine:rolling@sha256:7518a089eeb886ec8214d8c270c8911acaa41eabe44e8d0643b8feb561303b61

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
