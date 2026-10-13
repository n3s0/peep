package tui

import (
	"fmt"
	"os"

	"github.com/n3s0/peep/internal/httpclients"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	ipData *httpclients.ShifterIpInfo
}


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

func (m model) View() string {
	
	const (
		title string = "P33P TUI"
		subTitle string = "A simple TUI for viewing information about your public IP."
		ipHeader string = "IP INFORMATION"
		locationHeader string = "LOCATION"
		boxWidth int = 80
	)

	interiorWidth := boxWidth - 10
	
	rootBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FFFFFFFF")).
		Padding(1, 2).
		Margin(1, 1).
		Width(boxWidth).
		Align(lipgloss.Left)
	
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00000000")).
		Background(lipgloss.Color("#FFFFFFFF")).
		Width(interiorWidth).
		Align(lipgloss.Left)

	subTitleStyle := lipgloss.NewStyle().
		Width(interiorWidth).
		Align(lipgloss.Left)
	
	ipDataBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FFFFFFF")).
		Padding(1, 2).
		Margin(1, 1).
		Width(interiorWidth).
		Align(lipgloss.Left)
	
	ipDataHeader := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00000000")).
		Background(lipgloss.Color("#FFFFFFFF")).
		Width(interiorWidth - 5).
		Align(lipgloss.Left)

	locationDataBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FFFFFFF")).
		Padding(1, 2).
		Margin(1, 1).
		Width(interiorWidth).
		Align(lipgloss.Left)
	
	locationDataHeader := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00000000")).
		Background(lipgloss.Color("#FFFFFFFF")).
		Width(interiorWidth - 5).
		Align(lipgloss.Left)
	
	appTitle := titleStyle.Render(title) + "\n\n"

	appSubTitle := subTitleStyle.Render(subTitle) + "\n"

	ipDataBoxHeader := ipDataHeader.Render(ipHeader) + "\n"
	
	ipDataBoxBody := fmt.Sprintf("IP Address: %s\n", m.ipData.Ip)
	ipDataBoxBody += fmt.Sprintf("AS Name: %s\n", m.ipData.AsName)
	ipDataBoxBody += fmt.Sprintf("ASN: %d\n", m.ipData.Asn)
	ipDataBoxBody += fmt.Sprintf("ISP: %s\n", m.ipData.Isp)
	ipDataBoxBody += fmt.Sprintf("Org: %s\n", m.ipData.Org)
	ipDataBoxBody += fmt.Sprintf("Connection Type: %s\n", m.ipData.ConnectionType)
	ipDataBoxBody += fmt.Sprintf("User Type: %s\n", m.ipData.UserType)

	if !m.ipData.IsAnycast {
		ipDataBoxBody += fmt.Sprintf("Is Anycast: No\n")
	} else {
		ipDataBoxBody += fmt.Sprintf("Is Anycast: Yes\n")
	}

	ipDataBoxContent := lipgloss.JoinVertical(lipgloss.Left, ipDataBoxHeader, ipDataBoxBody)
	ipDataBox := ipDataBoxStyle.Render(ipDataBoxContent)

	locationDataBoxHeader := locationDataHeader.Render(locationHeader) + "\n"
	
	locationDataBoxBody := fmt.Sprintf("City: %s\n", m.ipData.City)
	locationDataBoxBody += fmt.Sprintf("City Geoname ID: %d\n", m.ipData.CityGeonameId)
	locationDataBoxBody += fmt.Sprintf("Postal Code: %s\n\n", m.ipData.Postal)
	locationDataBoxBody += fmt.Sprintf("Region: %s\n", m.ipData.Region)
	locationDataBoxBody += fmt.Sprintf("Region Code: %s\n", m.ipData.RegionCode)
	locationDataBoxBody += fmt.Sprintf("Region Geoname ID: %d\n\n", m.ipData.RegionGeonameId)
	locationDataBoxBody += fmt.Sprintf("District: %s\n", m.ipData.District)
	locationDataBoxBody += fmt.Sprintf("District Code: %d\n", m.ipData.DistrictCode)
	locationDataBoxBody += fmt.Sprintf("District Code ID: %d\n\n", m.ipData.DistrictCodeId)
	locationDataBoxBody += fmt.Sprintf("Country: %s\n", m.ipData.Country)
	locationDataBoxBody += fmt.Sprintf("Country Name: %s\n\n", m.ipData.CountryName)
	locationDataBoxBody += fmt.Sprintf("Timezone: %s\n\n", m.ipData.Timezone)
	locationDataBoxBody += fmt.Sprintf("Latitude: %f\n", m.ipData.Latitude)
	locationDataBoxBody += fmt.Sprintf("Longitude: %f\n", m.ipData.Longitude)

	locationDataBoxContent := lipgloss.JoinVertical(lipgloss.Left, locationDataBoxHeader, locationDataBoxBody)
	
	locationDataBox := locationDataBoxStyle.Render(locationDataBoxContent)
	
	subLayout := lipgloss.JoinVertical(lipgloss.Top, appTitle, appSubTitle, ipDataBox, locationDataBox)

	return rootBoxStyle.Render(subLayout)
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
