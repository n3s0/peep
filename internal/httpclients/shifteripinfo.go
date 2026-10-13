package httpclients

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
)

type ShifterIpInfo struct {
	Ip string `json:"ip"`
	City string `json:"city"`
	CityGeonameId uint `json:"city_geoname_id"`
	Region string `json:"region"`
	RegionCode string `json:"region_code"`
	RegionGeonameId uint `json:"region_geoname_id"`
	District string `json:"district"`
	DistrictCode uint `json:"district_code"`
	DistrictCodeId uint `json:"district_code_id"`
	Country string `json:"country"`
	CountryName string `json:"country_name"`
	CountryGeonameId uint `json:"country_geoname_id"`
	IsEu bool `json:"is_eu"`
	Continent string `json:"continent"`
	ContinentName string `json:"continent_name"`
	ContinentGeonameId uint `json:"continent_geoname_id"`
	Location string `json:"loc"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Postal string `json:"postal"`
	Timezone string `json:"timezone"`
	WeatherCode string `json:"weather_code"`
	Asn uint `json:"asn"`
	AsName string `json:"as_name"`
	Isp string `json:"isp"`
	Org string `json:"Org"`
	ConnectionType string `json:"connection_type"`
	UserType string `json:"user_type"`
	IsAnycast bool `json:"is_anycast"`
}

const (
	shifterIpInfoUrl string = "ip-info.com"
)

func GetShifterIpInfoData() (*ShifterIpInfo, error) {
	var jsonData ShifterIpInfo

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf("https://%s/json", shifterIpInfoUrl)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating new request: %v\n", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error executing the request: %v\n", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&jsonData)
	if err != nil {
		return nil, fmt.Errorf("Error with json decoder: %v\n", err)
	}

	return &jsonData, nil
}
