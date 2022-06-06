/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package lib

func AddNeighborUsage() {
	usage := `
Usage:
p4rt-client -addNeighbor \
      -server=$P4RUNTIME_ENDPOINT \
      -neighborName=$NEIGHBOR_IP \
      -destMAC=$DEST_MAC \
      -routerInterface=$INTERFACE_NAME \
      -neighborTable=$NEIGHBOR_TABLE  (OPTIONAL)\
      -setDestMacAction=$SET_DEST_MAC (OPTIONAL)
Fields:
-server:            string: ip address and listen port of P4Runtime service 
-neighborName         string: ip address of adjacent port
-destMAC            string: MAC address of adjacent port
-routerInterface    string: name of virtual interface used to pass traffic to this neighbor
-neighborTable      uint32: id associated with ingress.routing.neighbor_table from p4info.txt"
-setDestMacAction   uint32: id associated with ingress.routing.set_dst_mac from p4info.txt
e.g.
p4rt-client -addNeighbor \
            -server=10.128.100.209:9559 \
            -neighborName=192.168.2.2 \
            -routerInterface=intf-eth1 \
            -destMAC=00:07:43:4b:7f:50 \
          [ -neighborTable=33554496 ] \
          [ -setDestMacAction=16777217 ] 
`
	log.Info(usage)
}
func DelNeighborUsage() {
	usage := `
Usage:
p4rt-client -delNeighbor \
      -server=$P4RUNTIME_ENDPOINT \
      -routerInterface=$INTERFACE_NAME \
      -neighborName=$NEIGHBOR_IP \
      -neighborTable=$NEIGHBOR_TABLE  (OPTIONAL)
Fields:
-server:            string: ip address and listen port of P4Runtime service 
-routerInterface    string: name of virtual interface used to pass traffic to this neighbor
-neighborName         string: ip address of adjacent port
-neighborTable      uint32: id associated with ingress.routing.neighbor_table from p4info.txt" (OPTIONAL)

e.g.
p4rt-client -delNeighbor \
            -server=10.128.100.209:9559 \
            -routerInterface=intf-eth1 \
            -neighborName=192.168.2.2 \
          [ -neighborTable=33554496 ] 
`
	log.Info(usage)
}
