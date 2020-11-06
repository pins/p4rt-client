package models

type IPv4Config struct {
	VrfId *string
	DestinationCIDR *string
	NexthopId *string
	IPv4TableId uint32
	SetNexthopId uint32
}