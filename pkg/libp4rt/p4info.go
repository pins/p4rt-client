/*
 * Copyright (c) 2022-present Intel Corporation All Rights Reserved
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package libp4rt

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	p4config "github.com/p4lang/p4runtime/go/p4/config/v1"
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
)

func pushP4Info(client *P4rtClient, p4infoFilename *string) error {
	p4info, err := loadP4Info(p4infoFilename)
	if err != nil {
		message := fmt.Sprintf("loadP4Info failed %v\n", err)
		log.Error(&message)
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
	//js,_ := json.Marshal(req)
	message := fmt.Sprintf("Calling SetForwardingPipelineConfig with : %s", req.String())
	log.Debug(&message)

	_, err = client.Client.SetForwardingPipelineConfig(context.Background(), req)
	if err != nil {
		message := fmt.Sprintf("SetForwardPipelineConfig failed with %v\n", err)
		log.Error(&message)
	}
	return err
}
func loadP4Info(p4infoPath *string) (p4info p4config.P4Info, err error) {
	message := fmt.Sprintf("P4 Info: %s\n", *p4infoPath)
	log.Info(&message)
	p4infoBytes, err := ioutil.ReadFile(*p4infoPath)
	if err != nil {
		return
	}
	p4content := string(p4infoBytes)
	message = fmt.Sprintf("P4Info Contents %s", p4content)
	log.Debug(&message)
	err = proto.UnmarshalText(p4content, &p4info)
	return
}
