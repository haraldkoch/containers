DATE = formatdate( "YYYY.MM.DD", timestamp() )
APP = "actions-runner"
SOURCE = "https://github.com/actions/runner"
variable "GIT_SHA" {}

variable "VERSION" {
  // renovate: datasource=docker depName=ghcr.io/actions/actions-runner
  default = "2.336.0"
}

group "default" {
  targets = ["image-local"]
}

target "image" {
  inherits = ["docker-metadata-action"]
  args = {
    VERSION = "${VERSION}"
  }
  labels = {
    "org.opencontainers.image.vendor" = "haraldkoch"
    "org.opencontainers.image.source" = "https://github.com/haraldkoch/containers"
    "org.opencontainers.image.created" = "${DATE}"
    "org.opencontainers.image.revision" = "${GIT_SHA}"
    "org.opencontainers.image.title" = "${APP}"
    "org.opencontainers.image.url" = "${SOURCE}"
    "org.opencontainers.image.version" = "${VERSION}"
  }
  no-cache = true
}

target "image-local" {
  inherits = ["image"]
  output = ["type=docker"]
  tags = ["${APP}:${VERSION}"]
}

target "image-all" {
  inherits = ["image"]
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
}

target "docker-metadata-action" {}
