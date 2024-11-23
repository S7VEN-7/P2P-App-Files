build:
	@go build -o bin/p2p_app

run: build
	@./bin/p2p_app

test:
	@go test ./... -v