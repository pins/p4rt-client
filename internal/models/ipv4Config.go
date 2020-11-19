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
