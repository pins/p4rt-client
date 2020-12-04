package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//This takes a single argument in the form of 00:11:22:33:44:55 (hex) and outputs the IPv6 Link Local address
func main() {
	args := os.Args
	if len(args) != 2{
		usage()
	}
	macAddr := strings.TrimSpace(os.Args[1])
	octets := strings.Split(macAddr, ":")
	if len(octets)!=6{
		usage()
	}
	first, _ := strconv.ParseUint(octets[0], 16, 8)
	first = first ^ 2
	ipV6String := fmt.Sprintf("fe80::%x%s:%sff:fe%s:%s%s", first, octets[1], octets[2], octets[3], octets[4], octets[5])
	fmt.Println(ipV6String)
}
func usage(){
message :=`
Usage:
macToIpV6 00:11:22:33:44:55 
returns an IPv6 address
fe80::211:22ff:fe33:4455 in this case
`
log.Fatalln(message)
}
