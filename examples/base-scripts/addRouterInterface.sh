#!/bin/bash
set -x 
source $1
source ../vals-p4.sh

echo p4rt-client  -addRouterInt \
	-server=$P4RUNTIME_ENDPOINT \
	-routerTable=$ROUTER_TABLE \
	-routerInterface=$INTERFACE_NAME \
	-routerPortId=$ROUTER_PORT_ID \
	-routerIntMAC=$ROUTER_INTF_MAC  \
	-egressPort=$EGRESS_PORT \
	-setPortMac=$SET_PORT_MAC
