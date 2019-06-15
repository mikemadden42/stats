build:
	go build -ldflags "-s -w" -mod=vendor

run:
	go run main.go -mod=vendor

clean:
	go clean -mod=vendor

check:
	go fmt
	go vet
