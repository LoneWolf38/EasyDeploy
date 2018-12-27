package cmd

import (
		"fmt"
		"os"
		"github.com/spf13/cobra"
)



var CPath = os.Getenv("HOME")+"/.easyconfig.json"

var RootCmd = &cobra.Command{
	Use: "easydeploy",
	Short: "To show how your webapps looks in production stage",
	Long: `EasyDeploy helps user to deploy their webapps/websites in a cloud and test whether all their services are running and also test how their app will run in a production env`,
	Version: "0.1",
	Run: easydeploy,
}

func easydeploy(cmd *cobra.Command, args []string) {
	if len(args) <1 {
		fmt.Println(cmd.UsageString())	
	}
}

func init() {
	RootCmd.AddCommand(InitServiceCmd)
	RootCmd.AddCommand(DeployAppCmd)
	//RootCmd.AddCommand(DeleteAppCmd)
}
