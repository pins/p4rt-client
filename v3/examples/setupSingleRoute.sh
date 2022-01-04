#!/bin/bash

./p4rt-client -pushP4Info -server=10.70.2.3:9559  -p4info=middle-block.txt
./p4rt-client -addRouterInt -server 10.70.2.3:9559 -routerInterface="Ethernet0" -routerPortId=1000 -routerIntMAC="fc:bd:67:2b:4b:08"  -egressPort="Ethernet0"
./p4rt-client -addNeighbor -server 10.70.2.3:9559 -neighborName="fe80::207:43ff:fe4b:7f50" -routerInterface="Ethernet0" -destMAC="00:07:43:4b:7f:50"
./p4rt-client -addNextHop -server 10.70.2.3:9559 -routerInterface="Ethernet0" -neighborName="fe80::207:43ff:fe4b:7f50" -nextHopId="internet"
./p4rt-client -addIpV4 -server 10.70.2.3:9559  -routedNetwork=10.130.0.0/16 -nextHopId="internet"


