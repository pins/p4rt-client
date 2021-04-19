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
      -neighborIp=$NEIGHBOR_IP \
      -destMAC=$DEST_MAC \
      -routerInterface=$INTERFACE_NAME \
      -neighborTable=$NEIGHBOR_TABLE \
      -setDestMacAction=$SET_DEST_MAC
