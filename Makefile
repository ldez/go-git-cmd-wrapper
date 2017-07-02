.PHONY: all

default: generate test-unit

generate:
	go generate -x internal/generator.go

test-unit:
	go test ./... --cover