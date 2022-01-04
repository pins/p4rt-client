./p4rt-client -delIpV4 -server 10.70.2.3:9559 -routedNetwork=172.16.1.0/24 
./p4rt-client -server 10.70.2.3:9559 -delActionProfile -mpGroupId=group1


./p4rt-client -delNextHop -server 10.70.2.3:9559  -nextHopId="mp1"
./p4rt-client -delNextHop -server 10.70.2.3:9559  -nextHopId="mp2"
./p4rt-client -delNextHop -server 10.70.2.3:9559  -nextHopId="mp3"

./p4rt-client -delNeighbor -server 10.70.2.3:9559 -neighborName="fe80::207:43ff:fe4b:7f50" -routerInterface="Ethernet0"  
./p4rt-client -delNeighbor -server 10.70.2.3:9559 -neighborName="fe80::207:43ff:fe4b:7f50" -routerInterface="Ethernet8"  
./p4rt-client -delNeighbor -server 10.70.2.3:9559 -neighborName="fe80::207:43ff:fe4b:7f50" -routerInterface="Ethernet16"

./p4rt-client -delRouterInt -server 10.70.2.3:9559 -routerInterface="Ethernet0"
./p4rt-client -delRouterInt -server 10.70.2.3:9559 -routerInterface="Ethernet8"
./p4rt-client -delRouterInt -server 10.70.2.3:9559 -routerInterface="Ethernet16"
