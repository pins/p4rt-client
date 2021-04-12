#!/bin/bash
set -x
source $1
source ../vals-p4.sh

echo p4rt-client  -addIpV4 \
      -server=$P4RUNTIME_ENDPOINT \
      -vrf=$VRF_ID \
      -routedNetwork=$ROUTED_NETWORK \
      -nextHopId=$NEXTHOP_NAME \
      -ipv4table=$IPV4_ROUTE_TABLE \
      -setNextHop=$SET_NEXTHOP
