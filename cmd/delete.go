package cmd

import (
		"fmt"
		"os"
		"os/exec"
		"github.com/spf13/cobra"
)


var DeleteAppCmd = &cobra.Command{
	Use: "delete",
	Short: "To Delete Previous Servers",
	Run: deleteResources,
}

func deleteResources(cmd *cobra.Command, args []string) {
	fmt.Println("Deleting Servers ....")
}

func tfdelete() error{
	tfinit := exec.Command(TF,"destroy")
	tfinit.Dir = TF_FILES
	tfinit.Stdout = os.Stdout
	tfinit.Stdin = os.Stdin
	tfinit.Stderr = os.Stderr
	err := tfinit.Run()
	return err
}