package cli

import (
	"log"
	"github.com/spf13/cobra"
	"github.com/rafaelleonardocruz/aws-whois/pkg"
)

func NewFindCmd() *cobra.Command {
	return &cobra.Command{
		Use: "find [ip-address]",
		Short: "find which resouce is using an IP address"
		Long: ``,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args string) {
			// WIP
		}

	}
}

func Find(ip string) {
	log.Println("validating if is a valid address, and if is an AWS address")

	dns, region, err := pkg.Resolver(ip)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	instances, err := pkg.GetInstances(region)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	resoure, address, found := pkg.TryMatchEc2(ip, instances)

	if found {
		log.Println("Your resource is %v at %v region", resource, region)
		os.Exit(0)
	}

	eip, err := pkg.GetElasticIps(region)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	resouce, address, found = pkg.TryMatchEip(ip, eip)

	if found {
		log.Println("Your resource is %v at %v region", resource, region)
		os.Exit(0)
	}

	log.Println("This IP wasn't found at thins AWS account")
	os.Exit(1)

}
