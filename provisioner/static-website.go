package provisioner

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

)

const PORT = "22"
const USER = "ubuntu"
const servDir = "/var/www/html"


var startCommandList = []string {
			"sudo su -",
			"apt-get update",
			"apt-get install -y git apache2", 
			"systemctl enable apache2",
			"systemctl start apache2",
		} 
var apacheSystemCommand = "sudo systemctl restart apache2"

var gitCloneCommand = "git clone "

func StaticDeploy(url, CPath,projectName string) {
	viper.SetConfigFile(CPath)
	viper.SetConfigType("json")
	viper.ReadInConfig()
	server_ip := viper.GetString("server.ip")
	keyFile := viper.GetString("user.keypath")
	fmt.Println(keyFile+""+server_ip)
	svr := ServerConnInfo{
		Server: server_ip,
		Port: PORT,
		User: USER,
		Key: keyFile,
	}
	srvDir := fmt.Sprintf("/var/www/%s",projectName)
	github := gitCloneCommand+url
	startCommandList = append(startCommandList,	fmt.Sprintf("mkdir %s",srvDir))
 	startCommandList = append(startCommandList,github)

 	for _, cmd := range startCommandList {
 		success, execError := SSHCommandBool(cmd,svr) 
 		if success != true{
 			fmt.Println(execError)
 			os.Exit(1)
 		}
 	}
}