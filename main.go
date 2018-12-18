package main


import (
	"fmt"
	//"strings"
	"os"
	"os/exec"
	//"flag"
	//"bufio"
	"log"
	//"bytes"
)
var HOME = os.Getenv("HOME")
const TF = "/usr/local/bin/terraform"
var TF_FILES = HOME+"/terraform/"

var USER = os.Getenv("USER")
var PDIR string
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

func tfoutput() {
	tfoutput := exec.Command(TF,"output","-module=servers")
	tfoutput.Dir = TF_FILES
	tfoutput.Stdout = os.Stdout
	tfoutput.Stdin = os.Stdin
	tfoutput.Stderr = os.Stderr
	err := tfoutput.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s",tfoutput.Stdout)
}

func ServerOutput() {
	newconfig.SetConfigFile("./output.json")
	pro := newconfig.GetStringSlice("server.Projects")
	pro = append(pro, "go","static")
	newconfig.Set("server.Projects",pro[:])
	newconfig.WriteConfig()	
}

func FileCheck(path string) bool{
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}


func main() {
	
	//check if the config is present or not
	if FileCheck("./output.json"){
		fmt.Println("Config File Found ...")
		ReadConfig()	
	} else {
		fmt.Println("No Config File Found ...")
		ConfigInit()
		pro := []string{"django","flask","nodejs","java"}
		WriteServersDetails("123.3123.3123.312","www.example.com",pro[:])
		ServerOutput()
	}
	fmt.Println(HOME)
	tfinit()
	
	//ConfigInit()
	// WriteConfig()
	

	//Sequence of terraform commands...
	//1. init and get in the root module
	//2. plan and store the plan in a terr.tfplan file
	//3. apply and also destroy resources
	//tfinit() tfget() tfplan() tfapply() tfdestroy()
}