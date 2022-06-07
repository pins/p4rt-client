/*
 * Copyright (c) 2022-present Intel Corporation All Rights Reserved
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package helpUsage

func AddRouterIntUsage() {
	usage := `
Usage:
p4rt-client  -addRouterInt \
        -server=$P4RUNTIME_ENDPOINT \
        -routerInterface=$INTERFACE_NAME \
        -routerPortId=$ROUTER_PORT_ID \
        -routerIntMAC=$ROUTER_INTF_MAC  \
        -egressPort=$EGRESS_PORT \
        -routerTable=$ROUTER_TABLE (OPTIONAL) \
        -setPortMac=$SET_PORT_MAC  (OPTIONAL) 
Fields:
-server:            string: ip address and listen port of P4Runtime service 
-routerInterfaceId: string: unique name for virtual interface 
-routerIntPort:     uint32: unique port number to assign to virtual interface
-routerIntMAC:      string: MAC address to use for virtual interface - generally the same as used on all SONiC interfaces
-egressPort:        uint32: port number of physical port to use for egress
-routerTable:       uint32: table id for ingress.routing.router_interface_table for p4info.txt (OPTIONAL)
-setPortMac:        uint32: action id associated with ingress.routing.set_port_and_src_mac action in p4info.txt (OPTIONAL)
e.g.
 p4rt-client -addRouterInt \
             -server=10.128.100.209:9559 \
             -routerInterface=intf-eth1   \
             -routerPortId=1000  \
             -routerIntMAC=8c:ea:1b:17:64:0c  \
             -egressPort=Ethernet0   \
           [ -routerTable=33554497 ]  \
           [ -setPortMac=16777218 ]
`
	log.Info(usage)
}

func DelRouterIntUsage() {
	usage := `
Usage:
p4rt-client  -delRouterInt \
        -server=$P4RUNTIME_ENDPOINT \
        -routerInterface=$INTERFACE_NAME \
        -routerTable=$ROUTER_TABLE (OPTIONAL) \
        -setPortMac=$SET_PORT_MAC (OPTIONAL)
Fields:
-server:            string: ip address and listen port of P4Runtime service 
-routerInterfaceId: string: unique name for virtual interface 
-routerTable:       uint32: table id for ingress.routing.router_interface_table for p4info.txt (OPTIONAL)
-setPortMac:        uint32: action id associated with ingress.routing.set_port_and_src_mac action in p4info.txt (OPTIONAL)
e.g.
 p4rt-client -delRouterInt \
             -server=10.128.100.209:9559 \
             -routerInterface=intf-eth1   \
           [ -routerTable=33554497 ]  \
           [ -setPortMac=16777218 ]
`
	log.Info(usage)
}
