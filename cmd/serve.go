package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Pet project serve command",
	Run:   RunServeCommand,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func RunServeCommand(cmd *cobra.Command, args []string) {
	fmt.Println("pet project serve command called")
	fmt.Println(args)
}
