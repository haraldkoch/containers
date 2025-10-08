# syntax=docker/dockerfile:1

FROM docker.io/library/golang:1.25-alpine AS builder
ARG VERSION
ENV CGO_ENABLED=0
RUN go install -ldflags="-s -w" "github.com/haraldkoch/openweathermap-exporter@v${VERSION}"

FROM docker.io/library/alpine:3.22.2
RUN \
    apk add --no-cache \
        bash \
        ca-certificates \
        catatonit

COPY --from=builder /go/bin/openweathermap-exporter /app/openweathermap-exporter

EXPOSE 2112
ENTRYPOINT ["/usr/bin/catatonit", "--", "/app/openweathermap-exporter"]
