/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package lib

import (
	"encoding/json"
	"fmt"
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/v2/internal/logging"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc"
	"log"
)
var DEFAULT_DEVICE_ID uint64 = 183807201

type P4rtClient struct {
	Client     p4.P4RuntimeClient
	Stream     p4.P4Runtime_StreamChannelClient
	DeviceId   uint64
	ElectionId *p4.Uint128
	MasterChan chan bool
}
func NewP4rtClient(serverAddressPort *string)*P4rtClient{
	conn, err := grpc.Dial(*serverAddressPort, grpc.WithInsecure())
	if err != nil {
		message := fmt.Sprintf("Failed to dial switch: %s", err)
		logging.Error(&message)
	}
	client := p4.NewP4RuntimeClient(conn)
	p4client := &P4rtClient{
		Client:   client,
		DeviceId: DEFAULT_DEVICE_ID,
	}
	Init(p4client)
	if logging.GetDebug(){
		js,_ := json.Marshal(p4client)
		message := fmt.Sprintf("NewP4rtClient returning : %v",js)
		logging.Debug(&message)
	}
	return p4client
}
func Init(client *P4rtClient) (err error) {
	message := fmt.Sprintf("In INIT deviceId is %d\n", client.DeviceId)
	logging.Info(&message)
	client.MasterChan = make(chan bool)
	// Initialize stream for mastership and packet I/O
	client.Stream, err = client.Client.StreamChannel(context.Background())
	if err != nil {
		message := fmt.Sprintf("Error creating StreamChannel %v\n", err)
		logging.Error(&message)
	}
	go func() {
		for {
			res, err := client.Stream.Recv()
			if err != nil {
				message := fmt.Sprintf("stream recv error: %v\n", err)
				logging.Error(&message)
			} else if arb := res.GetArbitration(); arb != nil {
				if code.Code(arb.Status.Code) == code.Code_OK {
					log.Println("client is master")
					client.MasterChan <- true
				} else {
					message := fmt.Sprintf("Returned ElectionId %v\n", arb.ElectionId)
					logging.Info(&message)
					electionId := arb.ElectionId.Low
					newElectionId := &p4.Uint128{
						Low:  electionId + uint64(1),
						High: arb.ElectionId.High,
					}
					go client.SetMastership(newElectionId)
				}
			} else {
				message := "PacketOut message received"
				logging.Info(&message)
			}
		}
	}()
	return
}
func (client *P4rtClient) SetMastership(electionId *p4.Uint128) {
	message := fmt.Sprintf("SetMastership called with %v ElectionId \n", electionId)
	logging.Info(&message)
	client.ElectionId = electionId
	mastershipReq := &p4.StreamMessageRequest{
		Update: &p4.StreamMessageRequest_Arbitration{
			Arbitration: &p4.MasterArbitrationUpdate{
				DeviceId:   client.DeviceId,
				ElectionId: electionId,
			},
		},
	}
	err := client.Stream.Send(mastershipReq)
	if err != nil {
		message := fmt.Sprintf("Set Mastership failed with %v, unable to proceed", err)
		logging.Error(&message)
	}
}

func (client *P4rtClient)writeRequest(updates []*p4.Update)error{
	writeRequest := p4.WriteRequest{
		DeviceId:   client.DeviceId,
		RoleId:     0,
		ElectionId: client.ElectionId,
		Updates:    updates,
		Atomicity:  p4.WriteRequest_CONTINUE_ON_ERROR,
	}
	if logging.GetDebug(){
		js,_ := json.Marshal(writeRequest)
		message := fmt.Sprintf("client.Write being called with \n %v",js)
		//message := fmt.Sprintf("writeRequest called with %s ",writeRequest.String())
		logging.Debug(&message)
	}
	//TODO Currently response from Write call is "" if changes perhaps log
	_, err := client.Client.Write(context.Background(), &writeRequest)
	if err != nil {
		message := fmt.Sprintf("Failed calling Write %v \n", err)
		logging.Error(&message)
	}
	return err

}
