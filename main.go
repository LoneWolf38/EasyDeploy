package main


import (
	//"log"

	"os"
	//"github.com/LoneWolf38/EasyDeploy/cmd"
	"github.com/LoneWolf38/EasyDeploy/provisioner"
)

var CPath = os.Getenv("HOME")+"/.easyconfig.json"


func main() {

	//cmd.RootCmd.Execute()
	provisioner.StaticDeploy("linkto the projects",CPath)
}