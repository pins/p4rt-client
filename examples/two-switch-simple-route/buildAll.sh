#!/bin/bash
set -x
SCRIPT_DIR="../base-scripts"
#only need to push p4info.txt once per switch
$SCRIPT_DIR/pushP4Info.sh `pwd`/vals-bcm-server.sh
#$SCRIPT_DIR/pushP4Info.sh `pwd`/vals-wedge-server.sh 


#for i in vals-bcm-server.sh  vals-bcm-wedge.sh  vals-wedge-bcm.sh  vals-wedge-server.sh 
for i in vals-bcm-server.sh  vals-bcm-wedge.sh 
do
   $SCRIPT_DIR/addRouterInterface.sh  `pwd`/$i
   $SCRIPT_DIR/addNeighbor.sh  `pwd`/$i
   $SCRIPT_DIR/addNexthop.sh `pwd`/$i	
   $SCRIPT_DIR/addIpV4Route.sh `pwd`/$i  
done

