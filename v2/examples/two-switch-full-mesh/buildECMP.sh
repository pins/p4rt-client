#!/bin/bash
#
# Copyright 2020-present Open Networking Foundation
#
# SPDX-License-Identifier: Apache-2.0
#
#
set -x
VALS_FILE=`pwd`/$1
SCRIPT_DIR="../base-scripts"

source $VALS_FILE 
$SCRIPT_DIR/pushP4Info.sh $VALS_FILE

ALL_INTERFACES=`seq $FIRST_INTERFACE $SKIP_INTERFACE $LAST_INTERFACE` 
MESH_INTERFACES=`seq $FIRST_MESH_INTERFACE $SKIP_INTERFACE $LAST_MESH_INTERFACE`



export INTERFACE_ID=1
for i in $ALL_INTERFACES
do
  export INTERFACE_NAME=$INTERFACE_BASE_NAME$i
  export ROUTER_PORT_ID=$INTERFACE_ID
  export EGRESS_PORT=$i
  $SCRIPT_DIR/addRouter.sh $VALS_FILE
  export INTERFACE_ID=$(($INTERFACE_ID+1))
done

for i in $MESH_INTERFACES
do
   export NEIGHBOR_NAME=$BASE_NEIGHBOR_NAME$i
   export DEST_MAC=$MESH_DEST_MAC
   export INTERFACE_NAME=$INTERFACE_BASE_NAME$i
   export NEXTHOP_NAME=$MESH_NEXT_HOP_NAME$i
   $SCRIPT_DIR/addNeighbor.sh $VALS_FILE
   $SCRIPT_DIR/addNexthop.sh $VALS_FILE
done

#uplink
echo $UPLINK_DEST_MAC
echo $MESH_DEST_MAC

export NEIGHBOR_NAME=$UPLINK_NEIGHBOR_NAME
export DEST_MAC=$UPLINK_DEST_MAC 
export INTERFACE_NAME=$INTERFACE_BASE_NAME$UPLINK_INTERFACE_ID 
export NEXTHOP_NAME=$UPLINK_NEXT_HOP_NAME

$SCRIPT_DIR/addNeighbor.sh $VALS_FILE 
$SCRIPT_DIR/addNexthop.sh $VALS_FILE

#end uplink
NEXTHOP_LIST=

for i in $MESH_INTERFACES
do
   NEXTHOP_LIST="$NEXTHOP_LIST$MESH_NEXT_HOP_NAME$i:$ECMP_WEIGHT,"
done
#trim last comma
NEXTHOP_LIST=${NEXTHOP_LIST::-1}
export NEXTHOP_LIST
$SCRIPT_DIR/addActionProfile.sh $VALS_FILE

export ROUTED_NETWORK=$UPLINK_ROUTE_CIDR
export NEXTHOP_NAME=$UPLINK_NEXT_HOP_NAME
$SCRIPT_DIR/addIpV4Route.sh $VALS_FILE
export ROUTED_NETWORK=$MESH_ROUTE_CIDR
$SCRIPT_DIR/addIpV4RouteWCMP.sh $VALS_FILE



