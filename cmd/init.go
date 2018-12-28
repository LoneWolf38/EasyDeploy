package cmd

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/awserr"
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
	}

}


func ConfigInit() {
		fmt.Println("Creating a New config .... ")
		UserDetails("test/path", "test-key-pair")
		AWScreds(ValueInput("AWS access key"), ValueInput("AWS secret key"))
		fmt.Println("Config file created easyconfig.json...")
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
	WriteConfigFiles("KeyPath",keypath,CPath,"user")
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


func CreateKeyPair() {
	readConfig.SetConfigFile(CPath)
	readConfig.ReadInConfig()
	fmt.Println("Setting up Environment variables....")
	os.Setenv("AWS_ACCESS_KEY_ID",readConfig.GetString("aws.access_key"))
	os.Setenv("AWS_SECRET_ACCESS_KEY",readConfig.GetString("aws.secret_key"))
	AwsProvider()
}


func AwsProvider() {
	fmt.Println("Creating a KeyPair for EC2 Instances...")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
			Config: aws.Config{Region: aws.String("ap-south-1")},
		}))
	svc := ec2.New(sess)
	input := &ec2.CreateKeyPairInput{
		KeyName: aws.String("test-key-pair"),
	}
	keypair, err := svc.CreateKeyPair(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok{
			switch aerr.Code(){
				default: fmt.Println(aerr.Error())
			}
		}else{
			fmt.Println(err.Error())
		}
	}

	keyname := string(*keypair.KeyName)
	keypem := string(*keypair.KeyMaterial)
	//keyfp := *keypair.KeyFingerrint
	keyPath, err := CreateKeyPairFile(keyname,keypem)
	if err != nil{
		fmt.Println("Error in Writing key file")
		os.Exit(1)
	}else{
		readConfig.SetConfigFile(CPath)
		readConfig.Set("user.keyname",keyname)
		readConfig.Set("user.Keypath",keyPath)
		readConfig.WriteConfig()
	}
	fmt.Printf("KeyPair Created and stored in ~/.%s.pem",keyname)
	//result.GetKeyName()
}

func CreateKeyPairFile(name,content string) (string, error){
	fullPath := HOME+"/."+name+".pem"
	file,err := os.Create(fullPath)
	if err != nil{
		return fullPath, err
	}
	file.WriteString(content)
	fileerr := os.Chmod(fullPath,0400)
	return fullPath, fileerr
}