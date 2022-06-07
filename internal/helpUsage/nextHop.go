/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package helpUsage

func AddNextHopUsage() {
	usage := `
Usage:
p4rt-client -addNextHop \
            -server=$P4RUNTIME_ENDPOINT  \
            -routerInterface=$INTERFACE_NAME \
            -neighborName=` + "`" + `macToIpV6 $DEST_MAC` + "`" + ` \
            -nextHopId=$NEXTHOP_NAME \
            -nextHopTable=$NEXTHOP_TABLE   (OPTIONAL)\
            -setNextHopAction=$SET_NEXTHOP_ID (OPTIONAL)
Fields:
-server:            string: ip address and listen port of P4Runtime service 
-routerInterface    string: name of the virtual interface to use
-neighborName         string: this is the IPv6 link local address of the adjacent port
       macToIpV6 code is in the utils directory which converts a MAC address to IpV6 link local address
-nextHopId          string: unique name to identify the neighbor and interface combination
-nextHopTable       uint32: table id associated with ingress.routing.nexthop_table from p4info.txt
-setNextHopAction   uint32: action id associated with ingress.routing.set_nexthop_id from p4info.txt

e.g. Using IPv6 address directly:
p4rt-client  -addNextHop \
             -server=10.128.100.209:9559 \
             -neighborName=fe80::207:43ff:fe4b:7f50 \
             -routerInterface=intf-eth1 \
             -nextHopId=bcmserver \
           [ -setNextHopAction=16777219 ] \
           [ -nextHopTable=33554498 ]

e.g. Using macToIpV6 utility and Destination Mac:
p4rt-client  -addNextHop \
             -server=10.128.100.209:9559 \
             -neighborName=` + "`" + `macToIpV6 00:07:43:4b:7f:50 ` + "`" + ` \
             -routerInterface=intf-eth1 \
             -nextHopId=bcmserver \
           [ -nextHopTable=33554498 ]\
           [ -setNextHopAction=16777219 ]
`
	log.Info(usage)

}
func DelNextHopUsage() {
	usage := `
Usage:
p4rt-client -delNextHop \
            -server=$P4RUNTIME_ENDPOINT  \
            -nextHopId=$NEXTHOP_NAME \
            -nextHopTable=$NEXTHOP_TABLE   (OPTIONAL)
Fields:
-server:            string: ip address and listen port of P4Runtime service 
-nextHopId          string: unique name to identify the neighbor and interface combination
-nextHopTable       uint32: table id associated with ingress.routing.nexthop_table from p4info.txt

e.g. :
p4rt-client  -delNextHop \
             -server=10.128.100.209:9559 \
             -nextHopId=bcmserver \
           [ -nextHopTable=33554498 ] 

`
	log.Info(usage)

}
