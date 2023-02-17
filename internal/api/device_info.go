package api

import (
	"fmt"
	"net/http"
	"time"
)

type DeviceInfo struct {
	Service    string `json:"service"`
	Method     string `json:"method"`
	Parameters struct {
	} `json:"parameters"`
}

type DeviceInfoResponse struct {
	Status struct {
		Manufacturer                             string    `json:"Manufacturer"`
		ManufacturerOUI                          string    `json:"ManufacturerOUI"`
		ModelName                                string    `json:"ModelName"`
		Description                              string    `json:"Description"`
		ProductClass                             string    `json:"ProductClass"`
		SerialNumber                             string    `json:"SerialNumber"`
		HardwareVersion                          string    `json:"HardwareVersion"`
		SoftwareVersion                          string    `json:"SoftwareVersion"`
		RescueVersion                            string    `json:"RescueVersion"`
		ModemFirmwareVersion                     string    `json:"ModemFirmwareVersion"`
		EnabledOptions                           string    `json:"EnabledOptions"`
		AdditionalHardwareVersion                string    `json:"AdditionalHardwareVersion"`
		AdditionalSoftwareVersion                string    `json:"AdditionalSoftwareVersion"`
		SpecVersion                              string    `json:"SpecVersion"`
		ProvisioningCode                         string    `json:"ProvisioningCode"`
		UpTime                                   int       `json:"UpTime"`
		FirstUseDate                             time.Time `json:"FirstUseDate"`
		DeviceLog                                string    `json:"DeviceLog"`
		VendorConfigFileNumberOfEntries          int       `json:"VendorConfigFileNumberOfEntries"`
		ManufacturerURL                          string    `json:"ManufacturerURL"`
		Country                                  string    `json:"Country"`
		ExternalIPAddress                        string    `json:"ExternalIPAddress"`
		DeviceStatus                             string    `json:"DeviceStatus"`
		NumberOfReboots                          int       `json:"NumberOfReboots"`
		UpgradeOccurred                          bool      `json:"UpgradeOccurred"`
		ResetOccurred                            bool      `json:"ResetOccurred"`
		RestoreOccurred                          bool      `json:"RestoreOccurred"`
		StandbyOccurred                          bool      `json:"StandbyOccurred"`
		XSOFTATHOMECOMAdditionalSoftwareVersions string    `json:"X_SOFTATHOME-COM_AdditionalSoftwareVersions"`
		BaseMAC                                  string    `json:"BaseMAC"`
	} `json:"status"`
}

func (d *DeviceInfo) Do(args ...string) (*http.Response, error) {
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

func NewDeviceInfo() DeviceInfo {
	var l DeviceInfo
	l.Service = "DeviceInfo"
	l.Method = "get"
	return l
}
