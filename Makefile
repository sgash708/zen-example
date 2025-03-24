.PHONY: build run clean d-build d-run d-stop up down

# ローカル開発用
build:
	go build -o ./bin/apiserver main.go

run: build
	./bin/apiserver

clean:
	rm -rf ./bin

# Docker ビルド・実行用
d-build:
	docker build -t zen-example-api docker/go/

d-run:
	docker run -p 8080:8080 zen-example-api

d-stop:
	docker stop $$(docker ps -q --filter ancestor=zen-example-api)

# Docker Compose 用
up:
	docker compose up --build -d

down:
	docker compose down

# 依存関係インストール
deps:
	go mod tidy
	go mod download

# テスト実行
test:
	go test -v ./...

# リント実行
lint:
	go vet ./...

# ヘルプ
help:
	@echo "利用可能なコマンド:"
	@echo "  make build            - アプリケーションをビルド"
	@echo "  make run              - アプリケーションを実行"
	@echo "  make clean            - ビルド成果物を削除"
	@echo "  make docker-build     - Dockerイメージをビルド"
	@echo "  make docker-run       - Dockerコンテナを実行"
	@echo "  make docker-stop      - 実行中のコンテナを停止"
	@echo "  make docker-compose-up - Docker Composeでサービスを起動"
	@echo "  make docker-compose-down - Docker Composeでサービスを停止"
	@echo "  make deps             - 依存関係をインストール"
	@echo "  make test             - テストを実行"
	@echo "  make lint             - リントを実行"
