BINARY_NAME=gcal-cli

all: build

build:
	go build -o ${BINARY_NAME} .

fmt:
	go fmt ./...

test:
	go test -v ./...

clean:
	rm -f ${BINARY_NAME}
