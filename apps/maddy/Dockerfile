FROM docker.io/library/golang:1.25-alpine AS builder
ARG VERSION
ENV CGO_ENABLED=0

WORKDIR /maddy

RUN \
    apk add --no-cache bash ca-certificates curl \
    && curl -fsSL "https://github.com/foxcpp/maddy/archive/refs/tags/v${VERSION}.tar.gz" \
        | tar xzf - --strip-components=1 \
    && mkdir -p /pkg/data \
    && cp maddy.conf.docker /pkg/data/maddy.conf \
    && ./build.sh --builddir /tmp --destdir /pkg/ build install

FROM docker.io/library/alpine:3.22.1

RUN \
    apk add --no-cache \
        bash \
        ca-certificates \
        catatonit

COPY --from=builder /pkg/data/maddy.conf /data/maddy.conf
COPY --from=builder /pkg/usr/local/bin/maddy /bin/maddy

EXPOSE 25 143 993 587 465
ENTRYPOINT ["/usr/bin/catatonit", "--", "/bin/maddy", "-config", "/data/maddy.conf"]
