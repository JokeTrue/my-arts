lint:
	command -v golangci-lint >/dev/null 2>&1 || { go install github.com/golangci/golangci-lint/cmd/golangci-lint; } && \
    golangci-lint --config=./.golangci.yaml run ./...

build:
	go build -o .bin/my-arts ./cmd/main.go

test:
	go test -race -count 100 ./pkg/...

run:
	docker-compose -f ./docker-compose.yaml up -d

down:
	docker-compose -f ./docker-compose.yaml down