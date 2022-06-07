/*
 * Copyright (c) 2022-present Intel Corporation All Rights Reserved
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package models

//func RounterTableInsert(outerInterfaceTableId uint32,routerIntId string,routerIntPortId uint32,routerInterMAC string,setMacAndPort uint32){

type L3Config struct {
	RouterInterfaceTableId uint32
	RouterInterfaceId      *string
	EgressPort             *string
	RouterInterfacePortId  uint32
	RouterInterfaceMAC     *string
	SetMacAndPortId        uint32
}

func (l3config *L3Config) String() string {
	return ""
}
