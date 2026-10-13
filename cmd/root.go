package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"

	"github.com/n3s0/peep/internal/output"
	"github.com/n3s0/peep/internal/tui"
)

var (
	version bool
	detail bool
	simple bool
	startTui bool
	ipOnly bool
)

const appVersion string = "v0.1.0"

var rootCmd = &cobra.Command{
	Use: "peep",
	Short: "Pulls plublic IP information from various APIs",
	Args: cobra.MaximumNArgs(2),
	Long: `peep is a CLI/TUI application that provides information about your public IP.`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Printf("peep %s\n", appVersion)			
		}

		if startTui {
			tui.Execute()
		}

		if simple {
			fmt.Printf("This is simple output\n")
		}
		
		if detail {
			output.GetDetailedInfo()
		}
		
		if detail && simple{
			fmt.Printf("This is detailed output\n")
		}

		if ipOnly {
			output.GetIpOnly()
		}
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("Error running application: %v\n", err)
	}
	return nil
}

func init() {
	rootCmd.Flags().BoolVar(&version, "version", false, "Outputs the app version and quits")
	rootCmd.Flags().BoolVarP(&startTui, "tui", "t", false, "Launches the TUI interface")
	rootCmd.Flags().BoolVarP(&simple, "simple", "s", false, "Outputs a simple list of IP data")
	rootCmd.Flags().BoolVarP(&detail, "detail", "d", false, "Outputs a verbose list of IP data")
	rootCmd.Flags().BoolVarP(&ipOnly, "ip", "i", false, "Outputs only IP info")
}
