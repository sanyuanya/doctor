variable "TAG" {
  default = "v0.1.0"
}

group "default" {
  targets = ["task"]
}

target "task" {
  dockerfile = "Dockerfile"
  context = "."
  tags = ["sanyuanya-docker.pkg.coding.net/donglexiaochengxu/dongle/doctor:${TAG}"]
  platforms = ["linux/amd64", "linux/arm64"]
}