package main


import (
	"github.com/LoneWolf38/EasyDeploy/cmd"
  "fmt"
)


func main() {
  fmt.Println("Executing Main Program")
	cmd.RootCmd.Execute()
	
}
