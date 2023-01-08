BIN_NAME=translator-cli
BIN_DIR=bin

.PHONY: build
build:
	go build -o ${BIN_DIR}/${BIN_NAME} main.go

.PHONY: compile
compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o ${BIN_DIR}/${BIN_NAME}-freebsd-386 main.go

	GOOS=linux GOARCH=386 go build -o ${BIN_DIR}/${BIN_NAME}-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o ${BIN_DIR}/${BIN_NAME}-linux-amd64 main.go
	GOOS=linux GOARCH=arm go build -o ${BIN_DIR}/${BIN_NAME}-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o ${BIN_DIR}/${BIN_NAME}-linux-arm64 main.go

	GOOS=windows GOARCH=386 go build -o ${BIN_DIR}/${BIN_NAME}-windows-386 main.go
	GOOS=windows GOARCH=amd64 go build -o ${BIN_DIR}/${BIN_NAME}-windows-amd64 main.go
	GOOS=windows GOARCH=arm go build -o ${BIN_DIR}/${BIN_NAME}-windows-arm main.go

	GOOS=darwin GOARCH=amd64 go build -o ${BIN_DIR}/${BIN_NAME}-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o ${BIN_DIR}/${BIN_NAME}-darwin-arm64 main.go

.PHONY: run
run:
	./${BIN_DIR}/${BIN_NAME} main.go

.PHONY: test
test:
	go test -v ./**/*.go

.PHONY: test_coverage
test_coverage:
	go test ./ -coverprofile=coverage.out

.PHONY: clean
clean:
	go clean
	rm -rf ${BIN_DIR}/*

.PHONY: dep
dep:
	go mod download

.PHONY: vet
vet:
	go vet

.PHONY: all
all: build run