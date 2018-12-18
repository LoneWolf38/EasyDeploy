package main


import (
	"fmt"
	//"strings"
	"os"
	"os/exec"
	//"flag"
	"bufio"
	"log"
	//"bytes"
)

const TF = "/usr/local/bin/terraform"
const TF_FILES = "terraform/"
var HOME = os.Getenv("HOME")
var USER = os.Getenv("USER")
var PDIR string



// AWS Credentials entry

func AWScreds() {
	fmt.Print("Enter your AWS access_key: ")
	key := bufio.NewReader(os.Stdin)
	access_key, err := key.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Setting the environment variables.....")
	seterr := os.Setenv("TF_VAR_acc",access_key)
	if seterr != nil {
		log.Fatal(seterr)
	}
	
}

// Terraform Init Command


func tfinit() {
	tfinit := exec.Command(TF,"init")
	tfinit.Dir = TF_FILES
	tfinit.Stdout = os.Stdout
	tfinit.Stdin = os.Stdin
	tfinit.Stderr = os.Stderr
	err := tfinit.Run()
	if err != nil {
		log.Fatal(tfinit.Stderr)
	}
	fmt.Println(tfinit.Stdout)
}

// Terraform Plan command

func tfplan() {
	tfplan := exec.Command(TF,"plan","-out","terr.tfplan")
	tfplan.Dir = TF_FILES
	tfplan.Stdout = os.Stdout
	tfplan.Stdin = os.Stdin
	tfplan.Stderr = os.Stderr
	err := tfplan.Run()
	if err != nil {
		log.Fatal(tfplan.Stderr)
	}
	fmt.Println(tfplan.Stdout)
}

// Terraform Apply command

func tfapply() {
	tfapply := exec.Command(TF,"apply","terr.tfplan")
	tfapply.Dir = TF_FILES
	tfapply.Stdout = os.Stdout
	tfapply.Stdin = os.Stdin
	tfapply.Stderr = os.Stderr
	err := tfapply.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s",tfapply.Stdout)
}

// Terraform Show Command for users

func tfshow() {
	tfshow := exec.Command(TF,"show")
	tfshow.Dir = TF_FILES
	tfshow.Stdout = os.Stdout
	tfshow.Stdin = os.Stdin
	tfshow.Stderr = os.Stderr
	err := tfshow.Run()	
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%s",tfshow.Stdout)
}


// Terraform get command to load modules

func tfget() {
	tfget := exec.Command(TF,"get")
	tfget.Dir = TF_FILES
	tfget.Stdout = os.Stdout
	tfget.Stdin = os.Stdin
	tfget.Stderr = os.Stderr
	err := tfget.Run()	
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%s",tfget.Stdout)	
}

// Terraform destroy command


func tfdestroy() {
	tfdel := exec.Command(TF,"destroy")
	tfdel.Dir = TF_FILES
	tfdel.Stdout = os.Stdout
	tfdel.Stdin = os.Stdin
	tfdel.Stderr = os.Stderr
	err := tfdel.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All resources are deleted")
}

// Terraform output command for getting the output for a JSON Files

// func tfoutput() {
// 	tfoutput := exec.Command(TF,"output","-module=servers")
// 	tfoutput.Dir = TF_FILES
// 	tfoutput.Stdout = os.Stdout
// 	tfoutput.Stdin = os.Stdin
// 	tfoutput.Stderr = os.Stderr

// }


func main() {
	// Read the project directory  and project name and setting them as terraform env variables
	// dir := bufio.NewReader(os.Stdin)
	// PDIR, err := dir.ReadString('\n')
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// enverr := os.Setenv("TF_VAR_projectdir",PDIR)
	// if enverr != nil{
	// 	log.Fatal(enverr)
	// }
	// AWS credentials input 
	//AWScreds()
	//tfinit()
	fmt.Println("From main to config.go")
	ConfigInit()
	WriteConfig()

	//Sequence of terraform commands...
	//1. init and get in the root module
	//2. plan and store the plan in a terr.tfplan file
	//3. apply and also destroy resources
	//tfinit() tfget() tfplan() tfapply() tfdestroy()
}