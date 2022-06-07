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
