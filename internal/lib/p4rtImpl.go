package lib

import (
	"fmt"
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/internal/logging"
	"github.com/pins/p4rt-client/internal/models"
	"github.com/pins/p4rt-client/internal/tables"
	"log"
	"strconv"
	"strings"
)
var isMaster bool
var p4client *P4rtClient
/*
getP4Client: if p4client is nil - creates a new instance if not it simply returns existing instance
 */
func getP4Client(serverAddressPort *string)*P4rtClient{
	if p4client == nil {
		p4client = NewP4rtClient(serverAddressPort)
	}
	return p4client
}
/*
becomeMaster: if client is not yet master it initiates mastership otherwise it noops
 */
func becomeMaster(client *P4rtClient) bool {
	if isMaster {
		return isMaster
	}
	client.SetMastership(&p4.Uint128{High: 0, Low: 1})
	isMaster = <-client.MasterChan
	return isMaster
}
/*
PushP4Info: Pushes the P4Compiler generated info file to switch for table/action matching
 */
func PushP4Info(serverAddressPort *string,p4info *string){
	if *serverAddressPort==""||*p4info==""{
		PushP4infoUsage()
		message := fmt.Sprintf("PushP4Info called with ServerAddressPort : %s , p4info : %s",*serverAddressPort,*p4info)
		logging.Error(&message)
	}
	if becomeMaster(getP4Client(serverAddressPort)){
		pushP4Info(p4client,p4info)
	}

}
/*
AddRouterIntEntry: creates a sai/p4 virtual interface ref and binds it to a physical port on the switch
 */
func AddRouterIntEntry(serverAddressPort *string, routerInterfaceId *string, egressPort *uint, routerIntPort *uint,
	routerIntMAC *string, routerIntTableId *uint, setMacPort *uint) {
	if *serverAddressPort==""|| *routerIntTableId == 0 || *routerInterfaceId == "" ||
		*egressPort == 10000 || *routerIntPort == 0 || *routerIntMAC == "" || *setMacPort == 0 {
		AddRouterIntUsage()
		message := fmt.Sprintf("addRouterEntry called with serverAddressPort: %s,"+
			" routerInterface:%s, routerPortId:%d, routerIntMAC:%s"+
			" egressPort:%d routerTable:%d setPortMac:%d",
			*serverAddressPort,*routerInterfaceId,*routerIntPort,
			*routerIntMAC,*egressPort,*routerIntTableId,*setMacPort)
		logging.Error(&message)
		return
	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		l3Config := &models.L3Config{
			RouterInterfaceTableId: uint32(*routerIntTableId),
			RouterInterfaceId:      routerInterfaceId,
			EgressPort:             uint32(*egressPort),
			RouterInterfacePortId:  uint32(*routerIntPort),
			RouterInterfaceMAC:     routerIntMAC,
			SetMacAndPortId:        uint32(*setMacPort),
		}
		updates := tables.RouterTableInsert( l3Config)
		p4client.writeRequest(updates)
	}
}
/*
AddNeighborEntry: creates an entry with IpAddress and MAC of port adjacent to the referenced router interface
 */
func AddNeighborEntry(serverAddressPort *string, routerInterfaceId *string, neighborIp *string, destMAC *string,
	neighborTable *uint, setDestMac *uint) {
	if *serverAddressPort==""||*routerInterfaceId==""||*neighborIp==""||*destMAC==""||
		*neighborTable==0||*setDestMac==0{
		AddNeighborUsage()
		message := fmt.Sprintf("AddNeighborEntry called with serverAddressPort : %s routerInterfaceId : %s neighborIp : %s  "+
			"destMAC : %s neighborTable : %d setDestMac : %d",*serverAddressPort,*routerInterfaceId,*neighborIp,*destMAC,
			*neighborTable,*setDestMac)
		logging.Error(&message)
	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		neighborConfig := &models.NeighborConfig{
			RouterInterfaceId:     routerInterfaceId,
			NeighborIP:            neighborIp,
			DestinationMac:        destMAC,
			NeighborTableId:       uint32(*neighborTable),
			NeighborTableActionId: uint32(*setDestMac),
		}
		updates := tables.NeighborTableInsert( neighborConfig)
		p4client.writeRequest(updates)
	}
}

/*
AddNextHopEntry: creates a nexthop label for neighbor identified by router interface and adjacent ipV6 link local address
 */
func AddNextHopEntry(serverAddressPort *string, nextHopId *string, neighborIp *string,routerInterfaceId *string,
	nextHopTable *uint, nextHopAction *uint, ) {
	if *serverAddressPort==""||*nextHopTable==0||*nextHopId==""||*nextHopAction==0||
		*routerInterfaceId==""||*neighborIp==""{
		AddNextHopUsage()
		message := fmt.Sprintf("AddNextHopEntry called with serverAddressPort : %s routerInterfaceId : %s neighborIp : %s "+
			"nextHopId : %s nextHopTable : %d nextHopAction : %d", *serverAddressPort,*routerInterfaceId,*neighborIp,*nextHopId,
			*nextHopTable,*nextHopAction)
		logging.Error(&message)

	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		nexthopConfig := &models.NexthopConfig{
			NexthopTableId:    uint32(*nextHopTable),
			NexthopId:         nextHopId,
			SetNexthopId:      uint32(*nextHopAction),
			RouterInterfaceId: routerInterfaceId,
			NeighborIp:        neighborIp,
		}
		updates := tables.NextHopTableInsert( nexthopConfig)
		p4client.writeRequest(updates)
	}
}

/*
AddIpV4TableEntry: creates a route entry for a CIDR towards a previously labeled nexthop
 */
func AddIpV4TableEntry(serverAddressPort *string, vrfId *string, destNetwork *string, nextHopId *string,
	ipv4Table *uint, setNextHopId *uint) {
	if *serverAddressPort==""||*vrfId==""||*destNetwork==""||*nextHopId==""||*ipv4Table==0||*setNextHopId==0{
		AddIpV4EntryUsage()
		message := fmt.Sprintf("AddIpV4TableEntry called with serverAddressPort : %s, vrfId : %s destNetwork : %s " +
			"nextHopId : %s ipv4Table : %d setNextHopId : %d ",*serverAddressPort,*vrfId,*destNetwork,*nextHopId,
			*ipv4Table,*setNextHopId)
		logging.Error(&message)
	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		ipv4Config := &models.IPv4Config{
			VrfId:           vrfId,
			DestinationCIDR: destNetwork,
			NexthopId:       nextHopId,
			IPv4TableId:     uint32(*ipv4Table),
			SetNexthopId:    uint32(*setNextHopId),
		}
		updates := tables.Ipv4TableInsert( ipv4Config)
		p4client.writeRequest(updates)
	}
}

/*
CreateActionProfileEntry: creates a weighted group of nexthops in support of multipath routing
 */
func CreateActionProfileEntry(serverAddressPort *string, groupId *string, nexthops *string, actionProfileAction *uint,  nexthopAction *uint,) {
	if becomeMaster(getP4Client(serverAddressPort)) {
		if *serverAddressPort==""||*groupId==""||*nexthops==""||*actionProfileAction==0||*nexthopAction==0{
			message := fmt.Sprintf("CreateActionProfileEntry called with serverAddressPort : %s groupId : %s " +
				"nexthops : %s actionProfileAction : %d nexthopAction : %d",*serverAddressPort,*groupId,*nexthops,
				*actionProfileAction, *nexthopAction)
			logging.Error(&message)
		}
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
			ActionProfileId: uint32(*actionProfileAction),
			SetNexthopId:    uint32(*nexthopAction),
			GroupId:         groupId,
			Members:         profileMembers,
		}
		log.Println("Calling ActionProfileGroupInsert")
		updates := tables.ActionProfileGroupInsert(actionProfileGroup)
		p4client.writeRequest(updates)
	}
}

/*
AddIpV4WcmpEntry: creates a route entry for a CIDR towards a multipath group (ActionProfile)
 */
func AddIpV4WcmpEntry(serverAddressPort *string,vrf *string, dstNetwork *string, wcmpGroupId *string,
	ipv4TableId *uint, wcmpActionId *uint) {
	if *serverAddressPort==""||*vrf==""||*dstNetwork==""||*ipv4TableId==0||*wcmpGroupId==""||*wcmpActionId==0{
		AddIpV4EntryWcmpUsage()
		message := fmt.Sprintf("AddIpV4WcmpEntry called with " +
			"serverAddressPort : %s,dstNetwork : %s  wcmpGroupId : %s  ipv4TableId : %d wcmpActionId : %d ",
			*serverAddressPort,*dstNetwork,wcmpGroupId,ipv4TableId,wcmpActionId)
		logging.Error(&message)
	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		wcmpEntry := &models.IPv4Config{
			VrfId:           vrf,
			DestinationCIDR: dstNetwork,
			IPv4TableId:     uint32(*ipv4TableId),
			WcmpGroupId:     wcmpGroupId,
			SetWcmpGroupId:  uint32(*wcmpActionId),
		}
		updates := tables.Ipv4TableInsertWcmp(wcmpEntry)
		p4client.writeRequest(updates)
	}
}
func AddProfileMember(serverAddressPort *string,memberId *uint,
	                  nexthopId *string, profileId *uint,setNexthopId *uint){
	if becomeMaster(getP4Client(serverAddressPort)){
		profileMember := &models.ActionProfileGroupMember{
			MemberId: uint32(*memberId),
			NexthopId: nexthopId,
			ProfileId: uint32(*profileId),
			SetNexthopId: uint32(*setNexthopId),
		}
		updates := tables.ActionGroupMemberCreate(profileMember)
		p4client.writeRequest(updates)
	}
}
