package main

import (
		"fmt"
		"github.com/spf13/viper"
		"os"
		"bufio"
		"log"
		"strings"
)

var oldconfig = viper.New()
var newconfig = viper.New()


	

//Creating a new Config file  

func ConfigInit() {
		


}

//Writing Servers Details in config file servers.json

func ServersDetails(ip,publicdns, runningProjects string) {
	WriteConfig("ip",ip,"./servers.json","server")
	WriteConfig("public_dns",publicdns,"./servers.json","server")
	WriteConfig("Running-Projects",runningProjects,"./servers.json","server")		
}


//Writing User Details in the config file in user.json

func UserDetails(keypath, keyname string) {
	WriteConfig("PrivateKey",keypath,"./user.json","user")
	WriteConfig("KeyName",keyname,"./user.json","user")
}

//Writing aws creds in aws.json

func AWScreds(akey, skey string) {
	WriteConfig("access_key",akey,"./aws.json","aws")
	WriteConfig("secret_key",skey,"./aws.json","aws")
}

//Write Function

func WriteConfig(key, value, config, objectName string ) {
	newconfig.SetConfigFile(config)
	object := objectName + "." + key
	fmt.Println(object)
	newconfig.Set(object,value)
	newconfig.WriteConfig()
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

//Read Function for config

func ReadConfig(path string) {
	oldconfig.SetConfigFile(path)

}


