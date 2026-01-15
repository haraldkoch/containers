FROM docker.io/library/alpine:3.23.2
ARG VERSION

#hadolint ignore=DL3018
RUN apk add --no-cache bash ca-certificates catatonit openssh rsync~"${VERSION}"
ENTRYPOINT ["/usr/bin/catatonit", "--", "/usr/bin/rsync"]
