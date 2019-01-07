package cmd

import (
		"fmt"
		"os"
		"github.com/spf13/cobra"
)

var Region = "ap-south-1"
var KeyName = "jarjarbinks"
var ami = "ami-0d773a3b7bb2bb1c1"
var instancetype = "t2.micro"
var secName = "static-website"
var secDes = "A security group for allowing ports 80 and 22 and 443"
var repo = "lonewolf38.github.io"

var CPath = os.Getenv("HOME")+"/.easyconfig"

var RootCmd = &cobra.Command{
	Use: "EasyDeploy",
	Short: "To show how your webapps looks in production stage",
	Long: `EasyDeploy helps user to deploy their webapps/websites in a cloud environment`,
	Version: "0.1",
	Run: easydeploy,
}

func easydeploy(cmd *cobra.Command, args []string) {
	if len(args) <1 {
		fmt.Println(cmd.UsageString())	
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&Region, "region", Region, "To set the region of the EC2 Instance")
	RootCmd.PersistentFlags().StringVar(&KeyName, "key",KeyName,"Keypair name")
	RootCmd.AddCommand(InitServiceCmd)
	RootCmd.AddCommand(DeployAppCmd)
	RootCmd.AddCommand(DeleteAppCmd)
}
