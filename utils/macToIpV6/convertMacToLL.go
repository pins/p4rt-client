package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	macAddr := strings.TrimSpace(os.Args[1])
	octets := strings.Split(macAddr, ":")
	first,_ := strconv.ParseUint(octets[0],16,8)
	first = first ^ 2

	ipV6String := fmt.Sprintf("fe80::%x%s:%sff:fe%s:%s%s",first,octets[1],octets[2],octets[3],octets[4],octets[5])
	fmt.Println(ipV6String)
}




