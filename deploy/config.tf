#Variables declaration 

variable "access_key" {
  description = "AWS Access Key"
}

variable "secret_key" {
  description = "AWS Secret Key"
}

variable "instance_name" {
  description = "Aws instance name"
}

variable "private_key" {
  description = "Private key used for login"
}

variable "region" {
  description = "Region where the AWS instance will be created"
  default     = "ap-south-1"                                    #For mine
}

variable "ami_region" {
  description = "AMI available for regions"
  type        = "map"

  default = {
    "ap-south-1" = "ami-0d773a3b7bb2bb1c1"
  }
}

variable "inst_type" {
  description = "Type of instance as per the projects"
  type        = "map"
}

variable "firewall" {
  description = "Security Groups as per the app"
  type        = "map"

  default = {
    "ssh_only"       = ["22"]
    "http_only"      = ["22", "80"]
    "https_only"     = ["22", "443"]
    "static_website" = ["22", "80", "443"]
  }
}

variable "deploy_type" {
  description = "Type of Deployment"
  default     = "demo"
}

variable "keyname" {
  description = "KeyPair Name"
}
