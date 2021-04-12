package main

import (
	"flag"
	"fmt"
	"github.com/pins/p4rt-client/internal/lib"
	"github.com/pins/p4rt-client/internal/logging"
)



func main() {
	flag.Parse()
	if *debug {
		logging.SetDebug(true)
	}
	if *logfile != "" {
		logging.SetOutputFile(logfile)
	}
	if logging.GetDebug() {
		message := "p4rt-client called with the following flags"
		logging.Debug(&message)
		flag.VisitAll(func(flag *flag.Flag) {
			val := fmt.Sprintf("%v", flag.Value)
			if val != flag.DefValue {
				flagMessage := fmt.Sprintf("		 %s:%s", flag.Name, val)
				logging.Debug(&flagMessage)
			}
		})
	}
	if *help {
		if *pushP4Info {
			lib.PushP4infoUsage()
		} else if *addRouterInt {
			lib.AddRouterIntUsage()
		}else if *delRouterInt {
			lib.DelRouterIntUsage()
		} else if *addNeighbor {
			lib.AddNeighborUsage()
		}else if *delNeighbor {
			lib.DelNeighborUsage()
		} else if *addNextHop {
			lib.AddNextHopUsage()
		}else if *delNextHop {
			lib.DelNextHopUsage()
		} else if *addIpV4Entry {
			lib.AddIpV4EntryUsage()
		}else if *delIpV4Entry {
			lib.DelIpV4EntryUsage()
		} else if *createActionProfile {
			lib.CreateActionProfileUsage()
		}else if *delActionProfile {
			lib.DeleteActionProfileUsage()
		} else if *addIpV4EntryWcmp {
			lib.AddIpV4EntryWcmpUsage()
		}else if *delIpV4EntryWcmp {
			lib.DelIpV4EntryWcmpUsage()
		} else if *advanced {
			lib.ShowAdvancedUsage()
		} else {
			lib.BasicUsage()
		}
		return
	}
	//if multiple "actions are specified, they will be executed in this order
	//currently it is possible to create an interface, create a neighbor entry, create a nexthop label,
	//and then route a CIDR to the nexthop
	//creating a multipath group is not currently supported
	if *pushP4Info {
		lib.PushP4Info(serverAddressPort, p4info)
	}
	if *addRouterInt {
		lib.AddRouterIntEntry(serverAddressPort, routerInterfaceId, egressPort, routerIntPort, routerIntMAC, routerIntTableId, setMacPort)
	}
	if *addNeighbor {
		lib.AddNeighborEntry(serverAddressPort, routerInterfaceId, neighborName, destMAC, neighborTable, setDestMac)
	}
	if *addNextHop {
		lib.AddNextHopEntry(serverAddressPort, nextHopId, neighborName, routerInterfaceId, nextHopTable, nextHopAction)
	}
	if *addIpV4Entry {
		lib.AddIpV4TableEntry(serverAddressPort, vrfId, destNetwork, nextHopId, ipv4Table, setNextHopId)
	}
	if *createActionProfile {
		lib.CreateActionProfileEntry(serverAddressPort, groupId, nextHops, actionProfileId, nextHopAction)
	}
	if *addIpV4EntryWcmp {
		lib.AddIpV4WcmpEntry(serverAddressPort, vrfId, destNetwork, groupId, ipv4Table, setWcmpId)
	}
	if *addProfileMember {
		lib.AddProfileMember(serverAddressPort, memberId, nextHopId, profileId, setNextHopId)
	}
	/*
	In batched deletes the order of operations should be reverse
	 */
	if *delIpV4EntryWcmp {
		lib.DelIpV4WcmpEntry(serverAddressPort, vrfId, destNetwork, groupId, ipv4Table, setWcmpId)
	}
	if *delActionProfile {
		lib.DeleteActionProfileEntry(serverAddressPort, groupId, nextHops, actionProfileId, nextHopAction)
	}
	if *delIpV4Entry {
		lib.DelIpV4TableEntry(serverAddressPort, vrfId, destNetwork, nextHopId, ipv4Table, setNextHopId)
	}
	if *delNextHop {
		lib.DelNextHopEntry(serverAddressPort, nextHopId, neighborName, routerInterfaceId, nextHopTable, nextHopAction)
	}
	if *delNeighbor {
		lib.DelNeighborEntry(serverAddressPort, routerInterfaceId, neighborName, destMAC, neighborTable, setDestMac)
	}
	if *delRouterInt {
		lib.DeleteRouterIntEntry(serverAddressPort, routerInterfaceId, egressPort, routerIntPort, routerIntMAC, routerIntTableId, setMacPort)
	}
}
