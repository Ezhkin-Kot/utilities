package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop server",
	Long:  "Stop server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stop server")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
