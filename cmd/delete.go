package cmd

import (
		"fmt"
		//"os"
		//"os/exec"
		"github.com/spf13/cobra"
)


var DeleteAppCmd = &cobra.Command{
	Use: "delete",
	Short: "To Delete Previous Servers",
	Run: deleteResources,
}

func deleteResources(cmd *cobra.Command, args []string) {
	fmt.Println("Deleting Servers ....")
	// if tfdelete() != nil{

	// }
}
