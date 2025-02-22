run:
	go run ./...

lint:
	golangci-lint run

tsp-install:
	cd typespec && tsp install

tsp-compile: tsp-install
	cd typespec && tsp compile . --output-dir ../docs

code-gen:
	go run github.com/ogen-go/ogen/cmd/ogen@latest --target internal/router/ogen/ --clean docs/@typespec/openapi3/openapi.v1.yaml

gen: tsp-compile code-gen

test:
	go test --race --count=1 ./...

.PHONY: run lint tsp-install test