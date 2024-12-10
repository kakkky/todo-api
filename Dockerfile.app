FROM golang:1.23.1

WORKDIR /src

COPY ./ ./

WORKDIR /src/app

RUN go mod download

# 必要なツールをインストール

RUN go install github.com/air-verse/air@latest
RUN go install go.uber.org/mock/mockgen@v0.3.0
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.23.0
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# airを起動
CMD ["air"]
