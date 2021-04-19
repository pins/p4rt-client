/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package lib
import (
	"log"
      )
func BasicUsage() {
	usage :=`
p4rt-client is used to configure the SAI layer of a SONiC based switch via the P4Runtime Service
Options:
p4rt-client	-pushP4Info          is used to push the p4info.txt file to the switch for command interpretation
p4rt-client	-addRouterInt        is used to create a virtual interface and map it to a physical interface
p4rt-client	-addNeighbor         is used to define an adjacent entity (switch, router, server etc..)
p4rt-client	-addNextHop          is used to create a NextHop label for a interface & neighbor combination
p4rt-client	-addIpV4             is used to create a route entry and point it to a NextHop
p4rt-client	-addActionProfile    is used to join several NextHop entries into one entity for (E|U)cmp pathing
p4rt-client	-addIpV4Wcmp         is used to create a route entry and point it to an (E|U)cmp path
p4rt-client	-help                prints this message

global options:
-debug              : generated detailed debugging info
-logfile=$LOG_FILE  : direct output to file instead of stdout

For help on an individual option include the option with -help e.g.
p4rt-client -pushP4Info -help

To see list of available arguments
p4rt-client -h

To see instructions for multiple Options in a single invocation:
p4rt-client -help -advanced
`
log.Println(usage)
}

func PushP4infoUsage() {
	usage := `
Usage:
./p4rt-client -pushP4info -p4info=$P4_INFO_FILENAME`
	log.Println(usage)
}

func AddRouterIntUsage(){
	usage :=`
Usage:
p4rt-client  -addRouterInt \
        -server=$P4RUNTIME_ENDPOINT \
        -routerInterface=$INTERFACE_NAME \
        -routerPortId=$ROUTER_PORT_ID \
        -routerIntMAC=$ROUTER_INTF_MAC  \
        -egressPort=$EGRESS_PORT \
        -routerTable=$ROUTER_TABLE \
        -setPortMac=$SET_PORT_MAC
Fields:
-server:            string: ip address and listen port of P4Runtime service 
-routerInterfaceId: string: unique name for virtual interface 
-routerIntPort:     uint32: unique port number to assign to virtual interface
-routerIntMAC:      string: MAC address to use for virtual interface - generally the same as used on all SONiC interfaces
-egressPort:        uint32: port number of physical port to use for egress
-routerTable:       uint32: table id for ingress.routing.router_interface_table for p4info.txt
-setPortMac:        uint32: action id associated with ingress.routing.set_port_and_src_mac action in p4info.txt
e.g.
 p4rt-client -addRouterInt \
             -server=10.128.100.209:9559 \
             -routerInterface=intf-eth1   \
             -routerPortId=1000  \
             -routerIntMAC=8c:ea:1b:17:64:0c  \
             -egressPort=125   \
             -routerTable=33554497  \
             -setPortMac=16777218
`
	log.Println(usage)
}

func AddNeighborUsage(){
	usage :=`
Usage:
p4rt-client -addNeighbor \
      -server=$P4RUNTIME_ENDPOINT \
      -neighborIp=$NEIGHBOR_IP \
      -destMAC=$DEST_MAC \
      -routerInterface=$INTERFACE_NAME \
      -neighborTable=$NEIGHBOR_TABLE \
      -setDestMacAction=$SET_DEST_MAC
Fields:
-server:            string: ip address and listen port of P4Runtime service 
-neighborIp         string: ip address of adjacent port
-destMAC            string: MAC address of adjacent port
-routerInterface    string: name of virtual interface used to pass traffic to this neighbor
-neighborTable      uint32: id associated with ingress.routing.neighbor_table from p4info.txt"
-setDestMacAction   uint32: id associated with ingress.routing.set_dst_mac from p4info.txt
e.g.
p4rt-client -addNeighbor \
            -server=10.128.100.209:9559 \
            -neighborIp=192.168.2.2 \
            -routerInterface=intf-eth1 \
            -destMAC=00:07:43:4b:7f:50 \
            -neighborTable=33554496 \
            -setDestMacAction=16777217 
`
    log.Println(usage)
}
func AddNextHopUsage(){
	usage:=`
Usage:
p4rt-client -addNextHop \
            -server=$P4RUNTIME_ENDPOINT  \
            -routerInterface=$INTERFACE_NAME \
            -neighborIp=` +"`"+`macToIpV6 $DEST_MAC`+"`"+` \
            -nextHopId=$NEXTHOP_NAME \
            -nextHopTable=$NEXTHOP_TABLE  \
            -setNextHopAction=$SET_NEXTHOP_ID
Fields:
-server:            string: ip address and listen port of P4Runtime service 
-routerInterface    string: name of the virtual interface to use
-neighborIp         string: this is the IPv6 link local address of the adjacent port
       macToIpV6 code is in the utils directory which converts a MAC address to IpV6 link local address
-nextHopId          string: unique name to identify the neighbor and interface combination
-nextHopTable       uint32: table id associated with ingress.routing.nexthop_table from p4info.txt
-setNextHopAction   uint32: action id associated with ingress.routing.set_nexthop_id from p4info.txt

e.g. Using IPv6 address directly:
p4rt-client  -addNextHop \
             -server=10.128.100.209:9559 \
             -nextHopTable=33554498 \
             -neighborIp=fe80::207:43ff:fe4b:7f50 \
             -nextHopId=bcmserver \
             -setNextHopAction=16777219 \
             -routerInterface=intf-eth1

e.g. Using macToIpV6 utility and Destination Mac:
p4rt-client  -addNextHop \
             -server=10.128.100.209:9559 \
             -nextHopTable=33554498 \
             -neighborIp=fe80::207:43ff:fe4b:7f50 \
             -neighborIp=`+"`"+`macToIpV6 00:07:43:4b:7f:50 `+"`"+` \
             -nextHopId=bcmserver \
             -setNextHopAction=16777219 \
             -routerInterface=intf-eth1
`
	log.Println(usage)

}
func AddIpV4EntryUsage(){
	usage:=`
Usage:
p4rt-client  -addIpV4 \
      -server=$P4RUNTIME_ENDPOINT \
      -vrf=$VRF_ID \
      -routedNetwork=$ROUTED_NETWORK \
      -nextHopId=$NEXTHOP_NAME \
      -ipv4table=$IPV4_ROUTE_TABLE \
      -setNextHop=$SET_NEXTHOP

Fields:
-server:            string: ip address and listen port of P4Runtime service 
-vrf:               string: name of VRF to use default: vrf-0
-routedNetwork      string: CIDR of network you are setting up a route for
-nextHopId          string: name of nextHop the routed packets should be sent to
-ipvtable           uint32: id associated with ingress.routing.ipv4_table from p4info.txt
-setNextHop         uint32: id associated with action ingress.routing.set_nexthop_id from p4info.txt

e.g.
p4rt-client -addIpV4 \
            -server=10.128.100.209:9559 \
            -vrf=vrf-0 \
            -routedNetwork="172.16.2.0/24" \
            -nextHopId=bcmserver \
            -ipv4table=33554500 \
            -setNextHop=16777221 

`
log.Println(usage)
}
func CreateActionProfileUsage(){
	usage:=`
Usage:
NEXTHOP_LIST=$NEXTHOP_1:$WEIGHT_1,$NEXTHOP_2:$WEIGHT_2, ....
p4rt-client -addActionProfile \
            -server=$P4RUNTIME_ENDPOINT \
            -mpGroupId=$MULTIPATH_GROUP_NAME \ 
            -nextHopWeights=$NEXTHOP_LIST \
            -aProfileId=$ACTION_PROFILE_TABLE \
            -setNextHopAction=$SET_NEXTHOP

Fields:
-server:            string: ip address and listen port of P4Runtime service 
-mpGroupId          string: label to use for group of NextHops
-nextHopWeights     string: comma separated list of NextHops and weights in the form of NextHop_1:Weight_1,NextHop_2:Weight_2...\
                    that make up the members of the multipath group
-aProfileId:        uint32: table id associated with ingress.routing.wcmp_group_table from p4info.txt
-setNextHopAction:  uint32: id associated with action ingress.routing.set_nexthop_id from p4info.txt
e.g.
nexthopList=bcmInter2:1,bcmInter3:1,bcmInter4:1,bcmInter5:1
p4rt-client  -addActionProfile \
             -server=10.128.100.209:9559  \
             -mpGroupId=group1 \
             -nextHopWeights=$nexthopList \
             -aProfileId=33554499 \
             -setNextHopAction=16777221  
`
	log.Println(usage)
}
func AddIpV4EntryWcmpUsage(){
	usage:=`
Usage: 
p4rt-client -addIpV4Wcmp \
            -server=$P4RUNTIME_ENDPOINT  \
            -vrf=VRF_ID \
            -routedNetwork=$ROUTED_NETWORK \
            -mpGroupId=$MULTIPATH_GROUP_NAME \
            -ipv4table=$IPV4_ROUTE_TABLE \
            -setWcmpId=$SET_WCMP_ACTION

Fields:
-server:            string: ip address and listen port of P4Runtime service 
-vrf:               string: name of VRF to use default: vrf-0
-routedNetwork      string: CIDR of network you are setting up a route for
-mpGroupId          string: name of the multi-path group routed packets should be sent to
-ipvtable           uint32: id associated with ingress.routing.ipv4_table from p4info.txt
-setWcmpId          uint32: id associated with ingress.routing.set_wcmp_group_id

e.g.
p4rt-client -addIpV4Wcmp \
            -server=10.128.100.209:9559  \
            -vrf=vrf-0 \
            -routedNetwork="172.16.1.0/24" \
            -mpGroupId=group1 \
            -ipv4table=33554500 \
            -setWcmpId=16777220
`
	log.Println(usage)
}
func ShowAdvancedUsage(){
	usage :=`
Multiple directives can be called in a single invocation for example to add a Router Interface,
create a Neighbor Entry, assign a NextHop ID and add a Route to that NextHop:
p4rt-client  -addRouterInt  -addNeighbor -addNextHop -addIpV4 \
             -server=10.128.100.209:9559 \
             -routerInterface=intf-eth1   \
             -routerPortId=1000  \
             -routerIntMAC=8c:ea:1b:17:64:0c  \
             -egressPort=125   \
             -neighborIp=192.168.2.2 \
             -routerInterface=intf-eth1 \
             -destMAC=00:07:43:4b:7f:50 \
             -neighborIp=`+"`"+`macToIpV6 00:07:43:4b:7f:50 `+"`"+` \
             -nextHopId=bcmserver \
             -vrf=vrf-0 \
             -routedNetwork="172.16.2.0/24" \
             -routerTable=33554497  \
             -neighborTable=33554496 \
             -nextHopTable=33554498 \
             -ipv4table=33554500 \
             -setPortMac=16777218 \
             -setDestMacAction=16777217 \
             -setNextHopAction=16777219 \
             -setNextHop=16777221 
`
log.Println(usage)

}
