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

all: test clean build
