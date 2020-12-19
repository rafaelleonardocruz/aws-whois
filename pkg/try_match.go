package pkg

import "strings"

//TryMatchEc2 lookup insied ec2 instance list for a match with specified ip
func TryMatchEc2(ip string, ips []Instance) (resource string, address string, f bool) {
	found := false

	for _, addr := range ips {
		if strings.Contains(addr.IPAddress, ip) {
			return addr.InstanceID, addr.IPAddress, true
		}
	}

	return "", "", found
}

//TryMatchEip lookup insied EIP list for a match with specified ip
func TryMatchEip(ip string, ips []ElasticIp) (resource string, address string, f bool) {
	found := false

	for _, addr := range ips {
		if strings.Contains(addr.PublicIP, ip) {
			return addr.AllocationID, addr.PublicIP, true
		}
	}

	return "", "", found
}
