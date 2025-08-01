#!/usr/bin/bash

OUTPUT="${1:-main}"

sqlc generate -f internal/database/sqlc.yml
swag init -d cmd/api,internal/handler --parseDependency --parseInternal
go build -o $OUTPUT cmd/api/main.go
