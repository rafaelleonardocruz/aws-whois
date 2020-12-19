package pkg

import (
	"errors"
	"log"
	"net"
	"os"
	"strings"
)

func isValidAddress(ip string) (bool, error) {
	parsedIp := net.ParseIP(ip)

	if parsedIp == nil {
		return false, errors.New("Invalid IP address")
	}

	if strings.Count(ip, ":") >= 2 {
		return false, errors.New("It's a Ipv6 address and it's not supported yet")
	}

	log.Println("It's a valid address")
	return true, nil
}

func isPrivateAddress(ip string) (bool, error) {
	privateReservedAddress := []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"}

	isPrivate := false
	parsedIp := net.ParseIP(ip)

	for _, privateNet := range privateReservedAddress {
		_, pNet, _ := net.ParseCIDR(privateNet)

		if pNet.Contains(parsedIp) {
			isPrivate = true
		}
	}

	if isPrivate == true {
		return isPrivate, errors.New("This tool don't work with private address")
	}

	log.Println("It's a public ipv4 address")
	return isPrivate, nil
}

func getReverseDNSName(ip string) ([]string, error) {
	hostNames, err := net.LookupAddr(ip)
	if err != nil {
		return nil, errors.New("Unable to resolve the address")
	}

	log.Println("This address resovles to ", hostNames)
	return hostNames, nil
}

func isAwsAddress(a string) (bool, error) {
	if strings.Contains(a, "aws") {
		log.Println("This address is own by AWS")
		return true, nil
	}

	return false, errors.New("This is not an AWS Address")
}

func getAwsRegion(hostname string) (region string, err error) {
	regionList := []string{
		"us-east-2",
		"us-east-1",
		"us-west-1",
		"us-west-2",
		"af-south-1",
		"ap-east-1",
		"ap-south-1",
		"ap-northeast-3",
		"ap-northeast-2",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-northeast-1",
		"ca-central-1",
		"eu-central-1",
		"eu-west-1",
		"eu-west-2",
		"eu-south-1",
		"eu-west-3",
		"eu-north-1",
		"me-south-1",
		"sa-east-1",
	}

	var found string

	for _, r := range regionList {
		if strings.Contains(hostname, r) {
			found = r
		}
	}

	if len(region) > 0 {
		return found, nil
	}

	log.Println("This address has no region, we'll use your default")
	return "", nil
}

//Resolver get information about informed IP
func Resolver(ip string) (hostnames string, region string, err error) {
	if _, err := isValidAddress(ip); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	if _, err := isPrivateAddress(ip); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	dns, err := getReverseDNSName(ip)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	if _, err := isAwsAddress(dns[0]); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	reg, _ := getAwsRegion(dns[0])

	return dns[0], reg, nil
}
