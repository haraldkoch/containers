target "docker-metadata-action" {}

variable "VERSION" {
  // renovate: datasource=repology depName=alpine_3_21/rsync
  default = "3.4.0"
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
}

target "image-all" {
  inherits = ["image"]
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
}
