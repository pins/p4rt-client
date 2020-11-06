package main

import (
	"flag"
	p4 "github.com/p4lang/p4runtime/go/p4/v1"
	"github.com/pins/p4rt-client/internal/lib"
	"github.com/pins/p4rt-client/internal/models"
	"github.com/pins/p4rt-client/tables"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

// Authentication holds the login/password
type Authentication struct {
	Login    string
	Password string
	// RequireTransportSecurity indicates whether the credentials requires transport security
}

// GetRequestMetadata gets the current request metadata
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"login":    a.Login,
		"password": a.Password,
	}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return true
}
func main() {
	useSsl := flag.Bool("ssl", false, "use ssl")
	useAuth := flag.Bool("auth", false, "use auth")
	serverAddressPort := flag.String("server", "", "address and port of p4rt server on swich" )
	fqdn := flag.String("fqdn", "", "FQDN of the service to match what is in server.crt")
	crtFile := flag.String("cert", "cert/server.crt", "Public cert for server to establish tls session")
	//#########################################
	pushP4Info := flag.Bool("pushP4Info",false,"push P4Info text file to switch")
	p4info     := flag.String("p4info","","p4info text filename which describes p4 application")

	//#########################################
	addL3 := flag.Bool("addL3",false,"add entries needed for layer 3 routing")
	routerIntTableId :=flag.Uint("routerTable",0,"id for ingress.routing.router_interface_table from p4info.txt")
	routerInterfaceId := flag.String("routerInterface","","name to give to router interface")
	routerIntPort  := flag.Uint("routerPortId",1,"port number to assign to router interface port")
	routerIntMAC := flag.String("routerIntMAC","","MAC address to be used for router interface e.g. 00:00:00:11:22:dd")
	egressPort := flag.Uint("egressPort",0,"switch port to egress")
	setMacPort := flag.Uint("setPortMac",0,"action id associated with action ingress.routing.set_port_and_src_mac ")
	nextHopTable := flag.Uint("nextHopTable",0,"table id associated with ingress.routing.nexthop_table from p4info.txt")
	nextHopId    := flag.String("nextHopId","","name to associate with next hop entry")
	nextHopAction := flag.Uint("setNextHopAction",0,"action id associated with ingress.routing.set_nexthop_id from p4info.txt")
	neighborIp    := flag.String("neighborIp","","ip address of next hop neighbor eg. 10.10.10.2")
	destMAC   := flag.String("destMAC","","MAC address for neighbor IP e.g. 11:22:33:44:55:66")
	neighborTable := flag.Uint("neighborTable",0,"id associated with ingress.routing.neighbor_table from p4info.txt")
	setDestMac := flag.Uint("setDestMacAction",0,"id associated with ingress.routing.set_dst_mac from p4info.txt")
	vrfId      := flag.String("vrf","default","name of vrf to use")
	destNetwork := flag.String("routedNetwork","","CIDR of network to route e.g. 1.2.3.4/8")
	ipv4Table   := flag.Uint("ipv4table",0,"id associated with ingress.routing.ipv4_table from p4info.txt")
	setNextHopId := flag.Uint("setNextHop",0,"id associated with action ingress.routing.set_nexthop_id from p4info.txt")
	help := flag.Bool("help",false,"print usage help")

	flag.Parse()


	if *help{
		basicUsage()
		return
	}
	if *pushP4Info&&*addL3{
		basicUsage()
		return
	}


	if *pushP4Info{
		if *p4info==""||*serverAddressPort==""{
			pushP4infoUsage()
			return
		}
	}else if *addL3 {
		//sanitiy check
		if *routerIntTableId == 0 || *routerInterfaceId == "" || *routerIntMAC == "" || *setMacPort == 0 || *nextHopTable == 0 || *nextHopId == "" || *nextHopAction == 0 ||
			*egressPort==0||*neighborIp == "" || *destMAC == "" || *vrfId == "" || *destNetwork == "" || *ipv4Table == 0 || *setNextHopId == 0 || *serverAddressPort=="" {
			addL3Usage()
			return
		}
	}else{
		basicUsage()
		return
	}

	var conn *grpc.ClientConn
	var err error
	auth := Authentication{
		Login:    "admin",
		Password: "admin",
	}
	if *useSsl && *useAuth {

		creds, err := credentials.NewClientTLSFromFile(*crtFile, *fqdn)
		conn, err = grpc.Dial(*serverAddressPort, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
		if err != nil {
			log.Fatalf("could not load tls cert: %s", err)
		}
	} else if *useSsl {
		creds, err := credentials.NewClientTLSFromFile("cert/server.crt", *fqdn)
		conn, err = grpc.Dial(*serverAddressPort, grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Fatalf("could not load tls cert: %s", err)
		}
	} else if *useAuth {
		conn, err = grpc.Dial(*serverAddressPort, grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	} else {

		log.Println(*serverAddressPort)
		conn, err = grpc.Dial(*serverAddressPort, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to dial switch: %s", err)
		}
	}

	defer conn.Close()
	client := p4.NewP4RuntimeClient(conn)
	p4client:=&lib.P4rtClient{
		Client:client,
		DeviceId:uint64(183807201),
	}
	lib.Init(p4client)
	if *addL3{
		p4client.SetMastership(&p4.Uint128{High: 0, Low: 1})
		l3Config := &models.L3Config{
			RouterInterfaceTableId: uint32(*routerIntTableId),
			RouterInterfaceId: routerInterfaceId,
			EgressPort: uint32(*egressPort),
			RouterInterfacePortId: uint32(*routerIntPort),
			RouterInterfaceMAC: routerIntMAC,
			SetMacAndPortId: uint32(*setMacPort),
		}
		tables.RounterTableInsert(p4client,l3Config)
		neighborConfig := &models.NeighborConfig{
			RouterInterfaceId: routerInterfaceId,
			NeighborIP: neighborIp,
			DestinationMac: destMAC,
			NeighborTableId: uint32(*neighborTable),
			NeighborTableActionId: uint32(*setDestMac),
		}
		tables.NeighborTableInsert(p4client,neighborConfig)
		nexthopConfig := &models.NexthopConfig{
			NexthopTableId:uint32(*nextHopTable),
			NexthopId: nextHopId,
			SetNexthopId: uint32(*setNextHopId),
			RouterInterfaceId: routerInterfaceId,
			NeighborIp: neighborIp,
		}
		tables.NextHopTableInsert(p4client,nexthopConfig)
		ipv4Config := &models.IPv4Config{
			VrfId: vrfId,
			DestinationCIDR: destNetwork,
			NexthopId: nextHopId,
			IPv4TableId: uint32(*ipv4Table),
			SetNexthopId: uint32(*setNextHopId),
		}
		tables.Ipv4TableInsert(p4client,ipv4Config)
	}
}
func basicUsage(){
	log.Println("p4rt-client can communication with p4runtime on a sonic based switch")
	log.Println("\n\nTo Set the p4info file:")
	pushP4infoUsage()
	log.Println("\n\nTo Configure Layer 3 routing:")
	addL3Usage()
	log.Println("\n\nTo see more help on arguments:\n./p4rt-client -h")
}
func addL3Usage(){
	usage:= `
Usage:
./p4rt-client -addL3 \
-server=$P4RT_LISTEN_ADDRESS  \
-routerTable=$ROUTER_TABLE_ID \
-routerInterface=$ROUTER_INTERFACE_NAME \
-routerPortId=$ROUTER_PORT_ID \
-routerIntMAC=$ROUTER_INTERFACE_MAC_ADDR \
-setPortMac=$SET_PORT_MAC_ACTION \
-nextHopTable=$NEXT_HOP_TABLE \
-nextHopId=$NEXT_HOP_NAME \
-setNextHopAction=$NEXT_HOP_ACTION_ID \
-neighborIp=$NEIGHBOR_IP \
-destMAC=$DEST_MAC \
-neighborTable=$NEIGHBOR_TABLE_ID \
-setDestMacAction=$SET_DST_MAC_ACTION_ID \
-vrf=$VRF_NAME \
-routedNetwork=$ROUTED_NETWORK_CIDR \
-ipv4table=$IPV_TABLE_ID \
-setNextHop=$SET_NEXT_HOP_ACTION_ID`
	log.Println(usage)
}

func pushP4infoUsage(){
	usage:=`
Usage:
./p4rt-client -pushP4info -p4info=$P4_INFO_FILENAME`
log.Println(usage)
}