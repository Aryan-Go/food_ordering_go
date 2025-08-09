hello:
	@echo "Hello"

test:
	 go test -v ./package/middlewares

run:
	go run command/main.go