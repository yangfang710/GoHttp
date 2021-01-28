run:
	go run cmd/main.go

test:
	CGO_ENABLED=1 APP_ENV=test GO111MODULE=on go test -count=1 -race -v -cover ./...

build:
	mkdir -p bin
	rm -fr bin/server
	CGO_ENABLED=0 GO111MODULE=on go build -o bin/server cmd/main.go

generate:
	APP_ENV=test go generate ./...

doc:
	apidoc -i handler/ -o apidoc/

.PHONY: test build lint
lint:
	golangci-lint run ./...