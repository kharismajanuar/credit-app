run:
	@cp .env.example .env
	@go build .
	@go run .