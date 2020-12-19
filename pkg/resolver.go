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
	return isPrivate, nil
}

func getReverseDnsName(ip string) ([]string, error) {
	hostNames, err := net.LookupAddr(ip)
	if err != nil {
		return nil, errors.New("Unable to resolve the address")
	}
	return hostNames, nil
}

func isAwsAddress(a string) (bool, error) {
	if strings.Contains(a, "aws") {
		return true, nil
	}

	return false, errors.New("This is not an AWS Address")
}

func getAwsRegion(hostname string) (region string, err error) {
	parser := strings.Split(hostname, ".")

	return string(parser[1]), nil
}

func resolver(ip string) (hostnames string, region string, err error) {
	if _, err := isValidAddress(ip); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	if _, err := isPrivateAddress(ip); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	dns, err := getReverseDnsName(ip)

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
