#!/bin/bash
set -x
SCRIPT_DIR="../base-scripts"
#only need to push p4info.txt once per switch
echo ../pushP4Info.sh `pwd`/vals-bcm-server.sh
echo ../pushP4Info.sh `pwd`/vals-wedge-server.sh 


for i in vals-bcm-server.sh  vals-bcm-wedge.sh  vals-wedge-bcm.sh  vals-wedge-server.sh 
do
   echo $SCRIPT_DIR/addRouterInterface.sh  `pwd`/$i
   echo $SCRIPT_DIR/addNeighbor.sh  `pwd`/$i
   echo $SCRIPT_DIR/addNexthop.sh `pwd`/$i	
   echo $SCRIPT_DIR/addIpV4Route.sh `pwd`/$i  
done

