package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	var ipFormatted string
	for i, v := range ip {
		if i == len(ip)-1 {
			ipFormatted += fmt.Sprintf("%v", v)
			break
		}
		ipFormatted += fmt.Sprintf("%v.", v)
	}
	return ipFormatted
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
