/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package main

import "flag"

var version = flag.Bool("version", false, "show version of p4rt-client")
var debug = flag.Bool("debug", false, "enable debug logging")
var logfile = flag.String("logfile", "", "filename to use for logging")
var serverAddressPort = flag.String("server", "", "address and port of p4rt server on swich")

//#########################################
var pushP4Info = flag.Bool("pushP4Info", false, "push P4Info text file to switch")
var p4info = flag.String("p4info", "", "p4info text filename which describes p4 application")

//#########################################
//Router Interface Table
var addRouterInt = flag.Bool("addRouterInt", false, "add interface in router table")
var delRouterInt = flag.Bool("delRouterInt", false, "delete interface in router table")
var routerInterfaceId = flag.String("routerInterface", "", "name to give to router interface")
var routerIntPort = flag.Uint("routerPortId", 9999999, "port number to assign to router interface port")
var routerIntMAC = flag.String("routerIntMAC", "", "MAC address to be used for router interface e.g. 00:00:00:11:22:dd")
var egressPort = flag.String("egressPort", "", "switch port to egress")

//Neighbor Table
var addNeighbor = flag.Bool("addNeighbor", false, "create an entry in the neighbor table")
var delNeighbor = flag.Bool("delNeighbor", false, "delete an entry in the neighbor table")
var neighborName = flag.String("neighborName", "", "internal name for neighbor")
var destMAC = flag.String("destMAC", "", "MAC address for neighbor IP e.g. 11:22:33:44:55:66")

//routerInterfaceId = flag.String("routerInterface","","name to give to router interface")

//Next Hop Table
var addNextHop = flag.Bool("addNextHop", false, "add nexthop entry in nexthop table")
var delNextHop = flag.Bool("delNextHop", false, "delete nexthop entry in nexthop table")
var nextHopId = flag.String("nextHopId", "", "name to associate with next hop entry")

//var neighborName = flag.String("neighborName", "", "internal name for neighbor")

//add VRF
var addVrf = flag.Bool("addVRF", false, "add VRF entry")

//var vrfId = flag.String("vrf", "", "name of vrf to use")

//IpV4Table
var addIpV4Entry = flag.Bool("addIpV4", false, "add routing entry in ipv4_table")
var delIpV4Entry = flag.Bool("delIpV4", false, "delete routing entry in ipv4_table")
var vrfId = flag.String("vrf", "", "name of vrf to use")
var destNetwork = flag.String("routedNetwork", "", "CIDR of network to route e.g. 1.2.3.4/8")

//nextHopId    = flag.String("nextHopId","","name to associate with next hop entry")

//ActionProfileCreate
var createActionProfile = flag.Bool("addActionProfile", false, "add ActionProfileGroup with Members")
var delActionProfile = flag.Bool("delActionProfile", false, "delete ActionProfileGroup with Members")
var actionProfileId = flag.Uint("aProfileId", 33554499, "table id associated with ingress.routing.wcmp_group_table from p4info.txt")
var nextHops = flag.String("nextHopWeights", "", "list of nexthops and weights to be used in action profile")
var groupId = flag.String("mpGroupId", "", "group id to use for MultiPath group")

//setNextHopId = flag.Uint("setNextHop",0,"id associated with action ingress.routing.set_nexthop_id from p4info.txt")
//destNetwork = flag.String("routedNetwork","","CIDR of network to route e.g. 1.2.3.4/8")

var addIpV4EntryWcmp = flag.Bool("addIpV4Wcmp", false, "create a routing entry with a wcmp path")
var delIpV4EntryWcmp = flag.Bool("delIpV4Wcmp", false, "delete a routing entry with a wcmp path")

//vrfId      = flag.String("vrf","","name of vrf to use")
//destNetwork = flag.String("routedNetwork","","CIDR of network to route e.g. 1.2.3.4/8")
//ipv4Table   = flag.Uint("ipv4table",0,"id associated with ingress.routing.ipv4_table from p4info.txt")
//groupId = flag.String("mpGroupId","", "group id to use for MultiPath group")

//AddProfileMember currently not supported in SONiC - p4 SAI limitation
var addProfileMember = flag.Bool("addProfileMember", false, "add a \"member\" that can be assigned to a profile group")
var memberId = flag.Uint("memberId", 0, "unique member id ")

//var nextHopId = flag.String("nextHopId", "", "name to associate with next hop entry")
var profileId = flag.Uint("profileId", 299650760, "ingress.routing.wcmp_group_selector")

//var setNexthopId = flag.Uint()

var help = flag.Bool("help", false, "print usage help")
var advanced = flag.Bool("advanced", false, "show usage calling multiple entries at the same time")
