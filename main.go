package main


import (
	"fmt"
	"strings"
	"os"
	"os/exec"
	//"flag"
	"bufio"
	"log"
	"bytes"
)

const TF = "/usr/local/bin/terraform"
const TF_FILES = "/home/drake/Projects/EasyDeploy"
var HOME = os.Getenv("HOME")
var USER = os.Getenv("USER")
var PDIR string
var TFCMD [4]string 




func awshandle() {
	
}

func tfcmd(cmd string) {
	if strings.EqualFold(cmd, "apply"){
		TFCMD = [4]string{"apply","terr.plan","",""}
	}else if strings.EqualFold(cmd, "plan"){
		TFCMD = [4]string{"plan","-out","terr.plan",""}
	}else if strings.EqualFold(cmd, "init"){	
		TFCMD = [4]string{"init","","",""}
	}else {
		fmt.Printf("Wrong input")
	}
}


func tfhandle() {
	tfinit := exec.Command(TF,TFCMD[0],TFCMD[1],TFCMD[2],TFCMD[3])
	tfinit.Dir = TF_FILES
	var out bytes.Buffer
	tfinit.Stdout = &out
	err := tfinit.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(out.String())
}

func main() {
	dir := bufio.NewReader(os.Stdin)
	PDIR, err := dir.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	enverr := os.Setenv("TF_VAR_DIR",PDIR)
	if enverr != nil{
		log.Fatal(enverr)
	}
	tfcmd("init")
	tfhandle()
}