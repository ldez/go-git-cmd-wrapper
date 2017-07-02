.PHONY: all

default: generate

generate:
	go generate -x internal/generator.go
