package output

import (
	"fmt"

	"github.com/pterm/pterm"

	"github.com/n3s0/peep/internal/httpclients"
)

var (
	name string = "P33P"
)

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

	ipData := fmt.Sprintf("IP Address: %s\n", fetchedIpData.Ip)

	ipDataPanel := pterm.DefaultBox.WithTitle("IP Information").Sprint(ipData)

	locationData := fmt.Sprintf("City: %s\n", fetchedIpData.City)

	locationDataPanel := pterm.DefaultBox.WithTitle("Location").Sprint(locationData)

	panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{Data: ipDataPanel},
		{Data: locationDataPanel},
	}).Srender()

	spinner.Success("Data query complete... Output is below...")

	pterm.DefaultBox.WithTitle("p33p ip lookup").WithTitleTopLeft().WithRightPadding(0).WithBottomPadding(0).Println(panels)
}
