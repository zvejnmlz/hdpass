BINARY_NAME=hdpass.bin

build:
	go build -o bin/${BINARY_NAME} cmd/webserver/main.go

run:
	go build -o bin/${BINARY_NAME} cmd/webserver/main.go
	./bin/${BINARY_NAME}

clean:
	go clean
	rm bin/${BINARY_NAME}