SHELL := /bin/bash

# ================================================================
# GO

go-run:
	go run cmd/services/api/main.go

go-tidy:
	go mod tidy
	go mod vendor
