package pkg

import "strings"

//TryMatch lookup insied AWS resources list for a match with specified ip
func TryMatch(ip string) (resource Resource, f bool) {
	found := false

	for _, addr := range ResourceList {
		if strings.Contains(addr.PublicIP, ip) {
			return addr, true
		}
	}

	return Resource{"", "", ""}, found
}
