package tables

import (
	"context"
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/internal/lib"
	"github.com/pins/p4rt-client/internal/models"
	"log"
	"net"
	"strconv"
	"strings"
)

func Ipv4TableInsert(client *lib.P4rtClient, config *models.IPv4Config) (bool, error) {
	substrs := strings.Split(*config.DestinationCIDR, "/")

	log.Println(*config.DestinationCIDR)
	if len(substrs) != 2 {
		log.Printf("ipv4TableInsert invalid CIDR %s\n", config.DestinationCIDR)
	}
	destNetwork := (net.ParseIP(substrs[0]).To4())
	mask, err := strconv.ParseInt(substrs[1], 10, 32)

	updates := []*p4.Update{}

	destinationTableUpdate := p4.Update{
		Type: p4.Update_INSERT,
		Entity: &p4.Entity{
			Entity: &p4.Entity_TableEntry{
				TableEntry: &p4.TableEntry{
					TableId: config.IPv4TableId,
					Match: []*p4.FieldMatch{
						&p4.FieldMatch{
							FieldId: 1,
							FieldMatchType: &p4.FieldMatch_Exact_{
								Exact: &p4.FieldMatch_Exact{
									Value: []byte(*config.VrfId),
								},
							},
						}, &p4.FieldMatch{
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
									&p4.Action_Param{
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
	}
	success := false
	updates = append(updates, &destinationTableUpdate)
	b, err2 := writeUpdates(client, updates, err, success)
	if err2 != nil {
		return b, err2
	}
	return success, nil
}
func Ipv4TableInsertWcmp(client *lib.P4rtClient, config *models.IPv4Config) (bool, error) {
	substrs := strings.Split(*config.DestinationCIDR, "/")

	log.Println(*config.DestinationCIDR)
	if len(substrs) != 2 {
		log.Printf("ipv4TableInsert invalid CIDR %s\n", config.DestinationCIDR)
	}
	destNetwork := (net.ParseIP(substrs[0]).To4())
	mask, err := strconv.ParseInt(substrs[1], 10, 32)

	updates := []*p4.Update{}

	destinationTableUpdate := p4.Update{
		Type: p4.Update_INSERT,
		Entity: &p4.Entity{
			Entity: &p4.Entity_TableEntry{
				TableEntry: &p4.TableEntry{
					TableId: config.IPv4TableId,
					Match: []*p4.FieldMatch{
						&p4.FieldMatch{
							FieldId: 1,
							FieldMatchType: &p4.FieldMatch_Exact_{
								Exact: &p4.FieldMatch_Exact{
									Value: []byte(*config.VrfId),
								},
							},
						}, &p4.FieldMatch{
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
									&p4.Action_Param{
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
	}
	success := false
	updates = append(updates, &destinationTableUpdate)
	b, err2 := writeUpdates(client, updates, err, success)
	if err2 != nil {
		return b, err2
	}
	return success, nil
}

func writeUpdates(client *lib.P4rtClient, updates []*p4.Update, err error, success bool) (bool, error) {
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
	return false, nil
}
