package main

import (
	"bytes"
	"encoding/json"
	"github.com/kataras/golog"
	"net/http"
	"strconv"
)

type Request struct {
	Service    string      `json:"service"`
	Method     string      `json:"method"`
	Parameters interface{} `json:"parameters"`
}

type LoginResponse struct {
	Status int `json:"status"`
	Data   struct {
		ContextID string `json:"contextID"`
		Username  string `json:"username"`
		Groups    string `json:"groups"`
	}
}

type CurrentUserResponse struct {
	Status string `json:"status"`
	Data   struct {
		ContextID string `json:"contextID"`
		Username  string `json:"username"`
		Groups    string `json:"groups"`
	}
}

type LoginParameters struct {
	ApplicationName string `json:"applicationName"`
	Username        string `json:"username"`
	Password        string `json:"password"`
}

type FirewallRule struct {
	Id                   string `json:"id"`
	Origin               string `json:"origin"`
	SourceInterface      string `json:"sourceInterface"`
	SourcePort           string `json:"sourcePort"`
	DestinationPort      string `json:"destinationPort"`
	DestinationIPAddress string `json:"destinationIPAddress"`
	SourcePrefix         string `json:"sourcePrefix"`
	Protocol             string `json:"protocol"`
	Ipversion            int    `json:"ipversion"`
	Enable               bool   `json:"enable"`
	Persistent           bool   `json:"persistent"`
}

func AddFirewallRuleV6(ipv6 string, port int, name string, protocol string) error {
	var l FirewallRule
	l.Id = name
	l.Origin = "webui"
	l.SourceInterface = "data"
	l.SourcePort = ""
	l.DestinationPort = strconv.Itoa(port)
	l.DestinationIPAddress = ipv6
	l.SourcePrefix = ""
	l.Protocol = protocol
	l.Ipversion = 6
	l.Enable = true
	l.Persistent = true

	var r Request
	r.Service = "Firewall"
	r.Method = "setPinhole"
	r.Parameters = &l

	js, err := json.Marshal(&r)
	if err != nil {
		return err
	}

	golog.Info(string(js))

	post, err := Post("http://"+IP+"/ws", js, http.Header{
		"Content-Type":  []string{"application/x-sah-ws-4-call+json"},
		"Authorization": []string{"X-Sah " + ContextID},
		"Cookie":        []string{Cookie},
		"X-Context":     []string{ContextID},
	})
	if err != nil {
		return err
	}

	body := new(bytes.Buffer)
	_, err = body.ReadFrom(post.Body)
	if err != nil {
		return err
	}

	golog.Info(body.String())

	return nil
}

func GetCurrentUser() (*CurrentUserResponse, error) {
	var l Request
	l.Service = "HTTPService"
	l.Method = "getCurrentUser"

	js, err := json.Marshal(&l)
	if err != nil {
		return nil, err
	}

	post, err := Post("http://"+IP+"/ws", js, http.Header{
		"Content-Type":  []string{"application/x-sah-ws-4-call+json"},
		"Authorization": []string{"X-Sah " + ContextID},
		"Cookie":        []string{Cookie},
		"X-Context":     []string{ContextID},
	})

	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	_, err = body.ReadFrom(post.Body)
	if err != nil {
		return nil, err
	}

	golog.Info(body.String())

	return nil, nil
}

func GetContextID(u string, p string) (*LoginResponse, error) {

	var l Request
	l.Service = "sah.Device.Information"
	l.Method = "createContext"
	l.Parameters = &LoginParameters{
		ApplicationName: "webui",
		Username:        u,
		Password:        p,
	}

	js, err := json.Marshal(&l)
	if err != nil {
		return nil, err
	}

	post, err := Post("http://"+IP+"/ws", js, http.Header{
		"Content-Type":  []string{"application/x-sah-ws-4-call+json"},
		"Authorization": []string{"X-Sah-Login"},
	})
	if err != nil {
		return nil, err
	}

	golog.Info(post)

	Cookie = post.Header.Get("Set-Cookie")

	body := new(bytes.Buffer)
	_, err = body.ReadFrom(post.Body)
	if err != nil {
		return nil, err
	}

	var resp LoginResponse
	err = json.Unmarshal(body.Bytes(), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func Post(url string, body []byte, h http.Header) (*http.Response, error) {
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	r.Header = h
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	return res, nil
}
