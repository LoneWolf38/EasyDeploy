package main

import (
		"fmt"
		"github.com/spf13/viper"
		"os"
		"bufio"
		"log"
		"strings"
)

var ReadConfig = viper.New()
var WriteConfig = viper.New()
	

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
	writeConfig.SetConfigFile(filepath)
	writeConfig.Set("server.ip",ip)
	writeConfig.Set("server.dns",publicdns)
	writeConfig.Set("server.Projects",rProjects[:])
	writeConfig.WriteConfig()		
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
	writeConfig.SetConfigFile(config)
	object := objectName + "." + key
	writeConfig.Set(object,value)
	writeConfig.WriteConfig()
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
		fmt.Print(module[0].(map[string]interface{})["outputs"].(map[string]interface{})["ip"].(map[string]interface{})["value"])
}


