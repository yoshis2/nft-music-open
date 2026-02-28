#!/bin/bash

if [ ! -e "go.mod" ]; then
    go mod init
    go get -u github.com/labstack/echo/v4

    go install github.com/joho/godotenv/cmd/godotenv@latest
    go install go.uber.org/mock/mockgen@latest
    go install github.com/go-delve/delve/cmd/dlv@latest
    go install github.com/rubenv/sql-migrate/...@latest
    go install github.com/air-verse/air@latest
    go install github.com/swaggo/swag/cmd/swag@latest
fi

if [ ! -e "vendor" ]; then
    go mod tidy
    go mod vendor
fi

if [ ! -e "../logs/access.log" ]; then
    mkdir -p ../logs
    touch ../logs/access.log
    echo "created access log file"
fi

swag init

air
