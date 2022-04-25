.PHONY: api events-app

include .env
$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

infrastructure:
	@docker-compose up -d nats cockroachdb

api:
	@docker-compose up --build --abort-on-container-exit --exit-code-from api api

events-listener:
	@docker-compose up --build --abort-on-container-exit --exit-code-from events-listener events-listener

run-all: api events-app infrastructure

test:
	go test -v ./...