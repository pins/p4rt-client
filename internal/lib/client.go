package lib

import (
	"fmt"
	"log"
	"golang.org/x/net/context"
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"google.golang.org/genproto/googleapis/rpc/code"
)
type P4rtClient struct {
	Client         p4.P4RuntimeClient
	Stream         p4.P4Runtime_StreamChannelClient
	DeviceId       uint64
	ElectionId     *p4.Uint128
}


func  Init(c *P4rtClient) (err error) {
	// Initialize stream for mastership and packet I/O
	log.Println("INIT")
	c.Stream, err = c.Client.StreamChannel(context.Background())
	if err != nil {
		log.Fatalf("Error creating StreamChannel %v\n",err)
	}
	log.Printf("In INIT deviceId is %d\n",c.DeviceId)
	go func() {
		for {
			res, err := c.Stream.Recv()
			if err != nil {
				log.Fatalf("stream recv error: %v\n", err)
			} else if arb := res.GetArbitration(); arb != nil {
				if code.Code(arb.Status.Code) == code.Code_OK {
					log.Println("client is master")
				} else {
					log.Println("client is not master")
				}
			} else {
				fmt.Printf("stream recv: %v\n", res)
			}
		}
	}()
	return
}
func (c *P4rtClient) SetMastership(electionId *p4.Uint128) (err error) {
	c.ElectionId = electionId
	mastershipReq := &p4.StreamMessageRequest{
		Update: &p4.StreamMessageRequest_Arbitration{
			Arbitration: &p4.MasterArbitrationUpdate{
				DeviceId: c.DeviceId,
				ElectionId: electionId,
			},
		},
	}
	err = c.Stream.Send(mastershipReq)
	return
}

