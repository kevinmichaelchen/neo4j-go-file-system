.PHONY: all
all:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose stop

.PHONY: rebuild
rebuild:
	docker-compose up --build
