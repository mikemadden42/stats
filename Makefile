build:
	go build -ldflags "-s -w"

run:
	go run main.go

clean:
	go clean

check:
	go fmt
	go vet