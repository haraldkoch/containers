target "docker-metadata-action" {}

variable "APP" {
  default = "olivetin"
}

variable "VERSION" {
  // renovate: datasource=github-releases depName=OliveTin/OliveTin
  default = "2025.7.29"
}

variable "HTTPIE_VERSION" {
  // renovate: datasource=github-releases depName=httpie/cli
  default = "3.2.4"
}


variable "SOURCE" {
  default = "https://github.com/OliveTin/OliveTin"
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
    "linux/amd64"
  ]
}
