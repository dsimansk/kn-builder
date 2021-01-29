build:
	gofmt -s -w knb.go pkg
	go build -o knb
.PHONY: build