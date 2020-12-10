sample: *.go go.mod
	gofmt -w *.go
	go build -o yamlApi main.go

.PHONY: init
init:
	go mod init yamlApi

.PHONY: test
test:
	gofmt -w *.go
	go test
