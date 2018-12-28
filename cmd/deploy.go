package cmd

import (
		"fmt"
		"os"
		"os/exec"
		"github.com/spf13/cobra"
		
)


var HOME = os.Getenv("HOME")
const TF = "/usr/local/bin/terraform"
var TF_FILES = HOME+"/terraform/"


var DeployAppCmd = &cobra.Command{
	Use: "deploy",
	Short: "To Deploy the App",
	Run: deploy,
}

func deploy(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(CPath); os.IsNotExist(err) {
		fmt.Println("No Config File found...")
		fmt.Println("Creating a New one......")
		ConfigInit()	
	}
	fmt.Println("Initializing Action env....")
	if tfinit() != nil{
		fmt.Println("Terraform Init error")
	}
	fmt.Println("Server Planning....")
	if tfplan() != nil{
		fmt.Println("Terraform Plan error")
	}

	fmt.Println("Server Creation....")
	if tfapply() != nil{
		fmt.Println("Terraform Apply error")	
	}
}


