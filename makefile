build:
	@go build -o bin/MeChat

run: build
	@./bin/MeChat

test:
	@go test -v ./...
