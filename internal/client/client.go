package client

import (
	"LiveboxTerminal/internal/api"
	"encoding/json"
)

type Client struct {
	ContextID string
	Cookie    string
}

func (c *Client) GetDeviceInfo(ip string, contextID string, cookie string) (*api.DeviceInfoResponse, error) {
	di := api.NewDeviceInfo()
	resp, err := di.Do(ip, contextID, cookie)
	if err != nil {
		return nil, err
	}

	var r *api.DeviceInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) GetDhcp(ip string, contextID string, cookie string) (*api.WanDhcpResponse, error) {
	di := api.NewWanDhcp()
	resp, err := di.Do(ip, contextID, cookie)
	if err != nil {
		return nil, err
	}

	var r *api.WanDhcpResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) Login(ip string, username string, password string) (*api.LoginContextResponse, error) {
	ct := api.NewLoginContext(username, password)
	resp, err := ct.Do(ip)
	if err != nil {
		return nil, err
	}

	var r *api.LoginContextResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	c.ContextID = r.Data.ContextID
	c.Cookie = resp.Header.Get("Set-Cookie")

	return r, nil
}

func (c *Client) GetCurrentUser(ip string, contextID string, cookie string) (*api.CurrentUserResponse, error) {
	cr := api.NewCurrentUsers()
	resp, err := cr.Do(ip, contextID, cookie)
	if err != nil {
		return nil, err
	}

	var r *api.CurrentUserResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
