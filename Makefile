.PHONY: build run

build:
	docker build -t crm .
run:
	docker run -p 8080:8080 crm
compose-up:
	docker-compose up
compose-down:
	docker-compose down
