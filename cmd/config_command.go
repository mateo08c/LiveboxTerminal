package cmd

import (
	"github.com/kataras/golog"
	"github.com/spf13/cobra"
	"strings"
)

func SetConfig(cmd *cobra.Command, args []string) {
	file, err := cmd.Flags().GetString("path")
	if err != nil {
		panic(err)
	}

	ip, err := cmd.Flags().GetString("ip")
	if err != nil {
		panic(err)
	}

	username, err := cmd.Flags().GetString("username")
	if err != nil {
		panic(err)
	}

	password, err := cmd.Flags().GetString("password")
	if err != nil {
		panic(err)
	}

	err = Config.UpdateConfigFile(file, ip, username, password)
	if err != nil {
		panic(err)
	}
}

func PrintConfig(cmd *cobra.Command, args []string) {
	json, err := cmd.Flags().GetBool("json")
	if err != nil {
		panic(err)
	}

	if json {
		j, err := Config.ToJson()
		if err != nil {
			panic(err)
		}
		println(j)
	} else {
		golog.Infof("URL: http://%s", Config.Ip)
		golog.Infof("Ip: %s", Config.Ip)
		golog.Infof("Username: %s", Config.Username)
		golog.Infof("Password: %s", strings.Repeat("*", len(Config.Password)))
	}
}
