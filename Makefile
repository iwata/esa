GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

.PHONY: test
test:
	@go test -race -v $(GOPACKAGES)
