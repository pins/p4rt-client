package models

type ActionProfileGroupMember struct {
	MemberId  uint32
	Weight    int32
	NexthopId *string
	ProfileId uint32
	SetNexthopId    uint32
}
type ActionProfileGroup struct {
	ActionProfileId uint32
	GroupId         *string
	SetNexthopId    uint32
	DestinationCIDR *string
	Members         []*ActionProfileGroupMember
}
