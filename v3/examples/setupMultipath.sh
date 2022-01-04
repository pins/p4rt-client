#!/bin/bash

./p4rt-client -pushP4Info -server=10.70.2.3:9559  -p4info=middle-block.txt
./p4rt-client -addRouterInt -server 10.70.2.3:9559 -routerInterface="Ethernet0" -routerPortId=1000 -routerIntMAC="fc:bd:67:2b:4b:08"  -egressPort="Ethernet0"
./p4rt-client -addRouterInt -server 10.70.2.3:9559 -routerInterface="Ethernet8" -routerPortId=1001 -routerIntMAC="fc:bd:67:2b:4b:08"  -egressPort="Ethernet8"
./p4rt-client -addRouterInt -server 10.70.2.3:9559 -routerInterface="Ethernet16" -routerPortId=1002 -routerIntMAC="fc:bd:67:2b:4b:08"  -egressPort="Ethernet16"

./p4rt-client -addNeighbor -server 10.70.2.3:9559 -neighborName="fe80::207:43ff:fe4b:7f50" -routerInterface="Ethernet0"  -destMAC="00:07:43:4b:7f:50"
./p4rt-client -addNeighbor -server 10.70.2.3:9559 -neighborName="fe80::207:43ff:fe4b:7f50" -routerInterface="Ethernet8"  -destMAC="00:07:43:4b:7f:50"
./p4rt-client -addNeighbor -server 10.70.2.3:9559 -neighborName="fe80::207:43ff:fe4b:7f50" -routerInterface="Ethernet16" -destMAC="00:07:43:4b:7f:50"

./p4rt-client -addNextHop -server 10.70.2.3:9559 -routerInterface="Ethernet0" -neighborName="fe80::207:43ff:fe4b:7f50" -nextHopId="mp1"
./p4rt-client -addNextHop -server 10.70.2.3:9559 -routerInterface="Ethernet8" -neighborName="fe80::207:43ff:fe4b:7f50" -nextHopId="mp2"
./p4rt-client -addNextHop -server 10.70.2.3:9559 -routerInterface="Ethernet16" -neighborName="fe80::207:43ff:fe4b:7f50" -nextHopId="mp3"

nexthopList=mp1:100,mp2:105,mp3:110
./p4rt-client -server 10.70.2.3:9559 -addActionProfile -mpGroupId=group1  -nextHopWeights=$nexthopList 

./p4rt-client -addIpV4Wcmp -server 10.70.2.3:9559 -routedNetwork="172.16.1.0/24" -mpGroupId=group1




