# syntax=docker/dockerfile:1

FROM docker.io/library/golang:1.24-alpine AS builder
ARG VERSION
ENV CGO_ENABLED=0
RUN go install -a -ldflags "-X=main.volsyncVersion=${VERSION}" github.com/backube/volsync@${VERSION}

FROM docker.io/library/alpine:3.22.0
ARG VERSION
RUN \
    apk add --no-cache \
        acl \
        bash \
        ca-certificates \
        coreutils \
        curl \
        tzdata \
    && \
    curl -fsSL "https://github.com/backube/volsync/archive/refs/tags/${VERSION}.tar.gz" \
        | tar xzf - -C /tmp --strip-components=1 \
    && mkdir -p /mover-rclone \
        && cp /tmp/mover-rclone/active.sh /mover-rclone/ \
        && chmod a+rx /mover-rclone/*.sh \
    && mkdir -p /mover-restic \
        && cp /tmp/mover-restic/entry.sh /mover-restic/ \
        && chmod a+rx /mover-restic/*.sh \
    && rm -rf /tmp/*
COPY --from=ghcr.io/rclone/rclone:1.70.2 /usr/local/bin/rclone /usr/local/bin/rclone
COPY --from=docker.io/restic/restic:0.18.0 /usr/bin/restic /usr/local/bin/restic
COPY --from=builder /go/bin/volsync /manager
ENTRYPOINT ["/bin/bash"]
