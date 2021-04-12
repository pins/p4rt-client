#!/bin/bash
set -x 
WEG_MAC=00:90:fb:64:cc:9e
AR1_MAC=fc:bd:67:2b:4b:08
AR2_MAC=fc:bd:67:2b:c8:f8
AS_MAC=8c:ea:1b:17:64:0c
p4rt-client -addRouterInt -server=10.70.2.2:9559  -routerInterface=intf-eth29 -routerPortId=112 -routerIntMAC=$AS_MAC -egressPort=112 routerTable=33554497  -setPortMac=16777218
p4rt-client -addNeighbor  -server=10.70.2.2:9559 -neighborName=test-neighbor1 -destMAC=$AR2_MAC -routerInterface=intf-eth29
p4rt-client -addNextHop   -server=10.70.2.2:9559 -routerInterface=intf-eth29 -neighborName=test-neighbor1 -nextHopId=bcmInter 
p4rt-client -addIpV4      -server=10.70.2.2:9559 -vrf=vrf-0 -routedNetwork=172.16.1.0/24 -nextHopId=bcmInter 

