.PHONY: api events-app

infrastructure:
	@docker-compose up -d nats cockroachdb 

api:
	@docker-compose up --build --abort-on-container-exit --exit-code-from api api

events-listener:
	@docker-compose up --build --abort-on-container-exit --exit-code-from events-listener events-listener

run-all: api events-app infrastructure