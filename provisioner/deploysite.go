package provisioner

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func ServerConf(confPath string) ServerConnInfo {
	confg = viper.new()
	confg.SetConfigFile(confPath)
	confg.SetConfigType("json")
	confg.ReadInConfig()
	return ServerConnInfo{
		Server: confg.GetString("server.ip"),
		Port : "22",
		User : "ubuntu",
		key : confg.GetString("user.keypath"),
	}
}

func ServerSetup() {
	commandList := {
		"sudo apt update",
		"sudo apt install -y apt-transport-https ca-certificates curl software-properties-common"
		"curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -",
		"sudo add-apt-repository ""deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable""",
		"sudo apt update",
		"apt-cache policy docker-ce",
		"sudo apt install docker-ce",
		"sudo systemctl status docker",
		"sudo usermod -aG docker ubuntu",
		"sudo su - ubuntu"
	}
}