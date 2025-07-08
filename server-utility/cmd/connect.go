package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to server",
	Run: func(cmd *cobra.Command, args []string) {
		hostName := viper.GetString("hostName")
		serverPort := viper.GetString("serverPort")
		userName := viper.GetString("userName")

		Connect(hostName, userName, serverPort)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.Flags().
		StringP("server", "s", "", "Host name in Tailscale to connect")
	connectCmd.Flags().
		StringP("user", "u", "", "User name in server to connect")
	connectCmd.Flags().StringP("port", "p", "22", "Server port to connect")

	viper.BindPFlag("hostName", connectCmd.Flags().Lookup("server"))
	viper.BindPFlag("userName", connectCmd.Flags().Lookup("user"))
	viper.BindPFlag("serverPort", connectCmd.Flags().Lookup("port"))
}
