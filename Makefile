clear:
	rm -rf ./bin
install:
	brew install vault
build: clear
	go build -o ./bin/ ./cmd/...
run: build
	./bin/$(APP)
test: build
	go test -v ./...
test-coverage: build
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
test-coverage-ci: build
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
lint: build
	golangci-lint run --fast
docker-start:
	docker-compose -f ./deployment/docker-compose.yml --env-file .dev.env up -d --build
docker-stop:
	docker-compose -f ./deployment/docker-compose.yml --env-file .dev.env down
docker-ps:
	docker-compose -f ./deployment/docker-compose.yml --env-file .dev.env ps
docker-logs:
	docker-compose -f ./deployment/docker-compose.yml --env-file .dev.env logs -f --tail=100
