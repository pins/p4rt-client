#!/bin/bash
set -x

#only need to push p4info.txt once per switch
./pushP4Info.sh vals-bcm-server.sh
./pushP4Info.sh vals-wedge-server.sh 


for i in vals-bcm-server.sh  vals-bcm-wedge.sh  vals-wedge-bcm.sh  vals-wedge-server.sh 
do
   ./addRouterInterface.sh  $i
   ./addNeighbor.sh  $i
   ./addNexthop.sh $i	
   ./addIpV4Route.sh $i  
done

