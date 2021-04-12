package models

type NeighborConfig struct {
	RouterInterfaceId     *string
	NeighborName            *string
	DestinationMac        *string
	NeighborTableId       uint32
	NeighborTableActionId uint32
}
