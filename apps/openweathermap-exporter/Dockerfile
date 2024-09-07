ARG VERSION
FROM docker.io/library/golang:1.23-alpine as openweathermap-exporter
ARG VERSION
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT=""
ARG TARGETPLATFORM
ENV CGO_ENABLED=0 \
    GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH} \
    GOARM=${TARGETVARIANT}
RUN go install -ldflags="-s -w" "github.com/haraldkoch/openweathermap-exporter@v${VERSION}"

FROM docker.io/library/alpine:3.20.3
ARG TARGETPLATFORM
ARG VERSION
ARG CHANNEL

RUN \
    apk add --no-cache \
        bash \
        ca-certificates \
        catatonit

COPY --from=openweathermap-exporter /go/bin/openweathermap-exporter /app/openweathermap-exporter

EXPOSE 2112
ENTRYPOINT ["/usr/bin/catatonit", "--"]
CMD ["/app/openweathermap-exporter"]

LABEL org.opencontainers.image.source="https://github.com/haraldkoch/openweathermap-exporter"
