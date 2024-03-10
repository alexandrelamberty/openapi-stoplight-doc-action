MAIN_PACKAGE_PATH := ./
BINARY_NAME := stoplight-doc

## tidy: format code and tidy modfile
tidy:
	go fmt ./...
	go mod tidy -v

## build: build the application
build:
	go build -o=${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the application
run: build
	./${BINARY_NAME} -title=$(title) -file=$(file)
