##############################
 #
 # Copyright (c) 2022-present Intel Corporation All Rights Reserved
 # Copyright 2020-present Open Networking Foundation
 #
 # SPDX-License-Identifier: Apache-2.0
 #
##############################

BINARY_NAME=p4rt-client
OUTPUT_DIR=bin

build:
		go build -o bin/ ./...

run:
		./bin/${BINARY_NAME}
test:
		go test -coverprofile=c.out  ./...

build_and_run: build run

clean:
		go clean
			rm ${OUTPUT_DIR}/*
