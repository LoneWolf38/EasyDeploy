package cmd

import (
		"fmt"
		"os"
		"github.com/spf13/cobra"
)


var CPath = os.Getenv("HOME")+"/.easyconfig.json"



var RootCmd = &cobra.Command{
	Use: "EasyDeploy",
	Short: "To show how your webapps looks in production stage",
	Long: `EasyDeploy helps user to deploy their webapps/websites in a cloud environment`,
	Version: "1.0",
	Run: easydeploy,
}

func easydeploy(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
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
