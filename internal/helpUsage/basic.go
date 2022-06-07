/*
 * Copyright (c) 2022-present Intel Corporation All Rights Reserved
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package helpUsage

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("Help")

func BasicUsage() {
	usage := `
p4rt-client is used to configure the SAI layer of a SONiC based switch via the P4Runtime Service
Options:
p4rt-client	-pushP4Info          is used to push the p4info.txt file to the switch for command interpretation
p4rt-client	-addRouterInt        is used to create a virtual interface and map it to a physical interface
p4rt-client	-delRouterInt        is used to delete a virtual interface and map it to a physical interface
p4rt-client	-addNeighbor         is used to define an adjacent entity (switch, router, server etc..)
p4rt-client	-delNeighbor         is used to delete an adjacent entity (switch, router, server etc..)
p4rt-client	-addNextHop          is used to create a NextHop label for a interface & neighbor combination
p4rt-client	-delNextHop          is used to delete a NextHop label for a interface & neighbor combination
p4rt-client	-addIpV4             is used to create a route entry and point it to a NextHop
p4rt-client	-delIpV4             is used to delete a route entry and point it to a NextHop
p4rt-client	-addActionProfile    is used to join several NextHop entries into one entity for (E|U)cmp pathing
p4rt-client	-delActionProfile    is used to delete an ActionProfile entry
p4rt-client	-addIpV4Wcmp         is used to create a route entry and point it to an (E|U)cmp path
p4rt-client	-delIpV4Wcmp         is used to delete a route entry and point it to an (E|U)cmp path
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
`
	log.Info(usage)
}
