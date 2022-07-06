FROM golang:alpine3.16

RUN apk add --no-cache git \
 && go install github.com/go-delve/delve/cmd/dlv@latest \
 && go install github.com/cosmtrek/air@latest

WORKDIR /app/file-share/server

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD air -c .air.toml
