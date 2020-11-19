package tables

import (
	"context"
	"encoding/binary"
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/internal/lib"
	"github.com/pins/p4rt-client/internal/models"
	"log"
	"net"
)

func RounterTableInsert(client *lib.P4rtClient, config *models.L3Config) {
	routerInterfaceMAC, err := net.ParseMAC(*config.RouterInterfaceMAC)
	updates := []*p4.Update{}
	portId := make([]byte, 4)
	binary.BigEndian.PutUint32(portId, config.EgressPort)
	routeInterfaceUpdate := p4.Update{
		Type: p4.Update_INSERT,
		Entity: &p4.Entity{
			Entity: &p4.Entity_TableEntry{
				TableEntry: &p4.TableEntry{
					TableId: config.RouterInterfaceTableId,
					Match: []*p4.FieldMatch{
						&p4.FieldMatch{
							FieldId: 1,
							FieldMatchType: &p4.FieldMatch_Exact_{
								Exact: &p4.FieldMatch_Exact{
									Value: []byte(*config.RouterInterfaceId),
								},
							},
						},
					},
					Action: &p4.TableAction{
						Type: &p4.TableAction_Action{
							Action: &p4.Action{
								ActionId: config.SetMacAndPortId,
								Params: []*p4.Action_Param{
									&p4.Action_Param{
										ParamId: 1,
										Value:   portId,
									},
									&p4.Action_Param{
										ParamId: 2,
										Value:   routerInterfaceMAC,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	updates = append(updates, &routeInterfaceUpdate)
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
	}
	log.Println(resp)
}
