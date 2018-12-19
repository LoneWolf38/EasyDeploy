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

var filepath = "./output.json"
	

//Creating a new Config file  

func ConfigInit() {
		fmt.Println("Creating a New config .... ")
		UserDetails(ValueInput("Private Key File Path"), ValueInput("Key Name"))
		AWScreds(ValueInput("AWS access key"), ValueInput("AWS secret key"))
		fmt.Println("Config file created output.json...")
}

//Writing Servers Details in config file servers.json for 1 time only

func WriteServersDetails(ip,publicdns string, rProjects []string) {
	fmt.Println("Writing Server details in config")
	newconfig.SetConfigFile(filepath)
	newconfig.Set("server.ip",ip)
	newconfig.Set("server.dns",publicdns)
	newconfig.Set("server.Projects",rProjects[:])
	newconfig.WriteConfig()		
}


//Writing User Details in the config file in user.json

func UserDetails(keypath, keyname string) {
	WriteConfig("PrivateKey",keypath,filepath,"user")
	WriteConfig("KeyName",keyname,filepath,"user")
}

//Writing aws creds in aws.json

func AWScreds(akey, skey string) {
	WriteConfig("access_key",akey,filepath,"aws")
	WriteConfig("secret_key",skey,filepath,"aws")
}

//Write Function

func WriteConfig(key, value, config, objectName string ) {
	newconfig.SetConfigFile(config)
	object := objectName + "." + key
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

func ReadtfConfig() {
		var module []interface{}
		viper.SetConfigFile("./terraform.tfstate")
		viper.SetConfigType("json")
		viper.ReadInConfig()
		module = viper.Get("modules").([]interface{})
		//output value
		fmt.Print(module[0].(map[string]interface{})["outputs"].(map[string]interface{})["link"].(map[string]interface{})["value"])
}


