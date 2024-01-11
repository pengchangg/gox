all: test clean build

BINARY_NAME=gox

test:
	go test ./...

test-redis:
	go run redis/main.go

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/${BINARY_NAME}-windows main.go

run:
	go build -o ./bin/${BINARY_NAME} main.go
	./bin/${BINARY_NAME}

clean:
	go clean
	rm -rf ./bin

build-linux-amd64:
	mkdir -p build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/${BINARY_NAME}_linux_amd64 main.go

build-linux-arm64:
	mkdir -p build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/${BINARY_NAME}_linux_arm64 main.go
