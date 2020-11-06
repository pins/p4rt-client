package tables

import (
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/internal/lib"
	"github.com/pins/p4rt-client/internal/models"
	"log"
	"net"
	"context"
)

func NeighborTableInsert(client *lib.P4rtClient,config *models.NeighborConfig)(bool,error){

	destinationMac,err := net.ParseMAC(*config.DestinationMac)
	destinationIp := net.ParseIP(*config.NeighborIP)
	updates := []*p4.Update{}

		neighborTableUpdate := p4.Update{
			Type: p4.Update_INSERT,
			Entity: &p4.Entity{
				Entity: &p4.Entity_TableEntry{
					TableEntry: &p4.TableEntry{
						TableId: config.NeighborTableId,
						Match: []*p4.FieldMatch{
							&p4.FieldMatch{
								FieldId: 1,
								FieldMatchType: &p4.FieldMatch_Exact_{
									Exact: &p4.FieldMatch_Exact{
										Value: []byte(*config.RouterInterfaceId),
									},
								},
							}, &p4.FieldMatch{
								FieldId:              2,
								FieldMatchType:&p4.FieldMatch_Exact_{
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
										&p4.Action_Param{
											ParamId: 1,
											Value: destinationMac,
										},
									},

								},
							},
						},
					},
				},
			},
		}

	success := false
	updates=append(updates,&neighborTableUpdate)
	writeRequest := p4.WriteRequest{
		DeviceId: client.DeviceId,
		RoleId: 0,
		ElectionId: client.ElectionId,
		Updates: updates,
		Atomicity: p4.WriteRequest_CONTINUE_ON_ERROR,
	}
	resp,err :=client.Client.Write(context.Background(),&writeRequest)
	if err!=nil{
		log.Fatalf("Failed calling Write %v \n",err)
		return success,err
	}
	log.Println(resp)
	return success,nil
}