package models


type NeighborConfig struct{
	RouterInterfaceId *string
	NeighborIP *string
	DestinationMac *string
	NeighborTableId uint32
	NeighborTableActionId uint32

}