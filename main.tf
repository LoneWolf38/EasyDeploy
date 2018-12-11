module "servers" {
  source        = "./modules"
  access_key    = "AKIAJWKNY6ONSEZKXXZQ"
  secret_key    = "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"
  instance_name = "static-website"
  private_key   = "/home/drake/jarjarbinks.pem"
  region        = "ap-south-1"
  keyname       = "jarjarbinks"
  webapp_type   = "static"

  inst_type = {
    "demo"    = "t2.micro"
    "develop" = "t2.micro"
  }
}
