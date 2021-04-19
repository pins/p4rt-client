#!/bin/bash
#
# Copyright 2020-present Open Networking Foundation
#
# SPDX-License-Identifier: Apache-2.0
#
#
set -x
source $1
source ../vals-p4.sh

p4rt-client  -addIpV4 \
      -server=$P4RUNTIME_ENDPOINT \
      -vrf=$VRF_ID \
      -routedNetwork=$ROUTED_NETWORK \
      -nextHopId=$NEXTHOP_NAME \
      -ipv4table=$IPV4_ROUTE_TABLE \
      -setNextHop=$SET_NEXTHOP
