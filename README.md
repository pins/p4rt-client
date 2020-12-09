# p4rt-client
## Purpose
`p4rt-client` was developed to prototype behavior on a SONiC based switch via P4Runtime interface.  There is no provison for handling "packet-in", so while it can configure the switch, it is not a "runtime" controller.  
`macToIpV6` is a simple utility that takes a single argument, MAC Addr, in the format `00:11:22:33:44:55` and prints to STDOUT the IPv6 Link Local Address 

## Installation
To use `go get` you need to first set up git to use git@github.com: syntax instead of https://github  
`git config --global url.git@github.com:.insteadOf https://github.com/`  
Assuming your github credidentials are setup and you have a functional golang environment you should be able to use:  
1. `go get github.com/pins/p4rt-client`
2. `go get github.com/pins/p4rt-client/utils/macToIpV6`
---

## Usage 
```
$ p4rt-client --help  
2020/12/09 08:23:58   
p4rt-client is used to configure the SAI layer of a SONiC based switch via the P4Runtime Service  
Options:  
p4rt-client	-pushP4Info          is used to push the p4info.txt file to the switch for command interpretation  
p4rt-client	-addRouterInt        is used to create a virtual interface and map it to a physical interface  
p4rt-client	-addNeighbor         is used to define an adjacent entity (switch, router, server etc..)  
p4rt-client	-addNextHop          is used to create a NextHop label for a interface & neighbor combination  
p4rt-client	-addIpV4             is used to create a route entry and point it to a NextHop   
p4rt-client	-addActionProfile    is used to join several NextHop entries into one entity for (E|U)cmp pathing  
p4rt-client	-addIpV4Wcmp         is used to create a route entry and point it to an (E|U)cmp path  
p4rt-client	-help                prints this message  

global options:  
-debug              : generated detailed debugging info  
-logfile=$LOG_FILE  : direct output to file instead of stdout  
  
For help on an individual option include the option with -help e.g.  
p4rt-client -pushP4Info -help  
  
To see list of available arguments  
p4rt-client -h  
  
To see instructions for multiple Options in a single invocation:  
p4rt-client -help -advanced  
```
## Examples
See the Examples directory for different use cases
