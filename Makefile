.PHONY: run
run: 
	go run cmd/app/main.go 

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: mod
mod:
	go mod tidy

.PHONY: gen
gen:
	oapi-codegen --package server --generate types,gin ./api/swagger.yaml > ./internal/server/swagger.gen.go
	go fmt ./internal/server/swagger.gen.go

.PHONY: build
build:
	fmt
	go build cmd/app -o auth
