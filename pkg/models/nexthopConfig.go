/*
 * Copyright (c) 2022-present Intel Corporation All Rights Reserved
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package models

type NexthopConfig struct {
	NexthopTableId    uint32
	NexthopId         *string
	SetNexthopId      uint32
	RouterInterfaceId *string
	NeighborName      *string
}
