run:
	docker compose up -d

stop:
	@docker stop racket-front

down:
	@docker compose down

rm: down
	@docker image rm frontend-racket-front:latest

logs:
	docker logs racket-front