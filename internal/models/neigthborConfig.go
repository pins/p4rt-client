/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package models

type NeighborConfig struct {
	RouterInterfaceId     *string
	NeighborIP            *string
	DestinationMac        *string
	NeighborTableId       uint32
	NeighborTableActionId uint32
}
