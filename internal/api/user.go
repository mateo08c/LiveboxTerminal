package api

import (
	"fmt"
	"net/http"
)

type CurrentUser struct {
	Service    string `json:"service"`
	Method     string `json:"method"`
	Parameters struct {
	} `json:"parameters"`
}

type CurrentUserResponse struct {
	Status struct {
		User   string   `json:"user"`
		Groups []string `json:"groups"`
	} `json:"status"`
}

func (l *CurrentUser) Do(args ...string) (*http.Response, error) {
	var url string
	if len(args) == 0 {
		return nil, fmt.Errorf("missing ip")
	}
	url = fmt.Sprintf("http://%s/ws", args[0])
	contextID := args[1]
	cookie := args[2]

	resp, err := PostAuth(url, l, contextID, cookie)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewCurrentUsers() CurrentUser {
	var l CurrentUser
	l.Service = "HTTPService"
	l.Method = "getCurrentUser"
	return l
}
