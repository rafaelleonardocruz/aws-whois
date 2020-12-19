package cli

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute(version string) {
	//var ipAddress string

	var rootCmd = &cobra.Command{
		Use:     "aws-whois",
		Short:   "aws-whois",
		Long:    `aws-whois - found which resource has a certain IP address`,
		Version: version,
	}

	//	rootCmd.PersistentFlags().StringVarP(&ipAddress, "ip-address", "ip", "", "IP Address that you would like to find")

	rootCmd.AddCommand(NewFindCmd())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
