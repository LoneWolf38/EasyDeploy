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
	// Initializing
	newconfig.SetConfigName("output1")
	newconfig.AddConfigPath(".")
	newconfig.SetConfigType("json")
	fmt.Println("Creating fresh config...")
	username := configInput("UserName")
	newconfig.Set("user.name",username)
	access_key := configInput("AWS Access Key")
	newconfig.Set("user.access_key",access_key)
	secret_key := configInput("AWS Secret Key")
	newconfig.Set("user.secret_key",secret_key)
	newconfig.WriteConfig()
}


func WriteConfig() {
	newconfig.SetConfigName("output1")
	newconfig.AddConfigPath(".")
	newconfig.SetConfigType("json")
	newconfig.Set("server.ip","13.132.313.133")
	newconfig.Set("server.public_dns","www.example.com")
	newconfig.Set("server.ndock",123)
	newconfig.WriteConfig()
}

func configInput(s string) string{
	buffer := bufio.NewReader(os.Stdin)
	fmt.Print(s+": ")
	val,err := buffer.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.Trim(val,"\n")
}


func ReadConfig() {
	fmt.Println("Searching for config files")
	oldconfig.SetConfigName("output")
	oldconfig.AddConfigPath(".")
	oldconfig.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config not found....")
		ConfigInit()
	} else {
		keys := oldconfig.GetString("drake.access_key")
		fmt.Println(keys)
	}
}


