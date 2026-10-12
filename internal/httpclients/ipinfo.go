package httpclients

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type IpInfoResponse struct {
	Ip string `json:"ip"`
	Hostname string `json:"hostname"`
	City string `json:"city"`
	Region string `json:"region"`
	Country string `json:"country"`
	Location string `json:"loc"`
	Org string `json:"org"`
	Postal string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme string `json:"readme"`
}

const (
	ipInfoUrl string = "ipinfo.io"
)

func GetIpInfoData() (*IpInfoResponse, error) {
	var ipinfoData IpInfoResponse

	client := &http.Client{}

	fullUrl := fmt.Sprintf("https://%s/", ipInfoUrl)

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("Error generating request: %v\n", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %v\n", err)
	}
	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(&ipinfoData)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling JSON: %v\n", err)
	}

	return &ipinfoData, nil
}
