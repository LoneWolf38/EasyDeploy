#Declare the security groups, networks and keypairs

resource "aws_security_group" "http_s" {
  name = "https"

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 80
    to_port     = 80
    protocol    = "http"
    self        = true
    description = "Http"
  }

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 443
    to_port     = 443
    protocol    = "https"
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
}
