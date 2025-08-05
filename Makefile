.PHONY: run test build clean

run:
	go run main.go

test:
	go test ./...

build:
	go build -o bin/app main.go

clean:
	rm -rf bin/