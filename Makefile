ENTRYPOINT_NAME=transactions-api

# To generate swagger you must have swagger cli so use `go get -u github.com/swaggo/swag/cmd/swag`
documentation:
	@swag init -g rest/router.go

test:
	@go test -cover ./service/... ./rest/handler/...

test-report:
	@go test -covermode=count -coverprofile=report.out ./service/... ./rest/handler/...
	@go tool cover -html=report.out

clean:
	@rm -rf bin/

build: clean
	@go build -o bin/${ENTRYPOINT_NAME} cmd/main.go

run:
	@go run cmd/main.go

docker-clean:
	@docker-compose down
	@docker rmi everest/${ENTRYPOINT_NAME}

docker-build:
	@docker build -t everest/${ENTRYPOINT_NAME} .

docker-run:
	@docker-compose up -d --force-recreate --remove-orphans
