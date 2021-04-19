#!/bin/bash
#
# Copyright 2020-present Open Networking Foundation
#
# SPDX-License-Identifier: Apache-2.0
#
#
set -x
source $1

p4rt-client -pushP4Info -server=$P4RUNTIME_ENDPOINT  -p4info=../$P4INFO_FILE
