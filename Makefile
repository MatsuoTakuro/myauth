run_on_local:
	go run main.go

build:
	docker compose build

run:
	docker compose up

rm:
	docker compose rm app -f
