run r:
	@echo [Running Mini URL Service]
	@DEFAULT_URL_REDIRECT=https://www.apple.com MONGODB_URI="mongodb://root:password@localhost:27017/marketingDB?authSource=admin" JWT_TOKEN_SEED=supersecret REDIS_HOST=localhost:6379 APP_PORT=8081 APP_ENV=dev APP_VERSION=0.1 go run src/main.go

test t:
	@echo [Running Mini URL Service Tests]
	@go test ./src/...

.PHONY: run r