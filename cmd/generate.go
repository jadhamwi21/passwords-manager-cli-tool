package cmd

import (
	"fmt"
	"log"

	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"
)

var length int
var digits int
var symbols int
var uppercase bool
var repeat bool

func init() {

	generateCommand.Flags().IntVarP(&length, "characters", "c", 64, "NUMBER OF CHARACTERS")
	generateCommand.Flags().IntVarP(&digits, "digits", "d", 10, "NUMBER OF DIGITS")
	generateCommand.Flags().IntVarP(&symbols, "symbols", "s", 10, "NUMBER OF SYMBOLS")
	generateCommand.Flags().BoolVarP(&uppercase, "uppercase", "u", false, "ALLOW UPPERCASE")
	generateCommand.Flags().BoolVarP(&repeat, "repeat", "r", false, "ALLOW REPEAT CHARACTERS")
	RootCmd.AddCommand(generateCommand)
}

var generateCommand = &cobra.Command{
	Use:   "generate",
	Short: "generate password",
	Run: func(cmd *cobra.Command, args []string) {
		password, err := password.Generate(length, digits, symbols, uppercase, repeat)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(password)
	},
}
