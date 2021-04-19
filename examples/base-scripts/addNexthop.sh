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

p4rt-client -addNextHop \
      -server=$P4RUNTIME_ENDPOINT  \
      -routerInterface=$INTERFACE_NAME \
      -neighborIp=`macToIpV6 $DEST_MAC` \
      -nextHopId=$NEXTHOP_NAME \
      -nextHopTable=$NEXTHOP_TABLE  \
      -setNextHopAction=$SET_NEXTHOP_ID

