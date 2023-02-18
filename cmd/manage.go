package cmd

import (
	"github.com/jadhamwi21/passwords-manager-cli-tool/repositories"
	"github.com/spf13/cobra"
)

var PASSWORD_ID string
var PASSWORD_VALUE string

func init() {
	addCommand.Flags().StringVar(&PASSWORD_ID, "id", "", "PASSWORD ID")
	addCommand.Flags().StringVar(&PASSWORD_VALUE, "password", "", "PASSWORD VALUE")
	addCommand.MarkFlagRequired("id")
	addCommand.MarkFlagRequired("password")
	manageCommand.AddCommand(addCommand)
	updateCommand.Flags().StringVar(&PASSWORD_ID, "id", "", "PASSWORD ID")
	updateCommand.Flags().StringVar(&PASSWORD_VALUE, "password", "", "PASSWORD VALUE")
	updateCommand.MarkFlagRequired("id")
	updateCommand.MarkFlagRequired("password")
	manageCommand.AddCommand(updateCommand)
	removeCommand.Flags().StringVar(&PASSWORD_ID, "id", "", "PASSWORD ID")
	removeCommand.MarkFlagRequired("id")
	manageCommand.AddCommand(removeCommand)
	RootCmd.AddCommand(manageCommand)
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "add a new password with an id",
	Run: func(cmd *cobra.Command, args []string) {
		repositories.PasswordsRepository.AddPassword(PASSWORD_ID, PASSWORD_VALUE)
	},
}
var removeCommand = &cobra.Command{
	Use:   "remove",
	Short: "remove a password by id",
	Run: func(cmd *cobra.Command, args []string) {
		repositories.PasswordsRepository.RemovePassword(PASSWORD_ID)
	},
}
var updateCommand = &cobra.Command{
	Use:   "update",
	Short: "update a password by id",
	Run: func(cmd *cobra.Command, args []string) {
		repositories.PasswordsRepository.UpdatePassword(PASSWORD_ID, PASSWORD_VALUE)
	},
}

var manageCommand = &cobra.Command{
	Use:   "manage",
	Short: "manage user passwords",
}
