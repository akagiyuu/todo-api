#!/usr/bin/bash

OUTPUT="${1:-main}"

sqlc generate -f internal/database/sqlc.yml
swag init -d cmd/api,internal/handler
go build -o $OUTPUT cmd/api/main.go
