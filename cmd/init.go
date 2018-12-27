package cmd

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var InitServiceCmd = &cobra.Command{
	Use: "init",
	Short: "First time command",
	Run: StartInit,
}

var newConfig = viper.New()


func StartInit(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(CPath); os.IsNotExist(err) {
		ConfigInit()
	} else {
		fmt.Println("A Config File Found")
	}
}


func ConfigInit() {
		fmt.Println("Creating a New config .... ")
		UserDetails(ValueInput("Private Key File Path"), ValueInput("Key Name"))
		AWScreds(ValueInput("AWS access key"), ValueInput("AWS secret key"))
		fmt.Println("Config file created .easyconfig.json...")
}

//Writing Servers Details in config file servers.json for 1 time only

// func WriteServersDetails(ip,publicdns string, rProjects []string) {
// 	fmt.Println("Writing Server details in config")
// 	newConfig.SetConfigFile(CPath)
// 	newConfig.Set("server.ip",ip)
// 	newConfig.Set("server.dns",publicdns)
// 	newConfig.Set("server.Projects",rProjects[:])
// 	newConfig.WriteConfig()		
// }


//Writing User Details in the config file in user.json

func UserDetails(keypath, keyname string) {
	WriteConfigFiles("PrivateKey",keypath,CPath,"user")
	WriteConfigFiles("KeyName",keyname,CPath,"user")
}

//Writing aws creds in aws.json

func AWScreds(akey, skey string) {
	WriteConfigFiles("access_key",akey,CPath,"aws")
	WriteConfigFiles("secret_key",skey,CPath,"aws")
}

//Write Function

func WriteConfigFiles(key, value, config, objectName string ) {
	newConfig.SetConfigFile(CPath)
	object := objectName + "." + key
	newConfig.Set(object,value)
	newConfig.WriteConfig()
}

//Input Function

func ValueInput(s string) string{
	buffer := bufio.NewReader(os.Stdin)
	fmt.Print(s+": ")
	val,err := buffer.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.Trim(val,"\n")
}
