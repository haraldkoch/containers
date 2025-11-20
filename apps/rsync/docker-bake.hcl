target "docker-metadata-action" {}

variable "APP" {
  default = "rsync"
}

variable "VERSION" {
  // renovate: datasource=repology depName=alpine_3_22/rsync
  default = "3.4.1-r1"
}

variable "SOURCE" {
  default = "https://github.com/WayneD/rsync"
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
    "org.opencontainers.image.source" = "${SOURCE}"
  }
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
