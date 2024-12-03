FROM golang:1.23.1

WORKDIR /src
# ワーキングディレクトリにソースコードをコピー
COPY ./ ./ 

# github.com/kakkky/pkgの依存関係をダウンロード
WORKDIR /src/pkg
RUN go mod download

# github.com/kakkky/appの依存関係をダウンロード
WORKDIR /src/app
RUN go mod download

# 必要なツールをインストール
RUN go install github.com/air-verse/air@latest
RUN go install go.uber.org/mock/mockgen@v0.3.0
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.23.0

# airを起動
WORKDIR /src
CMD ["air"]