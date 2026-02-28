FROM golang:1.25-alpine

ENV GOPATH=/go
ENV GOBIN=${GOPATH}/bin
ENV PATH=${GOPATH}/bin:${PATH}

# フォルダ構成をGoPathの通りに設置
WORKDIR /go/src/nft-music/backend

RUN apk update && apk upgrade && \
    apk add --no-cache bash musl-dev git build-base curl

RUN curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.8.0

RUN go install github.com/joho/godotenv/cmd/godotenv@latest && \
    go install go.uber.org/mock/mockgen@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/air-verse/air@latest && \
    go install github.com/rubenv/sql-migrate/...@latest && \
    go install mvdan.cc/gofumpt@latest && \
    go install github.com/josharian/impl@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    go install github.com/ethereum/go-ethereum/cmd/abigen@latest

RUN go mod init

RUN go get -u github.com/labstack/echo/v4

RUN go mod vendor
