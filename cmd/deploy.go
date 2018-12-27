package cmd

import (
		"fmt"
		"github.com/spf13/cobra"
)


var DeployAppCmd = &cobra.Command{
	Use: "deploy",
	Run: deploy,
}

func deploy(cmd *cobra.Command, args []string) {
	
}