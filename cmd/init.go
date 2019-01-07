package cmd

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/LoneWolf38/EasyDeploy/provider"
)

var InitServiceCmd = &cobra.Command{
	Use: "init",
	Short: "First time command",
	Run: StartInit,
}

var newConfig = viper.New()
var readConfig = viper.New()


func StartInit(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(CPath); os.IsNotExist(err) {
		ConfigInit()
		CreateKeyPair()
	} else {
		fmt.Println("A Config File Found")
		os.Exit(1)
	}
}


func ConfigInit() {
		fmt.Println("Creating a New config .... ")
		UserDetails("test/path", KeyName ,ValueInput("Github"))
		acskey := os.Getenv("AWS_ACCESS_KEY_ID")
		seckey := os.Getenv("AWS_SECRET_ACCESS_KEY")
		if len(acskey) == 0 && len(seckey) == 0{
			AWScreds(acskey,seckey,Region)	
		}else{
			AWScreds(ValueInput("AWS access key"), ValueInput("AWS secret key"),Region)	
		}
		WriteServersDetails("Default faulty Ip","faluty dns","2321312","21323","2142142","asdsd")
		fmt.Println("Config file created easyconfig.json...")
}

// Writing Servers Details in config file 

func WriteServersDetails(ip,publicdns, secgroup, vpcid,instanceid, subnetid string) {
	fmt.Println("Writing Server details in config")
	newConfig.SetConfigFile(CPath)
	newConfig.Set("server.ip",ip)
	newConfig.Set("server.dns",publicdns)
	newConfig.Set("server.SecGroup",secgroup)
	newConfig.Set("server.VpcId",vpcid)
	newConfig.Set("server.SubnetId",subnetid)
	newConfig.Set("server.InstanceId",instanceid)
	newConfig.WriteConfig()		
}

//Writing User Details in the config file in user.json

func UserDetails(keypath, keyname, github string) {
	WriteConfigFiles("keypath",keypath,CPath,"user")
	WriteConfigFiles("keyname",keyname,CPath,"user")
	WriteConfigFiles("github",github,CPath,"user")
}

//Writing aws creds in aws.json

func AWScreds(akey, skey, region string) {
	WriteConfigFiles("access_key",akey,CPath,"aws")
	WriteConfigFiles("secret_key",skey,CPath,"aws")
	WriteConfigFiles("region",region,CPath,"aws")
}

//Write Function

func WriteConfigFiles(key, value, config, objectName string ) {
	newConfig.SetConfigFile(CPath)
	newConfig.SetConfigType("json")
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


func CreateKeyPair() {
	readConfig.SetConfigFile(CPath)
	readConfig.SetConfigType("json")
	readConfig.ReadInConfig()
	fmt.Println("Setting up Environment variables....")
	os.Setenv("AWS_ACCESS_KEY_ID",readConfig.GetString("aws.access_key"))
	os.Setenv("AWS_SECRET_ACCESS_KEY",readConfig.GetString("aws.secret_key"))
	EC2keyPairCreation(KeyName)
}


func EC2keyPairCreation(keyName string) {
	fmt.Println("Creating a KeyPair for EC2 Instances...")
	svc := provider.CreateEc2Session(Region)
	keypem := provider.CreateKey(keyName,svc)
	//keyfp := *keypair.KeyFingerrint
	keyPath, err := CreateKeyPairFile(keyName,keypem)
	if err != nil{
		fmt.Println("Error in Writing key file")
		os.Exit(1)
	}else{
		readConfig.SetConfigFile(CPath)
		readConfig.SetConfigType("json")
		readConfig.Set("user.keyname",keyName)
		readConfig.Set("user.Keypath",keyPath)
		readConfig.WriteConfig()
	}
	fmt.Printf("KeyPair Created and stored in ~/%s.pem",keyName)
	//result.GetKeyName()
}

func CreateKeyPairFile(name,content string) (string, error){
	fullPath := HOME+"/"+name+".pem"
	file,err := os.Create(fullPath)
	if err != nil{
		return fullPath, err
	}
	file.WriteString(content)
	fileerr := os.Chmod(fullPath,0400)
	return fullPath, fileerr
}