variable "DEFAULT_TAG" {
  default = "woodpecker-gitea-publisher-golang:local"
}

// Special target: https://github.com/docker/metadata-action#bake-definition
target "docker-metadata-action" {
  tags = ["${DEFAULT_TAG}"]
}

// Default target if none specified
group "default" {
  targets = ["image-local"]
}

target "image" {
  inherits = ["docker-metadata-action"]
}

target "image-local" {
  inherits = ["image"]
  output = ["type=docker"]
}

# // must check by parent image support multi-platform
# // doc: https://docs.docker.com/reference/cli/docker/buildx/build/#platform
# // most of can as: linux/amd64 linux/386 linux/arm64/v8 linux/arm/v7 linux/arm/v6 linux/ppc64le linux/s390x
# target "image-all" {
#   inherits = ["image"]
#   platforms = [
#     "linux/amd64",
#     "linux/arm64/v8"
#   ]
# }

// must check by parent image support multi-platform
// doc: https://docs.docker.com/reference/cli/docker/buildx/build/#platform
// golang with alpine can use as: linux/amd64 linux/386 linux/arm64/v8 linux/arm/v7 linux/ppc64le linux/s390x
target "image-all" {
  inherits = ["image"]
  platforms = [
    "linux/amd64",
    "linux/386",
    "linux/arm64/v8",
    "linux/arm/v7",
    "linux/ppc64le",
    "linux/s390x",
  ]
}