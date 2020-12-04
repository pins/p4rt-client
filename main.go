package main

import (
	"flag"
	"fmt"
	"github.com/pins/p4rt-client/internal/lib"
	"github.com/pins/p4rt-client/internal/logging"
)



func main() {
	flag.Parse()
	if *debug{
		logging.SetDebug(true)
	}
	if *logfile != ""{
		logging.SetOutputFile(logfile)
	}
	if logging.GetDebug(){
		message := "p4rt-client called with the following flags"
		logging.Debug(&message)
		flag.VisitAll(func(flag *flag.Flag ){
			val := fmt.Sprintf("%v",flag.Value)
			if val != flag.DefValue{
				flagMessage := fmt.Sprintf("		 %s:%s",flag.Name,val)
				logging.Debug(&flagMessage)
			}
		})
	}
	if *help {
		if *pushP4Info{
			lib.PushP4infoUsage()
		} else if *addRouterInt {
			lib.AddRouterIntUsage()
		} else if *addNeighbor {
			lib.AddNeighborUsage()
		} else if *addNextHop {
			lib.AddNextHopUsage()
		} else if *addIpV4Entry{
			lib.AddIpV4EntryUsage()
		} else if *createActionProfile {
			lib.CreateActionProfileUsage()
		} else if *addIpV4EntryWcmp {
			lib.AddIpV4EntryWcmpUsage()
		}else if *advanced{
			lib.ShowAdvancedUsage()
		} else{
			lib.BasicUsage()
		}
		return
	}
	//if multiple "actions are specified, they will be executed in this order
	//currently it is possible to create an interface, create a neighbor entry, create a nexthop label,
	//and then route a CIDR to the nexthop
	//creating a multipath group is not currently supported
	if *pushP4Info {
		lib.PushP4Info(serverAddressPort,p4info)
	}
	if *addRouterInt {
		lib.AddRouterIntEntry(serverAddressPort, routerInterfaceId, egressPort, routerIntPort, routerIntMAC, routerIntTableId, setMacPort)
	}
	if *addNeighbor {
		lib.AddNeighborEntry(serverAddressPort,routerInterfaceId, neighborIp, destMAC, neighborTable, setDestMac)
	}
	if *addNextHop {
		lib.AddNextHopEntry(serverAddressPort, nextHopId, neighborIp, routerInterfaceId, nextHopTable, nextHopAction)
	}
	if *addIpV4Entry {
		lib.AddIpV4TableEntry(serverAddressPort,vrfId, destNetwork, nextHopId, ipv4Table, setNextHopId)
	}
	if *createActionProfile {
		lib.CreateActionProfileEntry(serverAddressPort,  groupId, nextHops,actionProfileId, nextHopAction,)
	}
	if *addIpV4EntryWcmp {
		lib.AddIpV4WcmpEntry(serverAddressPort,vrfId, destNetwork, groupId,  ipv4Table, setWcmpId)
	}
}
