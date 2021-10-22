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

p4rt-client -addNeighbor \
      -server=$P4RUNTIME_ENDPOINT \
      -neighborName=$NEIGHBOR_NAME \
      -destMAC=$DEST_MAC \
      -routerInterface=$INTERFACE_NAME 
