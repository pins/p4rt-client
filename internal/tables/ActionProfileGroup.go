/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package tables

import (
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/internal/models"
)

/*
ActionProfileGroupInsert: used to generate the p4.Update struct to create a group of nexthops
 */
func ActionProfileGroupInsert( group *models.ActionProfileGroup) []*p4.Update {
	groupMembers := []*p4.ActionProfileAction{}
	for i := 0; i < len(group.Members); i++ {
		groupMember := &p4.ActionProfileAction{
			Action: &p4.Action{
				ActionId: group.SetNexthopId,
				Params: []*p4.Action_Param{
					{
						ParamId: 1,
						Value:   []byte(*group.Members[i].NexthopId),
					},
				},
			},
			Weight: group.Members[i].Weight,
		}
		groupMembers = append(groupMembers, groupMember)
	}
	updates := []*p4.Update{
		&p4.Update{
			Type: p4.Update_INSERT,
			Entity: &p4.Entity{
				Entity: &p4.Entity_TableEntry{
					TableEntry: &p4.TableEntry{
						TableId: group.ActionProfileId,
						Match: []*p4.FieldMatch{
							{
								FieldId: 1,
								FieldMatchType: &p4.FieldMatch_Exact_{
									Exact: &p4.FieldMatch_Exact{
										Value: []byte(*group.GroupId),
									},
								},
							},
						},
						Action: &p4.TableAction{
							Type: &p4.TableAction_ActionProfileActionSet{
								ActionProfileActionSet: &p4.ActionProfileActionSet{
									ActionProfileActions: groupMembers,
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
func ActionGroupMemberCreate(member *models.ActionProfileGroupMember)[]*p4.Update{
	updates := []*p4.Update{
		&p4.Update{
			Entity: &p4.Entity{
				Entity:&p4.Entity_ActionProfileMember{
					ActionProfileMember: &p4.ActionProfileMember{
						ActionProfileId: member.ProfileId,
						MemberId: member.MemberId,
						Action: &p4.Action{
							ActionId: member.SetNexthopId,
							Params: []*p4.Action_Param{
								{
									ParamId: 1,
								    Value:   []byte(*member.NexthopId),
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
