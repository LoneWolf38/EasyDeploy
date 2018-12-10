#output

output "Instance Created" {
  value = "${aws_instance.static.id}"
}
