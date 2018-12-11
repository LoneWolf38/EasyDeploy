#Main entry point of the app

provider "aws" {
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
  region     = "${var.region}"
}

resource "aws_security_group" "http_s" {
  name = "normal"

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    self        = true
    description = "Http"
  }

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    self        = true
    description = "Https"
  }

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    self        = true
    description = "Remote Login"
  }

  egress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
  }
}

resource "aws_instance" "static" {
  ami             = "${lookup(var.ami_region, var.region)}"
  instance_type   = "${lookup(var.inst_type, var.deploy_type)}"
  key_name        = "${var.keyname}"
  security_groups = ["${aws_security_group.http_s.id}"]
  subnet_id       = "${var.subnetid}"

  tags {
    Name = "${var.instance_name}"
  }

  connection {
    user        = "ubuntu"
    type        = "ssh"
    private_key = "${file(var.private_key)}"
    timeout     = "2m"
  }

  provisioner "remote-exec" {
    inline     = ["sleep 1m", "sudo apt-get update", "sudo apt-get -y install nginx"]
    on_failure = "continue"
  }

  #provisioner "file" {}
}
