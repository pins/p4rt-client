package tables

import (
	"context"
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/internal/lib"
	"github.com/pins/p4rt-client/internal/models"
	"log"
)

func ActionProfileGroupInsert(client *lib.P4rtClient, group *models.ActionProfileGroup) error {

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
	updates := []*p4.Update{}
	actionProfileGroupUpdate := &p4.Update{
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
	}
	updates = append(updates, actionProfileGroupUpdate)
	writeRequest := p4.WriteRequest{
		DeviceId:   client.DeviceId,
		RoleId:     0,
		ElectionId: client.ElectionId,
		Updates:    updates,
		Atomicity:  p4.WriteRequest_CONTINUE_ON_ERROR,
	}
	resp, err := client.Client.Write(context.Background(), &writeRequest)
	if err != nil {
		log.Fatalf("Failed calling Write %v \n", err)
		return err
	}
	log.Println(resp)
	return nil
}
