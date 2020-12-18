package main

import (
	"net"
	"strconv"
)

// These functions need to be implemented
func cidrToMask(value string) string {
	result := value

	cidr, _ := strconv.Atoi(value)
	if (cidr >= 1 && cidr <= 32) {
		mask := net.CIDRMask(cidr, 32)
		result = net.IP(mask).String()
		return result	
	}
	result = "Invalid"
	return result
}

func maskToCidr(value string) string {
	valid := ipv4Validation(value)

	var result string
	result = value

	if valid ==false {
		result = "Invalid"
		return result
	}

	ip := net.IPMask(net.ParseIP(value).To4())

        add, _ := ip.Size()

	result = strconv.Itoa(add)

	return result
}

func ipv4Validation(value string) bool {
	ip := net.ParseIP(value)
	var status bool
	status = true

	if ip == nil {
		status = false
	}	

	return status
}
