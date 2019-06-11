hooks:
	@cp -f hooks/pre-commit  .git/hooks/pre-commit

fmt:
	@go fmt ./...

vet:
	@go vet ./...

lint:
	@golangci-lint run ./...

agent:
	@go build -o agent main.go



#add other here


#here here


