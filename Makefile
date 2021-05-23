.PHONY: generate lint lint-fix test

default: generate lint test

generate:
	go generate -x internal/generator.go

clean-generate:
	 rm ./**/*_gen.go

test:
	go test ./... --cover

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix
