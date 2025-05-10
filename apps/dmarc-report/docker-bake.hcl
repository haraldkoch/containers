target "docker-metadata-action" {}

variable "VERSION" {
  // renovate: datasource=docker depName=gutmensch/dmarc-report
  default = "1.4.5"
}

variable "SOURCE" {
  default = "https://github.com/gutmensch/docker-dmarc-report"
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
