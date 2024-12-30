FROM golang:1.23.1

WORKDIR /src

COPY ./app/go.mod ./app/go.mod
COPY ./app/go.sum ./app/go.sum
WORKDIR /src/app
# github.com/[あなたの名前]/appモジュールの依存関係を全てコンテナ内にダウンロード
RUN go mod download
# srcディレクトリ配下にローカルのソースコードをコピー
COPY ./ ./

WORKDIR /src/app

# 必要なツールをインストール

RUN go install github.com/air-verse/air@latest
RUN go install go.uber.org/mock/mockgen@v0.3.0
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.23.0
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# airを起動
CMD ["air"]
