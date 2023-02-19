package cmd

import (
	"github.com/jadhamwi21/passwords-manager-cli-tool/repositories"
	"github.com/spf13/cobra"
)

var manage_password_id string
var manage_password_value string

func init() {
	addCommand.Flags().StringVar(&manage_password_id, "id", "", "PASSWORD ID")
	addCommand.Flags().StringVar(&manage_password_value, "password", "", "PASSWORD VALUE")
	addCommand.MarkFlagRequired("id")
	addCommand.MarkFlagRequired("password")
	manageCommand.AddCommand(addCommand)
	updateCommand.Flags().StringVar(&manage_password_id, "id", "", "PASSWORD ID")
	updateCommand.Flags().StringVar(&manage_password_value, "password", "", "PASSWORD VALUE")
	updateCommand.MarkFlagRequired("id")
	updateCommand.MarkFlagRequired("password")
	manageCommand.AddCommand(updateCommand)
	removeCommand.Flags().StringVar(&manage_password_id, "id", "", "PASSWORD ID")
	removeCommand.MarkFlagRequired("id")
	manageCommand.AddCommand(removeCommand)
	RootCmd.AddCommand(manageCommand)
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "add a new password with an id",
	Run: func(cmd *cobra.Command, args []string) {
		repositories.PasswordsRepository.AddPassword(manage_password_id, manage_password_value)
	},
}
var removeCommand = &cobra.Command{
	Use:   "remove",
	Short: "remove a password by id",
	Run: func(cmd *cobra.Command, args []string) {
		repositories.PasswordsRepository.RemovePassword(manage_password_id)
	},
}
var updateCommand = &cobra.Command{
	Use:   "update",
	Short: "update a password by id",
	Run: func(cmd *cobra.Command, args []string) {
		repositories.PasswordsRepository.UpdatePassword(manage_password_id, manage_password_value)
	},
}

var manageCommand = &cobra.Command{
	Use:   "manage",
	Short: "manage user passwords",
}
