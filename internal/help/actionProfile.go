/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package help

func CreateActionProfileUsage() {
	usage := `
Usage:
NEXTHOP_LIST=$NEXTHOP_1:$WEIGHT_1,$NEXTHOP_2:$WEIGHT_2, ....
p4rt-client -addActionProfile \
            -server=$P4RUNTIME_ENDPOINT \
            -mpGroupId=$MULTIPATH_GROUP_NAME \ 
            -nextHopWeights=$NEXTHOP_LIST \
            -aProfileId=$ACTION_PROFILE_TABLE  (OPTIONAL)\
            -setNextHopAction=$SET_NEXTHOP (OPTIONAL)

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
           [ -aProfileId=33554499 ] \
           [ -setNextHopAction=16777221 ]
`
	log.Info(usage)
}
func DeleteActionProfileUsage() {
	usage := `
Usage:
NEXTHOP_LIST=$NEXTHOP_1:$WEIGHT_1,$NEXTHOP_2:$WEIGHT_2, ....
p4rt-client -delActionProfile \
            -server=$P4RUNTIME_ENDPOINT \
            -mpGroupId=$MULTIPATH_GROUP_NAME \ 
            -aProfileId=$ACTION_PROFILE_TABLE (OPTIONAL)

Fields:
-server:            string: ip address and listen port of P4Runtime service 
-mpGroupId          string: label to use for group of NextHops
-aProfileId:        uint32: table id associated with ingress.routing.wcmp_group_table from p4info.txt

e.g.
p4rt-client  -delActionProfile \
             -server=10.128.100.209:9559  \
             -mpGroupId=group1 \
           [ -aProfileId=33554499 ] 
`
	log.Info(usage)
}
func ShowAdvancedUsage() {
	usage := `
Multiple directives can be called in a single invocation for example to add a Router Interface,
create a Neighbor Entry, assign a NextHop ID and add a Route to that NextHop:
p4rt-client  -addRouterInt  -addNeighbor -addNextHop -addIpV4 \
             -server=10.128.100.209:9559 \
             -routerInterface=intf-eth1   \
             -routerPortId=1000  \
             -routerIntMAC=8c:ea:1b:17:64:0c  \
             -egressPort=Ethernet0   \
             -neighborName=192.168.2.2 \
             -routerInterface=intf-eth1 \
             -destMAC=00:07:43:4b:7f:50 \
             -neighborName=` + "`" + `macToIpV6 00:07:43:4b:7f:50 ` + "`" + ` \
             -nextHopId=bcmserver \
             -vrf=vrf-0 \
             -routedNetwork="172.16.2.0/24" \
           [ -routerTable=33554497 ] \
           [ -neighborTable=33554496 ]\
           [ -nextHopTable=33554498 ]\
           [ -ipv4table=33554500 ]\
           [ -setPortMac=16777218 ]\
           [ -setDestMacAction=16777217 ]\
           [ -setNextHopAction=16777219 ]\
           [ -setNextHop=16777221 ]
`
	log.Info(usage)

}
