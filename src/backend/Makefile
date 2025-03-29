build:
	docker build -t racket-rave .

run: build
	docker run -p 8080:8081 -d --name racket-container racket-rave:latest

stop:
	@docker stop racket-container

rm: stop
	@docker rm racket-container
	@docker image rm racket-rave:latest

run-local:
	go run cmd/main.go

logs:
	docker logs racket-container

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

.PHONY: run stop rm run-local logs lint tsp-install test