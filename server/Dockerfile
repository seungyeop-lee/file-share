FROM golang:1.15-alpine

RUN apk add --no-cache git \
 && go get github.com/go-delve/delve/cmd/dlv \
 && go get github.com/cosmtrek/air

WORKDIR /app/file-share/server

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD air -c .air.toml