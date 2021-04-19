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

p4rt-client  -addRouterInt \
	-server=$P4RUNTIME_ENDPOINT \
	-routerInterface=$INTERFACE_NAME \
	-routerPortId=$ROUTER_PORT_ID \
	-routerIntMAC=$ROUTER_INTF_MAC  \
	-egressPort=$EGRESS_PORT 
