/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package tables

import (
	"fmt"
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/internal/logging"
	"github.com/pins/p4rt-client/internal/models"
	"net"
)

/*
NeighborTableInsert: used to generate the p4.Update struct needed to create a neighbor entry
 */
func NeighborTableInsert( config *models.NeighborConfig) []*p4.Update {

	destinationMac, err := net.ParseMAC(*config.DestinationMac)
	if err != nil{
		message := fmt.Sprintf("Failed to parse Mac %s %v",*config.DestinationMac,err)
		logging.Error(&message)
	}
	destinationIp := net.ParseIP(*config.NeighborIP)
	updates := []*p4.Update{
		{
			Type: p4.Update_INSERT,
			Entity: &p4.Entity{
				Entity: &p4.Entity_TableEntry{
					TableEntry: &p4.TableEntry{
						TableId: config.NeighborTableId,
						Match: []*p4.FieldMatch{
							{
								FieldId: 1,
								FieldMatchType: &p4.FieldMatch_Exact_{
									Exact: &p4.FieldMatch_Exact{
										Value: []byte(*config.RouterInterfaceId),
									},
								},
							}, {
								FieldId: 2,
								FieldMatchType: &p4.FieldMatch_Exact_{
									Exact: &p4.FieldMatch_Exact{
										Value: destinationIp,
									},
								},
							},
						},
						Action: &p4.TableAction{
							Type: &p4.TableAction_Action{
								Action: &p4.Action{
									ActionId: config.NeighborTableActionId,
									Params: []*p4.Action_Param{
										{
											ParamId: 1,
											Value:   destinationMac,
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
