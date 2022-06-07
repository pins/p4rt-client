/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package helpUsage

func AddIpV4EntryUsage() {
	usage := `
Usage:
p4rt-client  -addIpV4 \
      -server=$P4RUNTIME_ENDPOINT \
      -vrf=$VRF_ID (OPTIONAL)\
      -routedNetwork=$ROUTED_NETWORK \
      -nextHopId=$NEXTHOP_NAME \
      -ipv4table=$IPV4_ROUTE_TABLE  (OPTIONAL)\
      -setNextHop=$SET_NEXTHOP (OPTIONAL)

Fields:
-server:            string: ip address and listen port of P4Runtime service 
-vrf:               string: name of VRF to use default: vrf-0 can be ommitted for default vrf
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
          [ -ipv4table=33554500 ]\
          [ -setNextHop=16777221 ] 

`
	log.Info(usage)
}
func DelIpV4EntryUsage() {
	usage := `
Usage:
p4rt-client  -delIpV4 \
      -server=$P4RUNTIME_ENDPOINT \
      -vrf=$VRF_ID (OPTIONAL) \
      -routedNetwork=$ROUTED_NETWORK \
      -ipv4table=$IPV4_ROUTE_TABLE  (OPTIONAL)

Fields:
-server:            string: ip address and listen port of P4Runtime service 
-vrf:               string: name of VRF to use default: vrf-0, can be ommitted for default vrf
-routedNetwork      string: CIDR of network you are setting up a route for
-ipvtable           uint32: id associated with ingress.routing.ipv4_table from p4info.txt

e.g.
p4rt-client -delIpV4 \
            -server=10.128.100.209:9559 \
            -vrf=vrf-0 \
            -routedNetwork="172.16.2.0/24" \
          [ -ipv4table=33554500 ]

`
	log.Info(usage)
}

func AddIpV4EntryWcmpUsage() {
	usage := `
Usage: 
p4rt-client -addIpV4Wcmp \
            -server=$P4RUNTIME_ENDPOINT  \
            -vrf=VRF_ID \
            -routedNetwork=$ROUTED_NETWORK \
            -mpGroupId=$MULTIPATH_GROUP_NAME \
            -ipv4table=$IPV4_ROUTE_TABLE  (OPTIONAL)\
            -setWcmpId=$SET_WCMP_ACTION (OPTIONAL)

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
          [ -ipv4table=33554500 ]\
          [ -setWcmpId=16777220 ]
`
	log.Info(usage)
}
