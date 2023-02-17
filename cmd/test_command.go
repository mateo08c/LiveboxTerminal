package cmd

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/spf13/cobra"
	"reflect"
	"strings"
	"time"
)

func RunTest(cmd *cobra.Command, args []string) {
	start := time.Now()

	golog.Warn("Trying to login with ", Config.Username, " on ", Config.Ip)

	login, err := Client.Login(Config.Ip, Config.Username, Config.Password)
	if err != nil {
		golog.Errorf(err.Error())
		return
	}

	if login.Status != 0 {
		golog.Errorf("Login failed, please check your credentials.")
		return
	}

	golog.Info("Login successful in ", time.Since(start))
	golog.Info("ContextID: ", Client.ContextID)
	golog.Info("Cookie: ", Client.Cookie)

	golog.Warn("Trying to get current account infos")
	data, err := Client.GetCurrentUser(Config.Ip, Client.ContextID, Client.Cookie)
	if err != nil {
		golog.Errorf(err.Error())
		return
	}

	golog.Info("User: ", data.Status.User)
	golog.Info("User Groups: ", strings.Join(data.Status.Groups, ", "))

	golog.Warn("Trying to get device infos")
	device, err := Client.GetDeviceInfo(Config.Ip, Client.ContextID, Client.Cookie)
	if err != nil {
		golog.Errorf(err.Error())
		return
	}

	v := reflect.ValueOf(device.Status)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		golog.Info(fmt.Sprintf("%s: %v", typeOfS.Field(i).Name, v.Field(i).Interface()))
	}

	golog.Warn("Trying to get dhcp infos")
	dhcp, err := Client.GetDhcp(Config.Ip, Client.ContextID, Client.Cookie)
	if err != nil {
		golog.Errorf(err.Error())
		return
	}

	v = reflect.ValueOf(dhcp.Status.Dhcp.DhcpData)
	typeOfS = v.Type()

	for i := 0; i < v.NumField(); i++ {
		n := typeOfS.Field(i).Name
		if n == "SentOption" || n == "ReqOption" {
			continue
		}

		golog.Info(fmt.Sprintf("%s: %v", typeOfS.Field(i).Name, v.Field(i).Interface()))
	}

	golog.Warn("Trying to get dhcp sentoptions")
	for _, v := range dhcp.Status.Dhcp.DhcpData.SentOption {
		golog.Info(fmt.Sprintf("%s: Enable: %v - Tag: %v - Value: %v", v.Alias, v.Enable, v.Tag, v.Value))
	}

	golog.Warn("Trying to get dhcp ReqOption")
	for _, v := range dhcp.Status.Dhcp.DhcpData.ReqOption {
		golog.Info(fmt.Sprintf("%s: Enable: %v - Tag: %v - Value: %v", v.Alias, v.Enable, v.Tag, v.Value))
	}

}
