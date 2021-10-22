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


p4rt-client -addActionProfile \
      -server=$P4RUNTIME_ENDPOINT  \
      -nextHopWeights=$NEXTHOP_LIST \
      -mpGroupId=$WCMP_GROUP_NAME 
