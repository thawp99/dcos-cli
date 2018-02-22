package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use: "auth",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("auth called")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
