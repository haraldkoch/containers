# syntax=docker/dockerfile:1

FROM docker.io/library/golang:1.24-alpine AS builder
ARG VERSION
ENV CGO_ENABLED=0
RUN go install -ldflags="-s -w" "github.com/fffonion/tplink-plug-exporter@v${VERSION}"

FROM docker.io/library/alpine:3.22.0
RUN \
    apk add --no-cache \
        bash \
        ca-certificates \
        catatonit

COPY --from=builder /go/bin/tplink-plug-exporter /app/tplink-plug-exporter

EXPOSE 9233
ENTRYPOINT ["/usr/bin/catatonit", "--", "/app/tplink-plug-exporter"]
