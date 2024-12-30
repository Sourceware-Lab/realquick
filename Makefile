export
UID="$(shell id -u)"
GID="$(shell id -g)"
DOCKER_BUILDKIT=1

.ONESHELL:

run: down
	docker compose -f docker-compose.yml -f otel.docker-compose.yml up --remove-orphans --build frontend backend

run_local: down
	docker compose run -d --remove-orphans -p 5432:5432 postgres
	cd ./src/backend
	air

test: down
	docker compose -f ./docker-compose.yml -f ./test.docker-compose.yml up --abort-on-container-exit --remove-orphans --build

test_local: down
	docker compose run -d --remove-orphans -p 5432:5432 postgres
	cd src/backend
	go test -race ./...
	./fuzz.sh

prod: down
	docker compose -f ./docker-compose.yml -f otel.docker-compose.yml -f ./prod.docker-compose.yml up --remove-orphans --build

lint:
	cd src/backend
	go vet
	go fmt
	golangci-lint run --fix ./...

clean: down kill
	docker system prune -a -f
	docker volume prune -a -f
	docker network prune -f
	docker system prune --volumes --all

kill:
	docker compose down --remove-orphans
	- docker ps -q | xargs -r docker kill
	- docker stop `docker ps -qa`

down:
	docker compose down --remove-orphans

generate_sdk: down
	docker compose -f ./docker-compose.yml -p realquick up -d typescript_sdk_generator

