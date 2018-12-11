#output

output "Instance Created" {
  value = "${aws_instance.static.id}"
}

output "Website Hosted In" {
  value = "${aws_instance.static.public_dns}"
}
