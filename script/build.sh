#!/usr/bin/bash

OUTPUT="${1:-main}"

sqlc generate
go build -o $OUTPUT cmd/api/main.go
