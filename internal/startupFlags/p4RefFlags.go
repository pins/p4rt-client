/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package startupFlags

import "flag"

var routerIntTableId = flag.Uint("routerTable", 33554497, "id for ingress.routing.router_interface_table from p4info.txt")
var neighborTable = flag.Uint("neighborTable", 33554496, "id associated with ingress.routing.neighbor_table from p4info.txt")
var nextHopTable = flag.Uint("nextHopTable", 33554498, "table id associated with ingress.routing.nexthop_table from p4info.txt")
var vrfTable = flag.Uint("vrfTable", 33554506, "table id associated with ")
var ipv4Table = flag.Uint("ipv4table", 33554500, "id associated with ingress.routing.ipv4_table from p4info.txt")

var setMacPort = flag.Uint("setPortMac", 16777218, "action id associated with action ingress.routing.set_port_and_src_mac ")
var setDestMac = flag.Uint("setDestMacAction", 16777217, "id associated with ingress.routing.set_dst_mac from p4info.txt")
var nextHopAction = flag.Uint("setNextHopAction", 16777219, "action id associated with ingress.routing.set_nexthop from p4info.txt")
var setNextHopId = flag.Uint("setNextHop", 16777221, "id associated with action ingress.routing.set_nexthop_id from p4info.txt")
var addVrfAction = flag.Uint("addVrf", 24742814, "id associated with add vrf")
var setWcmpId = flag.Uint("setWcmpId", 16777220, "id associated with ingress.routing.set_wcmp_group_id")
