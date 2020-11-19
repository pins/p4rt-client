package lib

import (
	"fmt"
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"log"
)

type P4rtClient struct {
	Client     p4.P4RuntimeClient
	Stream     p4.P4Runtime_StreamChannelClient
	DeviceId   uint64
	ElectionId *p4.Uint128
	MasterChan chan bool
}

func Init(client *P4rtClient) (err error) {
	client.MasterChan = make(chan bool)
	// Initialize stream for mastership and packet I/O
	log.Println("INIT")
	client.Stream, err = client.Client.StreamChannel(context.Background())
	if err != nil {
		log.Fatalf("Error creating StreamChannel %v\n", err)
	}
	log.Printf("In INIT deviceId is %d\n", client.DeviceId)
	go func() {
		for {
			res, err := client.Stream.Recv()
			if err != nil {
				log.Fatalf("stream recv error: %v\n", err)
			} else if arb := res.GetArbitration(); arb != nil {
				if code.Code(arb.Status.Code) == code.Code_OK {
					log.Println("client is master")
					client.MasterChan <- true
				} else {
					log.Printf("Returned ElectionId %v\n", arb.ElectionId)
					electionId := arb.ElectionId.Low
					newElectionId := &p4.Uint128{
						Low:  electionId + uint64(1),
						High: arb.ElectionId.High,
					}

					go client.SetMastership(newElectionId)
				}
			} else {
				fmt.Printf("stream recv: %s\n", string(res.GetPacket().Payload))
			}
		}
	}()
	return
}
func (client *P4rtClient) SetMastership(electionId *p4.Uint128) {
	log.Printf("SetMastership called with %v ElectionId \n", electionId)
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
		log.Fatalf("Set Mastership failed with %v, unable to proceed", err)
		client.MasterChan <- false
	}
}
