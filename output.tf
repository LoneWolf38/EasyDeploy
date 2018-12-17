output "Website Deployed in" {
  value = "${module.servers.public_ip}"
}

output "Goto The link" {
  value = "https://${module.servers.hosted_in}"
}

output "projectname" {
  value = ""
}
