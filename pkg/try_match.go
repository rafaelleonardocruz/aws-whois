package pkg

import "strings"

//TryMatchEc2 lookup insied ec2 instance list for a match with specified ip
func TryMatchEc2(ip string, ips []Instance) (resource Instance, f bool) {
	found := false

	for _, addr := range ips {
		if strings.Contains(addr.IPAddress, ip) {
			return addr, true
		}
	}

	return Instance{"", ""}, found
}

//TryMatchEip lookup insied EIP list for a match with specified ip
func TryMatchEip(ip string, ips []ElasticIp) (resource ElasticIp, f bool) {
	found := false

	for _, addr := range ips {
		if strings.Contains(addr.PublicIP, ip) {
			return addr, true
		}
	}

	return ElasticIp{"", ""}, found
}
