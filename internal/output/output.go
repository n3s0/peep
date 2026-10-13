package output

import (
	"fmt"

	"github.com/pterm/pterm"

	"github.com/n3s0/peep/internal/httpclients"
)

var (
	name string = "P33P"
)

func GetIpOnly() {
	fetchedIpData, err := httpclients.GetShifterIpInfoData()
	if err != nil {
		fmt.Printf("Error getting IP info: %v", err)
	}

	fmt.Printf("%s\n", fetchedIpData.Ip)
}

func GetDetailedInfo() {
	spinner, _ := pterm.DefaultSpinner.
		WithSequence("▁", "▃", "▅", "▇", "▅", "▃").
		WithStyle(pterm.NewStyle(pterm.FgCyan)).
		Start("Querying IP info...")

	fetchedIpData, err := httpclients.GetShifterIpInfoData()
	if err != nil {
		spinner.Fail("Failed to get IP data...")
	}

	spinner.UpdateText("Finalizing IP report...")

	ipData := fmt.Sprintf("IP Address: %s\n\n", fetchedIpData.Ip)
	ipData += fmt.Sprintf("AS Name: %s\n", fetchedIpData.AsName)
	ipData += fmt.Sprintf("ASN: %d\n", fetchedIpData.Asn)
	ipData += fmt.Sprintf("ISP: %s\n", fetchedIpData.Isp)
	ipData += fmt.Sprintf("Org: %s\n", fetchedIpData.Org)
	ipData += fmt.Sprintf("Connection Type: %s\n", fetchedIpData.ConnectionType)
	ipData += fmt.Sprintf("User Type: %s\n", fetchedIpData.UserType)

	ipDataPanel := pterm.DefaultBox.WithTitle("IP Information").Sprint(ipData)

	locationData := fmt.Sprintf("City: %s\n", fetchedIpData.City)
	locationData += fmt.Sprintf("City Geoname ID: %d\n", fetchedIpData.CityGeonameId)
	locationData += fmt.Sprintf("Postal Code: %s\n\n", fetchedIpData.Postal)
	locationData += fmt.Sprintf("Region: %s\n", fetchedIpData.Region)
	locationData += fmt.Sprintf("Region Code: %s\n", fetchedIpData.RegionCode)
	locationData += fmt.Sprintf("Region Geoname ID: %d\n\n", fetchedIpData.RegionGeonameId)
	locationData += fmt.Sprintf("District: %s\n", fetchedIpData.District)
	locationData += fmt.Sprintf("District Code: %d\n", fetchedIpData.DistrictCode)
	locationData += fmt.Sprintf("District Code ID: %d\n\n", fetchedIpData.DistrictCodeId)
	locationData += fmt.Sprintf("Country: %s\n", fetchedIpData.Country)
	locationData += fmt.Sprintf("Country Name: %s\n\n", fetchedIpData.CountryName)
	locationData += fmt.Sprintf("Timezone: %s\n\n", fetchedIpData.Timezone)
	locationData += fmt.Sprintf("Latitude: %f\n", fetchedIpData.Latitude)
	locationData += fmt.Sprintf("Longitude: %f\n", fetchedIpData.Longitude)

	locationDataPanel := pterm.DefaultBox.WithTitle("Location").Sprint(locationData)

	panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{Data: ipDataPanel}},
		{{Data: locationDataPanel}},
	}).Srender()

	spinner.Success("Data query complete... Output is below...\n")

	pterm.DefaultBox.WithTitle("p33p ip lookup").WithTitleTopLeft().WithRightPadding(0).WithBottomPadding(0).Println(panels)
}
