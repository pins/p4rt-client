package models

//func RounterTableInsert(outerInterfaceTableId uint32,routerIntId string,routerIntPortId uint32,routerInterMAC string,setMacAndPort uint32){


type L3Config struct{
	RouterInterfaceTableId uint32
	RouterInterfaceId *string
    EgressPort uint32
	RouterInterfacePortId uint32
	RouterInterfaceMAC *string
	SetMacAndPortId uint32
}
func (l3config *L3Config)String()string{
	return ""
}