.PHONY: all test build format docker_up docker_down

all: build test

build:
	go mod download
	go build -o ./bin/music_library

test:
	go test ./...

docker_up:
	sudo docker compose up -d --build
	sudo docker logs -tf music_library

docker_stop:
	sudo docker compose stop

docker_down:
	sudo docker compose down

rebuild: clean
	go mod tidy
	make build

format:
	go fmt ./...

clean:
	rm -rf ./bin
	rm -rf ./data