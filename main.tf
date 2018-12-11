variable "acc" {
  description = "access key"
}

variable "secret" {
  description = "secret key"
}

variable "path" {
  description = "path to file"
}

variable "key" {
  description = "describe your variable"
}

module "servers" {
  source        = "./modules"
  access_key    = "${var.acc}"
  secret_key    = "${var.secret}"
  instance_name = "static-website"
  private_key   = "${var.path}"
  region        = "ap-south-1"
  keyname       = "${var.key}"
  webapp_type   = "static"

  inst_type = {
    "demo"    = "t2.micro"
    "develop" = "t2.micro"
  }
}
