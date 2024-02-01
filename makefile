build:
	@go build -o otpverification cmd/main.go

run: build
	@./otpverification

start:
	@go run cmd/main.go