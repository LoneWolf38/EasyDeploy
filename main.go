package main


import (
	"log"
	//"github.com/LoneWolf38/EasyDeploy/cmd"
	"github.com/LoneWolf38/EasyDeploy/provisioner"
)


func main() {

	//cmd.RootCmd.Execute()
	serverConfig := provisioner.ServerConnInfo{
		Server: "13.126.222.190",
		Port: "22",
		User: "ubuntu",
		Key: "/home/drake/jarjarbinks.pem",
	}


	command := "cat ~/success"
	output, _:= provisioner.SSHCommandString(command, serverConfig)
	log.Println("Result", output)
}