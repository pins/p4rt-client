package lib

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	p4config "github.com/p4lang/p4runtime/go/p4/config/v1"
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"io/ioutil"
	"log"
)

func PushP4Info(client *P4rtClient, p4infoFilename *string) error {
	p4info, err := loadP4Info(p4infoFilename)
	if err != nil {
		log.Printf("loadP4Info failed %v\n", err)
		return err
	}
	deviceConfig := []byte{}
	hash := md5.Sum(deviceConfig)
	cookie := binary.LittleEndian.Uint64(hash[:])

	config := &p4.ForwardingPipelineConfig{
		P4Info:         &p4info,
		P4DeviceConfig: deviceConfig,
		Cookie:         &p4.ForwardingPipelineConfig_Cookie{Cookie: cookie},
	}

	req := &p4.SetForwardingPipelineConfigRequest{
		DeviceId:   client.DeviceId,
		RoleId:     0, // not used
		ElectionId: client.ElectionId,
		Action:     p4.SetForwardingPipelineConfigRequest_VERIFY_AND_COMMIT,
		Config:     config,
	}
	_, err = client.Client.SetForwardingPipelineConfig(context.Background(), req)
	if err != nil {
		log.Printf("SetForwardPipelineConfig failed with %v\n", err)
	}
	return err
}
func loadP4Info(p4infoPath *string) (p4info p4config.P4Info, err error) {
	log.Printf("P4 Info: %s\n", *p4infoPath)
	p4infoBytes, err := ioutil.ReadFile(*p4infoPath)
	if err != nil {
		return
	}
	err = proto.UnmarshalText(string(p4infoBytes), &p4info)
	return
}
