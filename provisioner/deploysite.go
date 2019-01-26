package provisioner

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func ServerConf(confPath string) ServerConnInfo {
	confg := viper.New()
	confg.SetConfigFile(confPath)
	confg.SetConfigType("json")
	confg.ReadInConfig()
	return ServerConnInfo{
		Server: confg.GetString("server.ip"),
		Port : "22",
		User : "ubuntu",
		Key : confg.GetString("user.keypath"),
	}
}

func ServerSetup(url,CPath,repo string) {
	svr := ServerConf(CPath)

	commandList := []string{
		"sudo apt update",
		"sudo apt install -y apt-transport-https ca-certificates curl software-properties-common",
		"curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -",
		"sudo add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable\"",
		"sudo apt update",
		"apt-cache policy docker-ce",
		"sudo apt install docker-ce",
		"sudo systemctl status docker",
		"sudo usermod -aG docker ubuntu",
		"sudo su - ubuntu",
	}

	for _, cmd := range commandList {
 		success, execError := SSHCommandBool(cmd,svr) 
 		if success != true{
 			fmt.Println(execError)
 			os.Exit(1)
		}
  	}	
}