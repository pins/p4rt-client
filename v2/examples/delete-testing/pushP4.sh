#!/bin/bash
#
# Copyright 2020-present Open Networking Foundation
#
# SPDX-License-Identifier: Apache-2.0
#
#

p4rt-client -pushP4Info -server=10.70.2.2:9559 -p4info=middleblock.p4info.pb.txt
