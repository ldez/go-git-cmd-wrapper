.PHONY: generate check test fmt

default: generate check test

generate:
	go generate -x internal/generator.go

test:
	go test ./... --cover

check:
	golangci-lint run

fmt:
	gofmt -s -l -w .