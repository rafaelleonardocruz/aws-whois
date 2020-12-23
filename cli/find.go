package cli

import (
	"log"
	"os"

	"github.com/rafaelleonardocruz/aws-whois/pkg"
	"github.com/spf13/cobra"
)

func NewFindCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find [ip-address]",
		Short: "find which resouce is using an IP address",
		Long:  ``,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ip := args[0]
			err := Find(ip)
			if err != nil {
				log.Println("This IP wasn't found at thins AWS account")
				os.Exit(1)
			}
		},
	}
	return cmd
}

//Find func lookup fo address trough supported resources
func Find(ip string) error {
	log.Println("validating if is a valid address, and if is an AWS address")

	_, region, err := pkg.Resolver(ip)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	err = pkg.GetInstances(region)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	err = pkg.GetElasticIps(region)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	Resource, Found := pkg.TryMatch(ip)

	if Found {
		log.Println("Your resource is a", Resource.ResourceName, "its ID is", Resource.ResourceID, "And its IP is ", Resource.PublicIP)
		return nil
	}

	log.Println("Address not found")
	return nil

}
