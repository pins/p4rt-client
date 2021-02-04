#!/bin/bash
set -x
source $1

p4rt-client -pushP4Info -server=$P4RUNTIME_ENDPOINT  -p4info=../$P4INFO_FILE
