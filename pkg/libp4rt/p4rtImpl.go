/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package libp4rt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/pkg/models"
	"github.com/pins/p4rt-client/pkg/p4rtFormatter"
)

var isMaster bool
var p4client *P4rtClient

/*
getP4Client: if p4client is nil - creates a new instance if not it simply returns existing instance
*/
func getP4Client(serverAddressPort *string) *P4rtClient {
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
func PushP4Info(serverAddressPort *string, p4info *string) error {
	if *serverAddressPort == "" || *p4info == "" {
		message := fmt.Sprintf("PushP4Info called with ServerAddressPort : %s , p4info : %s", *serverAddressPort, *p4info)
		return errors.New(message)
	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		pushP4Info(p4client, p4info)
	}
	return nil
}

/*
AddRouterIntEntry: creates a sai/p4 virtual interface ref and binds it to a physical port on the switch
*/
func AddRouterIntEntry(serverAddressPort *string, routerInterfaceId *string, egressPort *string, routerIntPort *uint,
	routerIntMAC *string, routerIntTableId *uint, setMacPort *uint) error {
	if *serverAddressPort == "" || *routerInterfaceId == "" ||
		*egressPort == "" || *routerIntPort == 9999999 || *routerIntMAC == "" {
		message := fmt.Sprintf("addRouterEntry called with serverAddressPort: %s routerInterface:%s routerPortId:%d routerIntMAC:%s egressPort:%d routerTable:%d setPortMac:%d",
			*serverAddressPort, *routerInterfaceId, *routerIntPort, *routerIntMAC, *egressPort, *routerIntTableId, *setMacPort)
		return errors.New(message)
	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		l3Config := &models.L3Config{
			RouterInterfaceTableId: uint32(*routerIntTableId),
			RouterInterfaceId:      routerInterfaceId,
			EgressPort:             egressPort,
			RouterInterfacePortId:  uint32(*routerIntPort),
			RouterInterfaceMAC:     routerIntMAC,
			SetMacAndPortId:        uint32(*setMacPort),
		}
		updates := p4rtFormatter.RouterTableInsert(l3Config)
		p4client.writeRequest(updates)
	}
	return nil
}

/*
DeleteRouterIntEntry: creates a sai/p4 virtual interface ref and binds it to a physical port on the switch
*/
func DeleteRouterIntEntry(serverAddressPort *string, routerInterfaceId *string, routerIntTableId *uint) error {
	if *serverAddressPort == "" || *routerInterfaceId == "" {
		message := fmt.Sprintf("deleteRouterEntry called with serverAddressPort: %s, routerInterface:%s",
			*serverAddressPort, *routerInterfaceId)
		return errors.New(message)
	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		l3Config := &models.L3Config{
			RouterInterfaceTableId: uint32(*routerIntTableId),
			RouterInterfaceId:      routerInterfaceId,
		}
		updates := p4rtFormatter.RouterTableDelete(l3Config)
		p4client.writeRequest(updates)
	}
	return nil
}

/*
AddNeighborEntry: creates an entry with IpAddress and MAC of port adjacent to the referenced router interface
*/
func AddNeighborEntry(serverAddressPort *string, routerInterfaceId *string, neighborName *string, destMAC *string,
	neighborTable *uint, setDestMac *uint) error {
	if *serverAddressPort == "" || *routerInterfaceId == "" || *neighborName == "" || *destMAC == "" || *neighborTable == 0 || *setDestMac == 0 {
		message := fmt.Sprintf("AddNeighborEntry called with serverAddressPort : %s routerInterfaceId : %s neighborName : %s  "+
			"destMAC : %s neighborTable : %d setDestMac : %d", *serverAddressPort, *routerInterfaceId, *neighborName, *destMAC,
			*neighborTable, *setDestMac)
		return errors.New(message)
	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		neighborConfig := &models.NeighborConfig{
			RouterInterfaceId:     routerInterfaceId,
			NeighborName:          neighborName,
			DestinationMac:        destMAC,
			NeighborTableId:       uint32(*neighborTable),
			NeighborTableActionId: uint32(*setDestMac),
		}

		updates := p4rtFormatter.NeighborTableInsert(neighborConfig)
		p4client.writeRequest(updates)
	}
	return nil
}

/*
DelNeighborEntry: creates an entry with IpAddress and MAC of port adjacent to the referenced router interface
*/
func DelNeighborEntry(serverAddressPort *string, routerInterfaceId *string, neighborName *string, neighborTable *uint) error {
	if *serverAddressPort == "" || *routerInterfaceId == "" || *neighborName == "" || *neighborTable == 0 {
		message := fmt.Sprintf("DelNeighborEntry called with serverAddressPort : %s  routerInterfaceId : %sneighborName : %s  "+"neighborTable : %d ",
			*serverAddressPort, *routerInterfaceId, *neighborName, *neighborTable)
		return errors.New(message)
	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		neighborConfig := &models.NeighborConfig{
			RouterInterfaceId: routerInterfaceId,
			NeighborName:      neighborName,
			NeighborTableId:   uint32(*neighborTable),
		}
		updates := p4rtFormatter.NeighborTableDelete(neighborConfig)
		p4client.writeRequest(updates)
	}
	return nil
}

/*
AddNextHopEntry: creates a nexthop label for neighbor identified by router interface and adjacent ipV6 link local address
*/
func AddNextHopEntry(serverAddressPort *string, nextHopId *string, neighborName *string, routerInterfaceId *string,
	nextHopTable *uint, nextHopAction *uint) error {
	if *serverAddressPort == "" || *nextHopTable == 0 || *nextHopId == "" || *nextHopAction == 0 ||
		*routerInterfaceId == "" || *neighborName == "" {
		message := fmt.Sprintf("AddNextHopEntry called with serverAddressPort : %s routerInterfaceId : %s neighborName : %s "+
			"nextHopId : %s nextHopTable : %d nextHopAction : %d", *serverAddressPort, *routerInterfaceId, *neighborName, *nextHopId,
			*nextHopTable, *nextHopAction)
		return errors.New(message)

	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		nexthopConfig := &models.NexthopConfig{
			NexthopTableId:    uint32(*nextHopTable),
			NexthopId:         nextHopId,
			SetNexthopId:      uint32(*nextHopAction),
			RouterInterfaceId: routerInterfaceId,
			NeighborName:      neighborName,
		}
		updates := p4rtFormatter.NextHopTableInsert(nexthopConfig)
		p4client.writeRequest(updates)
	}
	return nil
}

/*
DelNextHopEntry: creates a nexthop label for neighbor identified by router interface and adjacent ipV6 link local address
*/
func DelNextHopEntry(serverAddressPort *string, nextHopId *string, nextHopTable *uint) error {
	if *serverAddressPort == "" || *nextHopTable == 0 || *nextHopId == "" {
		message := fmt.Sprintf("DelNextHopEntry called with serverAddressPort : %s nextHopId : %s nextHopTable : %d",
			*serverAddressPort, *nextHopId, *nextHopTable)
		return errors.New(message)

	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		nexthopConfig := &models.NexthopConfig{
			NexthopTableId: uint32(*nextHopTable),
			NexthopId:      nextHopId,
		}
		updates := p4rtFormatter.NextHopTableDelete(nexthopConfig)
		p4client.writeRequest(updates)
	}
	return nil
}

/*
AddVRF: add a vrf
*/
func AddVRF(serverAddressPort *string, vrfId *string, vrfTable *uint, addVrf *uint) {
	/*
		if becomeMaster(getP4Client(serverAddressPort)) {
			vrfConfig := &models.VrfConfig{
				VrfTableId: uint32(*vrfTable),
				VrfId:      vrfId,
				AddVrf:     uint32(*addVrf),
			}
			updates := p4rtFormatter.VrfTableInsert(vrfConfig)
			p4client.writeRequest(updates)
		}
	*/

}

/*
AddIpV4TableEntry: creates a route entry for a CIDR towards a previously labeled nexthop
*/
func AddIpV4TableEntry(serverAddressPort *string, vrfId *string, destNetwork *string, nextHopId *string,
	ipv4Table *uint, setNextHopId *uint) error {
	if *serverAddressPort == "" || *destNetwork == "" || *nextHopId == "" || *ipv4Table == 0 || *setNextHopId == 0 {
		message := fmt.Sprintf("AddIpV4TableEntry called with serverAddressPort : %s, vrfId : %s destNetwork : %s "+
			"nextHopId : %s ipv4Table : %d setNextHopId : %d ", *serverAddressPort, *vrfId, *destNetwork, *nextHopId,
			*ipv4Table, *setNextHopId)
		return errors.New(message)
	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		ipv4Config := &models.IPv4Config{
			VrfId:           vrfId,
			DestinationCIDR: destNetwork,
			NexthopId:       nextHopId,
			IPv4TableId:     uint32(*ipv4Table),
			SetNexthopId:    uint32(*setNextHopId),
		}
		updates := p4rtFormatter.Ipv4TableInsert(ipv4Config)
		p4client.writeRequest(updates)
	}
	return nil
}

/*
DelIpV4TableEntry: creates a route entry for a CIDR towards a previously labeled nexthop
*/
func DelIpV4TableEntry(serverAddressPort *string, vrfId *string, destNetwork *string, ipv4Table *uint) error {
	if *serverAddressPort == "" || *destNetwork == "" || *ipv4Table == 0 {
		message := fmt.Sprintf("DelIpV4TableEntry called with serverAddressPort : %s, vrfId : %s destNetwork : %s ipv4Table : %d",
			*serverAddressPort, *vrfId, *destNetwork, *ipv4Table)
		return errors.New(message)
	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		ipv4Config := &models.IPv4Config{
			VrfId:           vrfId,
			DestinationCIDR: destNetwork,
			IPv4TableId:     uint32(*ipv4Table),
		}
		updates := p4rtFormatter.Ipv4TableDelete(ipv4Config)
		p4client.writeRequest(updates)
	}
	return nil
}

/*
CreateActionProfileEntry: creates a weighted group of nexthops in support of multipath routing
*/
func CreateActionProfileEntry(serverAddressPort *string, groupId *string, nexthops *string, actionProfileAction *uint, nexthopAction *uint) error {
	if becomeMaster(getP4Client(serverAddressPort)) {
		if *serverAddressPort == "" || *groupId == "" || *nexthops == "" || *actionProfileAction == 0 || *nexthopAction == 0 {
			message := fmt.Sprintf("CreateActionProfileEntry called with serverAddressPort : %s groupId : %s "+
				"nexthops : %s actionProfileAction : %d nexthopAction : %d", *serverAddressPort, *groupId, *nexthops,
				*actionProfileAction, *nexthopAction)
			return errors.New(message)
		}
		profileMembers := []*models.ActionProfileGroupMember{}
		entries := strings.Split(*nexthops, ",")
		if becomeMaster(getP4Client(serverAddressPort)) {
			for i := 0; i < len(entries); i++ {
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
			updates := p4rtFormatter.ActionProfileGroupInsert(actionProfileGroup)
			p4client.writeRequest(updates)
		}
	}
	return nil
}

/*
DeleteActionProfileEntry: creates a weighted group of nexthops in support of multipath routing
*/
func DeleteActionProfileEntry(serverAddressPort *string, groupId *string, actionProfileAction *uint) error {
	if *serverAddressPort == "" || *groupId == "" || *actionProfileAction == 0 {
		message := fmt.Sprintf("DeleteProfileEntry called with serverAddressPort : %s groupId : %s  actionProfileAction : %d",
			*serverAddressPort, *groupId, *actionProfileAction)
		return errors.New(message)
	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		actionProfileGroup := &models.ActionProfileGroup{
			ActionProfileId: uint32(*actionProfileAction),
			GroupId:         groupId,
		}
		updates := p4rtFormatter.ActionProfileGroupDelete(actionProfileGroup)
		p4client.writeRequest(updates)
	}
	return nil
}

/*
AddIpV4WcmpEntry: creates a route entry for a CIDR towards a multipath group (ActionProfile)
*/
func AddIpV4WcmpEntry(serverAddressPort *string, vrf *string, dstNetwork *string, wcmpGroupId *string,
	ipv4TableId *uint, wcmpActionId *uint) error {
	if *serverAddressPort == "" || *dstNetwork == "" || *ipv4TableId == 0 || *wcmpGroupId == "" || *wcmpActionId == 0 {
		message := fmt.Sprintf("AddIpV4WcmpEntry called with "+
			"serverAddressPort : %s,dstNetwork : %s  wcmpGroupId : %s  ipv4TableId : %d wcmpActionId : %d ",
			*serverAddressPort, *dstNetwork, wcmpGroupId, ipv4TableId, wcmpActionId)
		return errors.New(message)
	}
	if becomeMaster(getP4Client(serverAddressPort)) {
		wcmpEntry := &models.IPv4Config{
			VrfId:           vrf,
			DestinationCIDR: dstNetwork,
			IPv4TableId:     uint32(*ipv4TableId),
			WcmpGroupId:     wcmpGroupId,
			SetWcmpGroupId:  uint32(*wcmpActionId),
		}
		updates := p4rtFormatter.Ipv4TableInsertWcmp(wcmpEntry)
		p4client.writeRequest(updates)
	}
	return nil
}
func AddProfileMember(serverAddressPort *string, memberId *uint,
	nexthopId *string, profileId *uint, setNexthopId *uint) error {
	if becomeMaster(getP4Client(serverAddressPort)) {
		profileMember := &models.ActionProfileGroupMember{
			MemberId:     uint32(*memberId),
			NexthopId:    nexthopId,
			ProfileId:    uint32(*profileId),
			SetNexthopId: uint32(*setNexthopId),
		}
		updates := p4rtFormatter.ActionGroupMemberCreate(profileMember)
		p4client.writeRequest(updates)
	}
	return nil
}
