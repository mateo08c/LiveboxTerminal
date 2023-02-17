package api

import (
	"fmt"
	"net/http"
)

type LoginContext struct {
	Service    string `json:"service"`
	Method     string `json:"method"`
	Parameters struct {
		ApplicationName string `json:"applicationName"`
		Username        string `json:"username"`
		Password        string `json:"password"`
	} `json:"parameters"`
}

type LoginContextResponse struct {
	Status int `json:"status"`
	Data   struct {
		ContextID string `json:"contextID"`
		Username  string `json:"username"`
		Groups    string `json:"groups"`
	} `json:"data"`
}

func (l *LoginContext) Do(args ...string) (*http.Response, error) {
	var url string
	if len(args) == 0 {
		return nil, fmt.Errorf("missing ip")
	}
	url = fmt.Sprintf("http://%s/ws", args[0])

	resp, err := Post(url, l)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewLoginContext(username string, password string) LoginContext {
	var l LoginContext
	l.Service = "sah.Device.Information"
	l.Method = "createContext"
	l.Parameters.ApplicationName = "webui"
	l.Parameters.Username = username
	l.Parameters.Password = password

	return l
}
