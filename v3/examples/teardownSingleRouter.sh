#!/bin/bash

./p4rt-client -delIpV4 -server 10.70.2.3:9559 -routedNetwork=10.130.0.0/16 
./p4rt-client -delNextHop -server 10.70.2.3:9559  -nextHopId="internet"
./p4rt-client -delNeighbor -server 10.70.2.3:9559 -neighborName="fe80::207:43ff:fe4b:7f50" -routerInterface="Ethernet0"
./p4rt-client -delRouterInt -server 10.70.2.3:9559 -routerInterface="Ethernet0"


