build:
	@go build -o bin/api

server: build
	@./bin/api
