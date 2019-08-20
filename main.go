package main

import (
	"fmt"

	"github.com/LoneWolf38/EasyDeploy/cmd"
)

func main() {
	fmt.Println("Executing Main Program")
	cmd.RootCmd.Execute()

}
