#
# Copyright 2020-present Open Networking Foundation
#
# SPDX-License-Identifier: Apache-2.0
#
#

P4RUNTIME_ENDPOINT=
#Listen address and port for the P4Runtime service running on the target switch

P4INFO_FILE="sai-p4info.txt"
#The file name of the generated p4info file to be uploaded

INTERFACE_NAME=
#Unique arbitrary name given to the logical interface

ROUTER_PORT_ID=
#Unique arbitrary number given to the logical interface

ROUTER_INTF_MAC=
#MAC Address to be used on logical interface - in SONiC world typically same as all other interfaces

EGRESS_PORT=
#Physical port of switch associated with logical port

NEIGHBOR_NAME=
#Friendly name to use for adjacent neighbor

DEST_MAC=
#MAC address of the adjacent neighbor

NEXTHOP_NAME=
#Unique arbitrary name given to Neighbor/Interface combination

ROUTED_NETWORK=
#CIDR of the network to route

VRF_ID=
#Name of VRF to use for routing this referenced network

