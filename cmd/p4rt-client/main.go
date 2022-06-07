/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package main

import (
	"flag"
	"fmt"

	h "github.com/pins/p4rt-client/internal/helpUsage"

	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/pins/p4rt-client/pkg/libp4rt"
)

var log = logging.GetLogger("p4rt-client")

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		*help = true
	}
	if *version {
		message := fmt.Sprintf("p4rt-client version: %s target: %s ", p4rt_client_version, p4rt_client_summary)
		log.Info(&message)
		return
	}
	if *help {
		if *pushP4Info {
			h.PushP4Usage()
		} else if *addRouterInt {
			h.AddRouterIntUsage()
		} else if *delRouterInt {
			h.DelRouterIntUsage()
		} else if *addNeighbor {
			h.AddNeighborUsage()
		} else if *delNeighbor {
			h.DelNeighborUsage()
		} else if *addNextHop {
			h.AddNextHopUsage()
		} else if *delNextHop {
			h.DelNextHopUsage()
		} else if *addIpV4Entry {
			h.AddIpV4EntryUsage()
		} else if *delIpV4Entry {
			h.DelIpV4EntryUsage()
		} else if *createActionProfile {
			h.CreateActionProfileUsage()
		} else if *delActionProfile {
			h.DeleteActionProfileUsage()
		} else if *addIpV4EntryWcmp {
			h.AddIpV4EntryWcmpUsage()
		} else if *advanced {
			h.ShowAdvancedUsage()
		} else {
			h.BasicUsage()
		}
		return
	}
	//if multiple "actions are specified, they will be executed in this order
	//currently it is possible to create an interface, create a neighbor entry, create a nexthop label,
	//and then route a CIDR to the nexthop
	//creating a multipath group is not currently supported
	if *pushP4Info {
		err := libp4rt.PushP4Info(serverAddressPort, p4info)
		if err != nil {
			log.Error(err)
			h.PushP4Usage()
		}
	}
	if *addRouterInt {
		err := libp4rt.AddRouterIntEntry(serverAddressPort, routerInterfaceId, egressPort, routerIntPort, routerIntMAC, routerIntTableId, setMacPort)
		if err != nil {
			log.Error(err)
			h.AddRouterIntUsage()
		}
	}
	if *addNeighbor {
		err := libp4rt.AddNeighborEntry(serverAddressPort, routerInterfaceId, neighborName, destMAC, neighborTable, setDestMac)
		if err != nil {
			log.Error(err)
			h.AddNeighborUsage()
		}
	}
	if *addNextHop {
		err := libp4rt.AddNextHopEntry(serverAddressPort, nextHopId, neighborName, routerInterfaceId, nextHopTable, nextHopAction)
		if err != nil {
			log.Error(err)
			h.AddNextHopUsage()
		}
	}
	if *addVrf {
		libp4rt.AddVRF(serverAddressPort, vrfId, vrfTable, addVrfAction)
	}
	if *addIpV4Entry {
		err := libp4rt.AddIpV4TableEntry(serverAddressPort, vrfId, destNetwork, nextHopId, ipv4Table, setNextHopId)
		if err != nil {
			log.Error(err)
			h.AddIpV4EntryUsage()
		}
	}
	if *createActionProfile {
		err := libp4rt.CreateActionProfileEntry(serverAddressPort, groupId, nextHops, actionProfileId, setNextHopId)
		if err != nil {
			log.Error(err)
			h.CreateActionProfileUsage()
		}
	}
	if *addIpV4EntryWcmp {
		err := libp4rt.AddIpV4WcmpEntry(serverAddressPort, vrfId, destNetwork, groupId, ipv4Table, setWcmpId)
		if err != nil {
			log.Error(err)
			h.AddIpV4EntryWcmpUsage()
		}
	}
	if *addProfileMember {
		err := libp4rt.AddProfileMember(serverAddressPort, memberId, nextHopId, profileId, setNextHopId)
		if err != nil {
			log.Error(err)
		}
	}
	/*
		In batched deletes the order of operations should be reverse
	*/
	if *delActionProfile {
		err := libp4rt.DeleteActionProfileEntry(serverAddressPort, groupId, actionProfileId)
		if err != nil {
			log.Error(err)
			h.DeleteActionProfileUsage()
		}
	}
	if *delIpV4Entry {
		err := libp4rt.DelIpV4TableEntry(serverAddressPort, vrfId, destNetwork, ipv4Table)
		if err != nil {
			log.Error(err)
			h.DelIpV4EntryUsage()
		}
	}
	if *delNextHop {
		err := libp4rt.DelNextHopEntry(serverAddressPort, nextHopId, nextHopTable)
		if err != nil {
			log.Error(err)
			h.DelNextHopUsage()
		}
	}
	if *delNeighbor {
		err := libp4rt.DelNeighborEntry(serverAddressPort, routerInterfaceId, neighborName, neighborTable)
		if err != nil {
			log.Error(err)
			h.DelNeighborUsage()
		}
	}
	if *delRouterInt {
		err := libp4rt.DeleteRouterIntEntry(serverAddressPort, routerInterfaceId, routerIntTableId)
		if err != nil {
			log.Error(err)
			h.DelRouterIntUsage()
		}
	}
}
