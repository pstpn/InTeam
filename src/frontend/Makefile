run:
	docker compose up -d

stop:
	@docker stop racket-front

rm: stop
	@docker rm racket-front
	@docker image rm frontend-racket-front:latest

run-local:
	npm run serve

logs:
	docker logs racket-front