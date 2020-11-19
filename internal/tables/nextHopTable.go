package tables

import (
	"context"
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/internal/lib"
	"github.com/pins/p4rt-client/internal/models"
	"log"
	//"net"
)

func NextHopTableInsert(client *lib.P4rtClient, config *models.NexthopConfig) (bool, error) {
	success := false
	updates := []*p4.Update{}
	log.Println(*config.RouterInterfaceId)
	log.Println(*config.NeighborIp)
	//destinationIp := net.ParseIP(*config.NeighborIp)
	nextHopTableUpdate := &p4.Update{
		Type: p4.Update_INSERT,
		Entity: &p4.Entity{
			Entity: &p4.Entity_TableEntry{
				TableEntry: &p4.TableEntry{
					TableId: config.NexthopTableId,
					Match: []*p4.FieldMatch{
						&p4.FieldMatch{
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
									&p4.Action_Param{
										ParamId: 1,
										Value:   []byte(*config.RouterInterfaceId),
									},
									&p4.Action_Param{
										ParamId: 2,
										//Value: destinationIp,
										Value: []byte(*config.NeighborIp),
									},
								},
							},
						},
					},
				},
			},
		},
	}
	updates = append(updates, nextHopTableUpdate)
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
		return success, err
	}
	log.Println(resp)
	return success, nil
}
