FROM docker.io/library/alpine:3.23.3
ARG VERSION

#hadolint ignore=DL3018
RUN apk add --no-cache bash ca-certificates catatonit openssh rsync~"${VERSION}"
    && addgroup -g 1000 rsync
    && adduser -u 1000 -G rsync -s /bin/sh -D rsync

ENTRYPOINT ["/usr/bin/catatonit", "--", "/usr/bin/rsync"]
