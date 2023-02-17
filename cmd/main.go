package cmd

import (
	"LiveboxTerminal/internal/client"
	"LiveboxTerminal/internal/config"
	"github.com/spf13/cobra"
)

var Config = &config.Config{}
var Client = &client.Client{}

func Run() error {
	Config.Init()

	rootCmd.PersistentFlags().StringVarP(&Config.Ip, "ip", "i", Config.Ip, "ip address")
	rootCmd.PersistentFlags().StringVarP(&Config.Username, "username", "u", Config.Username, "username")
	rootCmd.PersistentFlags().StringVarP(&Config.Password, "password", "p", Config.Password, "password")

	configCmd.Flags().StringP("ip", "i", "", "ip address")
	configCmd.Flags().StringP("username", "u", "", "username")
	configCmd.Flags().StringP("password", "p", "", "password")
	_ = configCmd.MarkFlagRequired("password")

	printConfigCmd.Flags().Bool("json", false, "print json")

	configCmd.AddCommand(printConfigCmd)
	rootCmd.AddCommand(configCmd, testCommand)

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

var rootCmd = &cobra.Command{
	Use:   "LiveboxTerminal",
	Short: "LiveboxTerminal is a CLI for Livebox",
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate a config file",
	Run:   SetConfig,
}

var printConfigCmd = &cobra.Command{
	Use:   "print",
	Short: "Display config values",
	Run:   PrintConfig,
}

var testCommand = &cobra.Command{
	Use:   "test",
	Short: "Start a series of tests",
	Run:   RunTest,
}
