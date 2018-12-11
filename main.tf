variable "acc" {
  description = "access key"
}

variable "secret" {
  description = "secret key"
}

variable "keypath" {
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
  private_key   = "${var.keypath}"
  region        = "ap-south-1"
  keyname       = "${var.key}"
  webapp_type   = "static"

  inst_type = {
    "demo"    = "t2.micro"
    "develop" = "t2.micro"
  }
}

resource "null_resource" "static" {
  triggers {
    instance = "${module.servers.instance_id}"
  }

  depends_on = ["module.servers"]

  provisioner "file" {
    connection {
      type        = "ssh"
      user        = "ubuntu"
      private_key = "${file(var.keypath)}"
      timeout     = "2m"
      host        = "${module.servers.public_ip}"
    }

    source      = "/home/drake/upload.html"
    destination = "/var/www/html/index.html"
    on_failure  = "continue"
  }
}
