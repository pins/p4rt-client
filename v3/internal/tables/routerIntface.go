/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package tables

import (
   "fmt"
   //sai "github.com/pins/p4rt-client/v2/protos/sai_go"
   p4 "github.com/p4lang/p4runtime/go/p4/v1"
   "github.com/pins/p4rt-client/v3/internal/logging"
   "github.com/pins/p4rt-client/v3/internal/models"
   "net"
)
/*
RouterTableInsert: generate the p4.Update struct needed to create a Router Interface Entry
 */
func RouterTableInsert( config *models.L3Config) []*p4.Update{
   routerInterfaceMAC, err := net.ParseMAC(*config.RouterInterfaceMAC)
   if err != nil{
      message := fmt.Sprintf("Failed to parse %s %v",*config.RouterInterfaceMAC,err)
      logging.Error(&message)
   }

   //portStr := fmt.Sprintf("Ethernet%d",config.EgressPort)
   portStr := fmt.Sprintf("%d",config.EgressPort)
   portId := []byte(portStr)
   //binary.BigEndian.PutUint32(portId, config.EgressPort)
   updates := []*p4.Update{
      &p4.Update{
         Type: p4.Update_INSERT,
         Entity: &p4.Entity{
            Entity: &p4.Entity_TableEntry{
               TableEntry: &p4.TableEntry{
                  TableId: config.RouterInterfaceTableId,
                  Match: []*p4.FieldMatch{
                     {
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
                              {
                                 ParamId: 1,
                                 Value:   portId,
                              },
                              {
                                 ParamId: 2,
                                 Value:   routerInterfaceMAC,
                                 //Value: []byte(*config.RouterInterfaceMAC),
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
RouterTableDelete: generate the p4.Update struct needed to remove a Router Interface Entry
*/
func RouterTableDelete( config *models.L3Config) []*p4.Update{
   routerInterfaceMAC, err := net.ParseMAC(*config.RouterInterfaceMAC)
   if err != nil{
      message := fmt.Sprintf("Failed to parse %s %v",*config.RouterInterfaceMAC,err)
      logging.Error(&message)
   }
#portStr := fmt.Sprintf("%d",config.EgressPort)
   portId := []byte(config.EgressPort)
//   binary.BigEndian.PutUint32(portId, config.EgressPort)
   updates := []*p4.Update{
      &p4.Update{
         Type: p4.Update_DELETE,
         Entity: &p4.Entity{
            Entity: &p4.Entity_TableEntry{
               TableEntry: &p4.TableEntry{
                  TableId: config.RouterInterfaceTableId,
                  Match: []*p4.FieldMatch{
                     {
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
                              {
                                 ParamId: 1,
                                 Value:   portId,
                              },
                              {
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
      },
   }
   return updates
}
