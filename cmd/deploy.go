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


func tfinit() error{
	tfinit := exec.Command(TF,"init")
	tfinit.Dir = TF_FILES
	tfinit.Stdout = os.Stdout
	tfinit.Stdin = os.Stdin
	tfinit.Stderr = os.Stderr
	err := tfinit.Run()
	return err
}

// Terraform Plan command

func tfplan() error{
	tfplan := exec.Command(TF,"plan","-out","terr.tfplan")
	tfplan.Dir = TF_FILES
	tfplan.Stdout = os.Stdout
	tfplan.Stdin = os.Stdin
	tfplan.Stderr = os.Stderr
	err := tfplan.Run()
	return err
}

// Terraform Apply command

func tfapply() error{
	tfapply := exec.Command(TF,"apply","terr.tfplan")
	tfapply.Dir = TF_FILES
	tfapply.Stdout = os.Stdout
	tfapply.Stdin = os.Stdin
	tfapply.Stderr = os.Stderr
	err := tfapply.Run()
	return err
}

func tfget() error{
	tfget := exec.Command(TF,"get")
	tfget.Dir = TF_FILES
	tfget.Stdout = os.Stdout
	tfget.Stdin = os.Stdin
	tfget.Stderr = os.Stderr
	err := tfget.Run()	
	return err
}
