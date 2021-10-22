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
      -neighborName=$NEIGHBOR_NAME \
      -nextHopId=$NEXTHOP_NAME 
