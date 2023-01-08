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
	GOOS=windows GOARCH=386 go build -o ${BIN_DIR}/${BIN_NAME}-windows-386 main.go

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