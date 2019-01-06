package main


import (
	//"log"

	 ///"os"
	"github.com/LoneWolf38/EasyDeploy/cmd"
	//"github.com/LoneWolf38/EasyDeploy/provisioner"
)




func main() {

	cmd.RootCmd.Execute()
	//provisioner.StaticDeploy("linkto the projects",os.Getenv("HOME"))
}