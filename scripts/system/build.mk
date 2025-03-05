deploy:
	go build -o /tmp/codegen-cli cmd/codegen-cli/main.go
	@sudo cp /tmp/codegen-cli /usr/local/bin/