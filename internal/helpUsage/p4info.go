/*
 * Copyright (c) 2022-present Intel Corporation All Rights Reserved
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package helpUsage

func PushP4Usage() {
	usage := `
Usage:
./p4rt-client -pushP4info -p4info=$P4_INFO_FILENAME`
	log.Info(usage)
}
