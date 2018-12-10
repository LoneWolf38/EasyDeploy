access_key = ""
secret_key = ""
instance_name = "MyWebsite"
private_key = "${file("/home/drake/jarjarbinks.pem")}"
inst_type = {
	"demo" = "t2.micro"
	"develop"  = "t2.micro"
}

webapp_type = "static"
keyname = "jarjarbinks"