package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long:  "Start server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start server")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
