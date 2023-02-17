package api

import (
	"fmt"
	"net/http"
)

type WanDhcp struct {
	Service    string `json:"service"`
	Method     string `json:"method"`
	Parameters struct {
		Mibs string `json:"mibs"`
	} `json:"parameters"`
}

type WanDhcpResponse struct {
	Status struct {
		Dhcp struct {
			DhcpData struct {
				DHCPStatus                 string `json:"DHCPStatus"`
				LastConnectionError        string `json:"LastConnectionError"`
				Renew                      bool   `json:"Renew"`
				IPAddress                  string `json:"IPAddress"`
				SubnetMask                 string `json:"SubnetMask"`
				IPRouters                  string `json:"IPRouters"`
				DNSServers                 string `json:"DNSServers"`
				DHCPServer                 string `json:"DHCPServer"`
				LeaseTime                  int    `json:"LeaseTime"`
				LeaseTimeRemaining         int    `json:"LeaseTimeRemaining"`
				Uptime                     int    `json:"Uptime"`
				DSCPMark                   int    `json:"DSCPMark"`
				PriorityMark               int    `json:"PriorityMark"`
				Formal                     bool   `json:"Formal"`
				BroadcastFlag              int    `json:"BroadcastFlag"`
				CheckAuthentication        bool   `json:"CheckAuthentication"`
				AuthenticationInformation  string `json:"AuthenticationInformation"`
				ResetOnPhysDownTimeout     int    `json:"ResetOnPhysDownTimeout"`
				RetransmissionStrategy     string `json:"RetransmissionStrategy"`
				RetransmissionRenewTimeout int    `json:"RetransmissionRenewTimeout"`
				SendMaxMsgSize             bool   `json:"SendMaxMsgSize"`
				SentOption                 map[int]struct {
					Enable bool   `json:"Enable"`
					Alias  string `json:"Alias"`
					Tag    int    `json:"Tag"`
					Value  string `json:"Value"`
				} `json:"SentOption"`
				ReqOption map[int]struct {
					Enable bool   `json:"Enable"`
					Alias  string `json:"Alias"`
					Tag    int    `json:"Tag"`
					Value  string `json:"Value"`
				} `json:"ReqOption"`
			} `json:"dhcp_data"`
		} `json:"dhcp"`
	} `json:"status"`
}

func (d *WanDhcp) Do(args ...string) (*http.Response, error) {
	var url string
	if len(args) == 0 {
		return nil, fmt.Errorf("missing ip")
	}
	url = fmt.Sprintf("http://%s/ws", args[0])
	contextID := args[1]
	cookie := args[2]

	resp, err := PostAuth(url, d, contextID, cookie)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewWanDhcp() WanDhcp {
	var l WanDhcp
	l.Service = "NeMo.Intf.data"
	l.Method = "getMIBs"
	l.Parameters.Mibs = "dhcp"
	return l
}
