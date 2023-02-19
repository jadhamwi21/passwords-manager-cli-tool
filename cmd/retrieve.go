package cmd

import (
	"github.com/jadhamwi21/passwords-manager-cli-tool/repositories"
	"github.com/spf13/cobra"
)

var retrieve_all bool
var retrieve_password_id string

func init() {
	retrieveCommand.Flags().BoolVar(&retrieve_all, "all", false, "RETRIEVE ALL PASSWORDS")
	retrieveCommand.Flags().StringVar(&retrieve_password_id, "id", "", "RETRIEVE PASSWORD BY ID")
	RootCmd.AddCommand(retrieveCommand)
}

var retrieveCommand = &cobra.Command{
	Use:   "retrieve",
	Short: "retrieve passwords",
	Run: func(cmd *cobra.Command, args []string) {
		if retrieve_all {
			repositories.PasswordsRepository.RetrieveAllPasswords()
		}
		if retrieve_password_id != "" {
			repositories.PasswordsRepository.RetrievePasswordById(retrieve_password_id)
		}
	},
}
