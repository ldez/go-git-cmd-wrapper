.PHONY: all

default: generate checks test-unit

generate:
	go generate -x internal/generator.go

test-unit:
	go test ./... --cover

checks:
	golangci-lint run
