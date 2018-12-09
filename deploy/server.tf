#Declare resources and provisions
resource "aws_instance" "myweb" {
  ami             = "${lookup(var.ami_region, var.region)}"
  instance_type   = "${lookup(var.inst_type, var.deploy_type)}"
  key_name        = "${var.keyname}"
  security_groups = ["${aws_security_group.http_s.id}"]

  tags {
    Name = "${var.instance_name}"
  }
}
