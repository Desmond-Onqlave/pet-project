package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Pet project serve command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pet project serve command called")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
