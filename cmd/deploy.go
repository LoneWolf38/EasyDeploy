package cmd

import (
		"fmt"
		"os"
		"github.com/spf13/cobra"
		"github.com/LoneWolf38/EasyDeploy/provider"
		"github.com/spf13/viper"
		 // "github.com/aws/aws-sdk-go/aws"
		 // "github.com/aws/aws-sdk-go/aws/session"
		 // "github.com/aws/aws-sdk-go/service/ec2"
)

var HOME = os.Getenv("HOME")


var DeployAppCmd = &cobra.Command{
	Use: "deploy",
	Short: "To Deploy the App",
	Run: deploy,
}

func deploy(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(CPath); os.IsNotExist(err) {
		fmt.Println("No Config File found...")
		fmt.Println("Creating a New one......")
		ConfigInit()
		os.Exit(1)	
	}else{

	} 
}

func InfoDetails() {
	updateConfig := viper.New()
	fmt.Println("Collecting Info...")
	svc := provider.CreateEc2Session(Region)
	vpcId := provider.VpcDetails(svc)
	subnetId := provider.SubnetDetails(svc)
	updateConfig.SetConfigFile(CPath)
	updateConfig.SetConfigType("json")
	updateConfig.Set("server.SubnetId",subnetId)
	updateConfig.Set("server.VpcId",vpcId)
	updateConfig.WriteConfig()
}
