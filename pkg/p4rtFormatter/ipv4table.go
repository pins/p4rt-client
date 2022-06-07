/*
 * Copyright (c) 2022-present Intel Corporation All Rights Reserved
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package p4rtFormatter

import (
	"log"
	"net"
	"strconv"
	"strings"

	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/pkg/models"
)

//TODO test this with IPv6 routes and if works as is rename the functions and go file appropriately
/*
Ipv4TableInsert: used to generate the p4.Update struct needed to create a route entry to a specific nexthop
*/
func Ipv4TableInsert(config *models.IPv4Config) []*p4.Update {
	substrs := strings.Split(*config.DestinationCIDR, "/")

	log.Println(*config.DestinationCIDR)
	if len(substrs) != 2 {
		log.Printf("ipv4TableInsert invalid CIDR %s\n", config.DestinationCIDR)
	}
	destNetwork := (net.ParseIP(substrs[0]).To4())
	mask, err := strconv.ParseInt(substrs[1], 10, 32)
	if err != nil {
		//message := fmt.Sprintf("Failure to parse CIDR %v\n", err)
		//logging.Error(&message)
	}

	updates := []*p4.Update{
		&p4.Update{
			Type: p4.Update_INSERT,
			Entity: &p4.Entity{
				Entity: &p4.Entity_TableEntry{
					TableEntry: &p4.TableEntry{
						TableId: config.IPv4TableId,
						Match: []*p4.FieldMatch{
							{
								FieldId: 1,
								FieldMatchType: &p4.FieldMatch_Exact_{
									Exact: &p4.FieldMatch_Exact{
										Value: []byte(*config.VrfId),
									},
								},
							}, {
								FieldId: 2,
								FieldMatchType: &p4.FieldMatch_Lpm{
									Lpm: &p4.FieldMatch_LPM{
										Value:     destNetwork,
										PrefixLen: int32(mask),
									},
								},
							},
						},
						Action: &p4.TableAction{
							Type: &p4.TableAction_Action{
								Action: &p4.Action{
									ActionId: config.SetNexthopId,
									Params: []*p4.Action_Param{
										{
											ParamId: 1,
											Value:   []byte(*config.NexthopId),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return updates
}

/*
Ipv4TableInsertWcmp: used to generate the p4.Update struct needed to create a route entry to a multipath group
*/
func Ipv4TableInsertWcmp(config *models.IPv4Config) []*p4.Update {
	substrs := strings.Split(*config.DestinationCIDR, "/")

	log.Println(*config.DestinationCIDR)
	if len(substrs) != 2 {
		log.Printf("ipv4TableInsert invalid CIDR %s\n", config.DestinationCIDR)
	}
	destNetwork := (net.ParseIP(substrs[0]).To4())
	mask, err := strconv.ParseInt(substrs[1], 10, 32)
	if err != nil {
		//message := fmt.Sprintf("Failure to parse CIDR %v\n", err)
		//logging.Error(&message)
	}

	updates := []*p4.Update{
		&p4.Update{
			Type: p4.Update_INSERT,
			Entity: &p4.Entity{
				Entity: &p4.Entity_TableEntry{
					TableEntry: &p4.TableEntry{
						TableId: config.IPv4TableId,
						Match: []*p4.FieldMatch{
							{
								FieldId: 1,
								FieldMatchType: &p4.FieldMatch_Exact_{
									Exact: &p4.FieldMatch_Exact{
										Value: []byte(*config.VrfId),
									},
								},
							}, {
								FieldId: 2,
								FieldMatchType: &p4.FieldMatch_Lpm{
									Lpm: &p4.FieldMatch_LPM{
										Value:     destNetwork,
										PrefixLen: int32(mask),
									},
								},
							},
						},
						Action: &p4.TableAction{
							Type: &p4.TableAction_Action{
								Action: &p4.Action{
									ActionId: config.SetWcmpGroupId,
									Params: []*p4.Action_Param{
										{
											ParamId: 1,
											Value:   []byte(*config.WcmpGroupId),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return updates
}

/*
Ipv4TableDelete: used to generate the p4.Update struct needed to create a route entry to a specific nexthop
*/
func Ipv4TableDelete(config *models.IPv4Config) []*p4.Update {
	substrs := strings.Split(*config.DestinationCIDR, "/")

	log.Println(*config.DestinationCIDR)
	if len(substrs) != 2 {
		log.Printf("ipv4TableInsert invalid CIDR %s\n", config.DestinationCIDR)
	}
	destNetwork := (net.ParseIP(substrs[0]).To4())
	mask, err := strconv.ParseInt(substrs[1], 10, 32)
	if err != nil {
		//message := fmt.Sprintf("Failure to parse CIDR %v\n", err)
		//logging.Error(&message)
	}

	updates := []*p4.Update{
		&p4.Update{
			Type: p4.Update_DELETE,
			Entity: &p4.Entity{
				Entity: &p4.Entity_TableEntry{
					TableEntry: &p4.TableEntry{
						TableId: config.IPv4TableId,
						Match: []*p4.FieldMatch{
							{
								FieldId: 1,
								FieldMatchType: &p4.FieldMatch_Exact_{
									Exact: &p4.FieldMatch_Exact{
										Value: []byte(*config.VrfId),
									},
								},
							}, {
								FieldId: 2,
								FieldMatchType: &p4.FieldMatch_Lpm{
									Lpm: &p4.FieldMatch_LPM{
										Value:     destNetwork,
										PrefixLen: int32(mask),
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return updates
}
