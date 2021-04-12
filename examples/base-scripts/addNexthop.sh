#!/bin/bash
set -x
source $1
source ../vals-p4.sh

echo p4rt-client -addNextHop \
      -server=$P4RUNTIME_ENDPOINT  \
      -routerInterface=$INTERFACE_NAME \
      -neighborIp=`macToIpV6 $DEST_MAC` \
      -nextHopId=$NEXTHOP_NAME \
      -nextHopTable=$NEXTHOP_TABLE  \
      -setNextHopAction=$SET_NEXTHOP_ID

