/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package tables

import (
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/v2/internal/models"
)
/*
NextHopTableInsert: used to generate the p4.Update struct needed to create a nexthop entry
 */
func NextHopTableInsert(config *models.NexthopConfig) []*p4.Update {
	updates := []*p4.Update{
		&p4.Update{
			Type: p4.Update_INSERT,
			Entity: &p4.Entity{
				Entity: &p4.Entity_TableEntry{
					TableEntry: &p4.TableEntry{
						TableId: config.NexthopTableId,
						Match: []*p4.FieldMatch{
							{
								FieldId: 1,
								FieldMatchType: &p4.FieldMatch_Exact_{
									Exact: &p4.FieldMatch_Exact{
										Value: []byte(*config.NexthopId),
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
											Value:   []byte(*config.RouterInterfaceId),
										},
										{
											ParamId: 2,
											//Value: destinationIp,
											Value: []byte(*config.NeighborName),
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
NextHopTableDelete: used to generate the p4.Update struct needed to create a nexthop entry
*/
func NextHopTableDelete(config *models.NexthopConfig) []*p4.Update {
	updates := []*p4.Update{
		&p4.Update{
			Type: p4.Update_DELETE,
			Entity: &p4.Entity{
				Entity: &p4.Entity_TableEntry{
					TableEntry: &p4.TableEntry{
						TableId: config.NexthopTableId,
						Match: []*p4.FieldMatch{
							{
								FieldId: 1,
								FieldMatchType: &p4.FieldMatch_Exact_{
									Exact: &p4.FieldMatch_Exact{
										Value: []byte(*config.NexthopId),
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
											Value:   []byte(*config.RouterInterfaceId),
										},
										{
											ParamId: 2,
											//Value: destinationIp,
											Value: []byte(*config.NeighborName),
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
