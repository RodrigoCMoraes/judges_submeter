APP_NAME=judges_submeter
PORT=10080

build:
	mkdir -p ./bin && go build -o ./bin/runner ./cmd/${APP_NAME}

run:
	go run ./cmd/${APP_NAME}/main.go

get-dependencies:
	go mod tidy

format:
	go fmt ./...

lint:
	golangci-lint run

build-image:
	DOCKER_BUILDKIT=0 docker build -t ${APP_NAME} .

run-image:
	docker run --rm -it -p ${PORT}:${PORT} --env-file=./.env ${APP_NAME}

tests:
	go test -v ./test