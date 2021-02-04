#!/bin/bash
set -x
source $1
source ../vals-p4.sh


p4rt-client -addActionProfile \
      -server=$P4RUNTIME_ENDPOINT  \
      -nextHopWeights=$NEXTHOP_LIST \
      -mpGroupId=$WCMP_GROUP_NAME \
      -aProfileId=$WCMP_GROUP_TABLE  \
      -setNextHopAction=$SET_NEXTHOP

