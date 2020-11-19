package main

import (
	"flag"
	"fmt"
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/internal/lib"
	"github.com/pins/p4rt-client/internal/models"
	"github.com/pins/p4rt-client/internal/tables"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"strings"
)

var isMaster bool

func main() {

	serverAddressPort := flag.String("server", "", "address and port of p4rt server on swich")
	//#########################################
	pushP4Info := flag.Bool("pushP4Info", false, "push P4Info text file to switch")
	p4info := flag.String("p4info", "", "p4info text filename which describes p4 application")

	//#########################################
	//Router Interface Table
	addRouterInt := flag.Bool("addRouterInt", false, "add interface in router table")
	routerIntTableId := flag.Uint("routerTable", 0, "id for ingress.routing.router_interface_table from p4info.txt")
	routerInterfaceId := flag.String("routerInterface", "", "name to give to router interface")
	routerIntPort := flag.Uint("routerPortId", 1, "port number to assign to router interface port")
	routerIntMAC := flag.String("routerIntMAC", "", "MAC address to be used for router interface e.g. 00:00:00:11:22:dd")
	egressPort := flag.Uint("egressPort", 10000, "switch port to egress")
	setMacPort := flag.Uint("setPortMac", 0, "action id associated with action ingress.routing.set_port_and_src_mac ")

	//Neighbor Table
	addNeighbor := flag.Bool("addNeighbor", false, "create an entry in the neighbor table")
	neighborIp := flag.String("neighborIp", "", "ip address of next hop neighbor eg. 10.10.10.2")
	destMAC := flag.String("destMAC", "", "MAC address for neighbor IP e.g. 11:22:33:44:55:66")
	neighborTable := flag.Uint("neighborTable", 0, "id associated with ingress.routing.neighbor_table from p4info.txt")
	setDestMac := flag.Uint("setDestMacAction", 0, "id associated with ingress.routing.set_dst_mac from p4info.txt")
	//routerInterfaceId := flag.String("routerInterface","","name to give to router interface")

	//Next Hop Table
	addNextHop := flag.Bool("addNextHop", false, "add nexthop entry in nexthop table")
	nextHopTable := flag.Uint("nextHopTable", 0, "table id associated with ingress.routing.nexthop_table from p4info.txt")
	nextHopId := flag.String("nextHopId", "", "name to associate with next hop entry")
	nextHopAction := flag.Uint("setNextHopAction", 0, "action id associated with ingress.routing.set_nexthop_id from p4info.txt")
	//neighborIp    := flag.String("neighborIp","","ip address of next hop neighbor eg. 10.10.10.2")

	//IpV4Table
	addIpV4Entry := flag.Bool("addIpV4", false, "add routing entry in ipv4_table")
	vrfId := flag.String("vrf", "default", "name of vrf to use")
	destNetwork := flag.String("routedNetwork", "", "CIDR of network to route e.g. 1.2.3.4/8")
	ipv4Table := flag.Uint("ipv4table", 0, "id associated with ingress.routing.ipv4_table from p4info.txt")
	setNextHopId := flag.Uint("setNextHop", 0, "id associated with action ingress.routing.set_nexthop_id from p4info.txt")
	//nextHopId    := flag.String("nextHopId","","name to associate with next hop entry")

	//ActionProfileCreate
	createActionProfile := flag.Bool("addActionProfile", false, "add ActionProfileGroup with Members")
	actionProfileId := flag.Uint("aProfileId", 0, "table id associated with ingress.routing.wcmp_group_table from p4info.txt")
	nextHops := flag.String("nextHopWeights", "", "list of nexthops and weights to be used in action profile")
	groupId := flag.String("mpGroupId", "", "group id to use for MultiPath group")
	//setNextHopId := flag.Uint("setNextHop",0,"id associated with action ingress.routing.set_nexthop_id from p4info.txt")
	//destNetwork := flag.String("routedNetwork","","CIDR of network to route e.g. 1.2.3.4/8")

	addIpV4EntryWcmp := flag.Bool("addIpV4Wcmp", false, "create a routing entry with a wcmp path")
	//vrfId      := flag.String("vrf","default","name of vrf to use")
	//destNetwork := flag.String("routedNetwork","","CIDR of network to route e.g. 1.2.3.4/8")
	//ipv4Table   := flag.Uint("ipv4table",0,"id associated with ingress.routing.ipv4_table from p4info.txt")
	//groupId := flag.String("mpGroupId","", "group id to use for MultiPath group")
	setWcmpId := flag.Uint("setWcmpId", 0, "id associated with ingress.routing.set_wcmp_group_id")
	help := flag.Bool("help", false, "print usage help")
	flag.Parse()

	if *help {
		basicUsage()
		return
	}

	var conn *grpc.ClientConn
	var err error

	conn, err = grpc.Dial(*serverAddressPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial switch: %s", err)
	}

	defer conn.Close()
	client := p4.NewP4RuntimeClient(conn)
	p4client := &lib.P4rtClient{
		Client:   client,
		DeviceId: uint64(183807201),
	}
	lib.Init(p4client)

	if *pushP4Info {
		if becomeMaster(p4client) {
			lib.PushP4Info(p4client, p4info)
		}
	}
	if *addRouterInt {
		addRouterIntEntry(routerIntTableId, routerInterfaceId, egressPort, routerIntPort, routerIntMAC, setMacPort, p4client)
	}
	if *addNeighbor {
		addNeighborEntry(routerInterfaceId, neighborIp, destMAC, neighborTable, setDestMac, p4client)
	}

	if *addNextHop {
		addNextHopEntry(nextHopTable, nextHopId, nextHopAction, routerInterfaceId, neighborIp, p4client)
	}

	if *addIpV4Entry {
		addIpV4TableEntry(vrfId, destNetwork, nextHopId, ipv4Table, setNextHopId, p4client)
	}

	if *createActionProfile {
		createActionProfileEntry(*actionProfileId, nextHops, *nextHopAction, groupId, p4client)
	}

	if *addIpV4EntryWcmp {
		addIpV4WcmpEntry(vrfId, destNetwork, uint32(*ipv4Table), groupId, uint32(*setWcmpId), p4client)
	}
}
func becomeMaster(client *lib.P4rtClient) bool {
	if isMaster {
		return isMaster
	}
	client.SetMastership(&p4.Uint128{High: 0, Low: 1})
	isMaster = <-client.MasterChan
	return isMaster
}

func addIpV4WcmpEntry(vrf *string, dstNetwork *string, ipv4TableId uint32, wcmpGroupId *string, wcmpActionId uint32, p4client *lib.P4rtClient) {
	if becomeMaster(p4client) {
		wcmpEntry := &models.IPv4Config{
			VrfId:           vrf,
			DestinationCIDR: dstNetwork,
			IPv4TableId:     ipv4TableId,
			WcmpGroupId:     wcmpGroupId,
			SetWcmpGroupId:  wcmpActionId,
		}
		tables.Ipv4TableInsertWcmp(p4client, wcmpEntry)
	}
}

func addRouterIntEntry(routerIntTableId *uint, routerInterfaceId *string, egressPort *uint, routerIntPort *uint, routerIntMAC *string, setMacPort *uint, p4client *lib.P4rtClient) {
	if *routerIntTableId == 0 || *routerInterfaceId == "" || *egressPort == 10000 || *routerIntPort == 0 || *routerIntMAC == "" || *setMacPort == 0 {
		addRouterIntUsage()
		return
	}
	if becomeMaster(p4client) {
		l3Config := &models.L3Config{
			RouterInterfaceTableId: uint32(*routerIntTableId),
			RouterInterfaceId:      routerInterfaceId,
			EgressPort:             uint32(*egressPort),
			RouterInterfacePortId:  uint32(*routerIntPort),
			RouterInterfaceMAC:     routerIntMAC,
			SetMacAndPortId:        uint32(*setMacPort),
		}
		log.Println("Calling RouterTableInsert")
		tables.RounterTableInsert(p4client, l3Config)
	}
}

func addNeighborEntry(routerInterfaceId *string, neighborIp *string, destMAC *string, neighborTable *uint, setDestMac *uint, p4client *lib.P4rtClient) {
	if becomeMaster(p4client) {
		neighborConfig := &models.NeighborConfig{
			RouterInterfaceId:     routerInterfaceId,
			NeighborIP:            neighborIp,
			DestinationMac:        destMAC,
			NeighborTableId:       uint32(*neighborTable),
			NeighborTableActionId: uint32(*setDestMac),
		}
		log.Println("Calling NeighborTableInsert")
		tables.NeighborTableInsert(p4client, neighborConfig)
	}
}

func addNextHopEntry(nextHopTable *uint, nextHopId *string, nextHopAction *uint, routerInterfaceId *string, neighborIp *string, p4client *lib.P4rtClient) {
	if becomeMaster(p4client) {
		nexthopConfig := &models.NexthopConfig{
			NexthopTableId:    uint32(*nextHopTable),
			NexthopId:         nextHopId,
			SetNexthopId:      uint32(*nextHopAction),
			RouterInterfaceId: routerInterfaceId,
			NeighborIp:        neighborIp,
		}

		log.Println(*routerInterfaceId)
		log.Println(*neighborIp)
		log.Println("Calling NexthopTableInsert")
		tables.NextHopTableInsert(p4client, nexthopConfig)
	}
}

func addIpV4TableEntry(vrfId *string, destNetwork *string, nextHopId *string, ipv4Table *uint, setNextHopId *uint, p4client *lib.P4rtClient) {
	if becomeMaster(p4client) {
		ipv4Config := &models.IPv4Config{
			VrfId:           vrfId,
			DestinationCIDR: destNetwork,
			NexthopId:       nextHopId,
			IPv4TableId:     uint32(*ipv4Table),
			SetNexthopId:    uint32(*setNextHopId),
		}
		log.Println("Calling Ipv4TableInsert")
		tables.Ipv4TableInsert(p4client, ipv4Config)
	}
}
func createActionProfileEntry(actionProfileAction uint, nexthops *string, nexthopAction uint, groupId *string, p4client *lib.P4rtClient) {
	if becomeMaster(p4client) {
		profileMembers := []*models.ActionProfileGroupMember{}
		entries := strings.Split(*nexthops, ",")
		for i := 0; i < len(entries); i++ {
			fmt.Println(entries[i])
			subs := strings.Split(entries[i], ":")
			fmt.Printf("NextHop:%s Weight: %s\n", subs[0], subs[1])
			weight, _ := strconv.ParseInt(subs[1], 10, 32)
			profileMember := &models.ActionProfileGroupMember{
				Weight:    int32(weight),
				NexthopId: &subs[0],
			}
			profileMembers = append(profileMembers, profileMember)
		}
		actionProfileGroup := &models.ActionProfileGroup{
			ActionProfileId: uint32(actionProfileAction),
			SetNexthopId:    uint32(nexthopAction),
			GroupId:         groupId,
			Members:         profileMembers,
		}
		log.Println("Calling ActionProfileGroupInsert")
		tables.ActionProfileGroupInsert(p4client, actionProfileGroup)
	}
}

func basicUsage() {
	log.Println("p4rt-client can communication with p4runtime on a sonic based switch")
	log.Println("\n\nTo Set the p4info file:")
	pushP4infoUsage()
	log.Println("\n\nTo Configure Layer 3 routing:")
	addL3Usage()
	log.Println("\n\nTo see more help on arguments:\n./p4rt-client -h")
}
func addRouterIntUsage() {
	/*addRouterInt := flag.Bool("addRouterInt",false,"add interface in router table")
	routerIntTableId :=flag.Uint("routerTable",0,"id for ingress.routing.router_interface_table from p4info.txt")
	routerInterfaceId := flag.String("routerInterface","","name to give to router interface")
	routerIntPort  := flag.Uint("routerPortId",1,"port number to assign to router interface port")
	routerIntMAC := flag.String("routerIntMAC","","MAC address to be used for router interface e.g. 00:00:00:11:22:dd")
	egressPort := flag.Uint("egressPort",10000,"switch port to egress")
	setMacPort := flag.Uint("setPortMac",0,"action id associated with action ingress.routing.set_port_and_src_mac ")*/

	log.Println("failing")
}
func addL3Usage() {
	usage := `
Usage:
./p4rt-client -addL3 \
-server=$P4RT_LISTEN_ADDRESS  \
-routerTable=$ROUTER_TABLE_ID \
-routerInterface=$ROUTER_INTERFACE_NAME \
-routerPortId=$ROUTER_PORT_ID \
-routerIntMAC=$ROUTER_INTERFACE_MAC_ADDR \
-setPortMac=$SET_PORT_MAC_ACTION \
-nextHopTable=$NEXT_HOP_TABLE \
-nextHopId=$NEXT_HOP_NAME \
-setNextHopAction=$NEXT_HOP_ACTION_ID \
-neighborIp=$NEIGHBOR_IP \
-destMAC=$DEST_MAC \
-neighborTable=$NEIGHBOR_TABLE_ID \
-setDestMacAction=$SET_DST_MAC_ACTION_ID \
-vrf=$VRF_NAME \
-routedNetwork=$ROUTED_NETWORK_CIDR \
-ipv4table=$IPV_TABLE_ID \
-setNextHop=$SET_NEXT_HOP_ACTION_ID`
	log.Println(usage)
}

func pushP4infoUsage() {
	usage := `
Usage:
./p4rt-client -pushP4info -p4info=$P4_INFO_FILENAME`
	log.Println(usage)
}
