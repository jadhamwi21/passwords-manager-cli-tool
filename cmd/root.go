package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "passwords-manager-cli",
	Short: "manage user passwords, and generate new passwords",
	Long:  "this command can help you user manage his/her passwords, such as add,delete,update passwords, with also a special feature of generating a password",
}

func Execute() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
