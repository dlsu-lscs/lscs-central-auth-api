build:
	@go build -o bin/lscs-central-auth

run: build
	@./bin/lscs-central-auth

test:
	@go test -v ./...
