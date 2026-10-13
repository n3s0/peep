package tui

import (
	"fmt"
	"os"

	"github.com/n3s0/peep/internal/httpclients"

	tea "github.com/charmbracelet/bubbletea"
)


type model struct {
	ipData *httpclients.ShifterIpInfo
}

const title = "P33P TUI"

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
			case "q":
				return m, tea.Quit
		}
	}
	return m, nil
}

/*
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
*/

func (m model) View() string {
	s := fmt.Sprintf("---- %s ----\n\n", title)
	
	s += fmt.Sprintf("[+] IP INFO [+]\n\n")

	s += fmt.Sprintf("IP Address: %s\n", m.ipData.Ip)
	s += fmt.Sprintf("ASN: %d\n", m.ipData.Asn)
	s += fmt.Sprintf("AS Name: %s\n", m.ipData.AsName)
	s += fmt.Sprintf("ISP: %s\n", m.ipData.Isp)
	s += fmt.Sprintf("Org: %s\n", m.ipData.Org)
	s += fmt.Sprintf("Connection Type: %s\n", m.ipData.ConnectionType)
	s += fmt.Sprintf("User Type: %s\n", m.ipData.UserType)

	if !m.ipData.IsAnycast {
		s += fmt.Sprintf("Is Anycast: No\n")
	} else {
		s += fmt.Sprintf("Is Anycast: Yes\n")
	}
	
	s += fmt.Sprintf("\n[!] LOCATION INFO [!]\n\n")
	
	s += fmt.Sprintf("City: %s\n", m.ipData.City)
	s += fmt.Sprintf("City Geoname ID: %d\n\n", m.ipData.CityGeonameId)

	s += fmt.Sprintf("Region: %s\n", m.ipData.Region)
	s += fmt.Sprintf("Region Code: %s\n", m.ipData.RegionCode)
	s += fmt.Sprintf("Region Geoname ID: %d\n\n", m.ipData.RegionGeonameId)

	s += fmt.Sprintf("District: %s\n", m.ipData.District)
	s += fmt.Sprintf("District Code: %d\n", m.ipData.DistrictCode)
	s += fmt.Sprintf("District Code ID: %d\n\n", m.ipData.DistrictCodeId)

	s += fmt.Sprintf("Country: %s\n", m.ipData.Country)
	s += fmt.Sprintf("Country Name: %s\n\n", m.ipData.CountryName)

	s += fmt.Sprintf("Latitude: %f\n", m.ipData.Latitude)
	s += fmt.Sprintf("Longitude: %f\n\n", m.ipData.Longitude)

	s += fmt.Sprintf("Postal Code: %s\n\n", m.ipData.Postal)

	s += fmt.Sprintf("Timezone: %s\n\n", m.ipData.Timezone)

	s += "Press 'r' (refresh) || 'q' (exit)\n"

	return s
}

func Execute() {
	fetchedIpData, err := httpclients.GetShifterIpInfoData()
	if err != nil {
		fmt.Printf("%v", err)
	}

	p := tea.NewProgram(model{ipData: fetchedIpData})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting program: %v", err)
		os.Exit(1)
	}
}
