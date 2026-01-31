target "docker-metadata-action" {}

variable "APP" {
  default = "prometheus-dnssec-exporter"
}

variable "VERSION" {
  // renovate: datasource=github-releases depName=haraldkoch/prometheus-dnssec-exporter
  default = "0.7.30"
}

variable "SOURCE" {
  default = "https://github.com/haraldkoch/prometheus-dnssec-exporter"
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
