#
# Copyright 2020-present Open Networking Foundation
#
# SPDX-License-Identifier: Apache-2.0
#
#
P4RUNTIME_ENDPOINT="10.70.2.2:9559"
P4INFO_FILE="sai-p4info.txt"

FIRST_INTERFACE=0
SKIP_INTERFACE=4
LAST_INTERFACE=124

FIRST_MESH_INTERFACE=0
LAST_MESH_INTERFACE=120
INTERFACE_BASE_NAME="Ethernet"
UPLINK_INTERFACE_ID=$LAST_INTERFACE
UPLINK_NEIGHBOR_NAME="SERVER_UPLINK"
BASE_NEIGHBOR_NAME="MESH_NEIGHBOR_"

ROUTER_INTF_MAC="8c:ea:1b:17:64:0c"
MESH_DEST_MAC="00:90:fb:64:cc:9e"
UPLINK_DEST_MAC="00:07:43:4b:7f:50"
MESH_NEXT_HOP_NAME="WEDGE_"
UPLINK_NEXT_HOP_NAME="BROADCOM_UPLINK_NEXT_HOP"
ECMP_WEIGHT=100
WCMP_GROUP_NAME="wedge-mesh"
UPLINK_ROUTE_CIDR="172.16.2.0/24"
MESH_ROUTE_CIDR="172.16.1.0/24"
VRF_ID="vrf-0"

