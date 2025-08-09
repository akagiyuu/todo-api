#!/usr/bin/bash

OUTPUT="${1:-main}"

sqlc generate -f internal/database/sqlc.yml
go build -o $OUTPUT cmd/api/main.go
