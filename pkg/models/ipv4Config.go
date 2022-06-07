/*
 * Copyright (c) 2022-present Intel Corporation All Rights Reserved
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package models

type IPv4Config struct {
	VrfId           *string
	DestinationCIDR *string
	NexthopId       *string
	WcmpGroupId     *string
	IPv4TableId     uint32
	SetNexthopId    uint32
	SetWcmpGroupId  uint32
}
