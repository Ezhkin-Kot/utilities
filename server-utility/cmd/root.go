package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Aboba",
	Long:  "Abooooooooooooooooooba",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigFile("../.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Error: .env file not found: %v", err)
			fmt.Println(
				"Read how to configure it: https://github.com/Ezhkin-Kot/utilities/blob/master/server-utility/README.md",
			)
		} else {
			log.Fatalf("Error while reading config: %v", err)
		}
	}

	viper.Set("wolMac", viper.GetString("WOL_MAC"))
	viper.Set("routerName", viper.GetString("DEFAULT_ROUTER_NAME"))
	viper.Set("hostName", viper.GetString("DEFAULT_HOST_NAME"))
  viper.Set("userName", viper.GetString("DEFAULT_USER_NAME"))
  viper.Set("hostPort", viper.GetString("DEFAULT_HOST_PORT"))
	viper.Set("routerPort", viper.GetString("DEFAULT_ROUTER_PORT"))
}
