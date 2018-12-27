#output

output "instance_id" {
  value = "${aws_instance.static.id}"
}

output "hosted_in" {
  value = "${aws_instance.static.public_dns}"
}

output "public_ip" {
  value = "${aws_instance.static.public_ip}"
}
