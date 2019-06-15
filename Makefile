build:
	go build -ldflags "-s -w" -mod=vendor

run:
	go run -mod=vendor main.go

clean:
	go clean -mod=vendor

check:
	go fmt
	go vet
