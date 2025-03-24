# ビルド
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# ソースコード全体をコピー
COPY . .

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o apiserver main.go

# 実行用の軽量イメージ
FROM alpine:latest

# 証明書インストール（HTTPS対応）
RUN apk --no-cache add ca-certificates

WORKDIR /app

# ビルドしたバイナリだけコピー
COPY --from=builder /app/apiserver .

EXPOSE 8080

# ヘルスチェック用
RUN apk --no-cache add wget

CMD ["./apiserver"]
