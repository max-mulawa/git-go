.PHONY: build
build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/git-go ./main.go

test:
	go test ./...