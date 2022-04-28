BINARY_NAME=draw
INSTALL_DIR="/usr/bin/"

SRC=$(wildcard *.go)

all: draw

draw: ${SRC}
	go build -o "${BINARY_NAME}"

install: draw
	cp "${BINARY_NAME}" "${INSTALL_DIR}${BINARY_NAME}"

run: draw
	./${BINARY_NAME}

clean:
	@go clean
	@rm -f "${BINARY_NAME}"

.PHONY: all clean run install
