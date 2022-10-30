.PHONY: build
build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/git-go ./main.go

install: build
	sudo cp ./bin/git-go /usr/local/bin/
test:
	go test ./...