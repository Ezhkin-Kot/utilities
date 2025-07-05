package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect [-s server-name] [-u user] [-p port]",
	Short: "Connect to server",
	Long:  "Connect to server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Connect to server")
	},
}

func init() {
	connectCmd.Flags().
		StringVarP(&serverName, "server", "s", "server", "server name to connect")
	connectCmd.Flags().
		StringVarP(&userName, "user", "u", "root", "user name in server to connect")
	connectCmd.Flags().StringVarP(&port, "port", "p", "22", "port to connect")
	rootCmd.AddCommand(connectCmd)
}
