# Zenフレームワークを使用したDDDアーキテクチャの例

このプロジェクトは[Unkey Zenフレームワーク](https://github.com/unkeyed/unkey/tree/main/go/pkg/zen)を使用して、ドメイン駆動設計（DDD）の原則に従ったGoアプリケーション（Go 1.24）のサンプルです。

## アーキテクチャ

このプロジェクトは標準的なDDD階層アーキテクチャに従っています：

1. **ドメイン層**: コアビジネスロジックとエンティティ
   - エンティティ: User
   - リポジトリインターフェース
   - ドメインサービス
   
2. **アプリケーション層**: ユースケースを調整するアプリケーションサービス
   - 入出力用DTO
   - アプリケーションサービス
   
3. **インフラストラクチャ層**: 技術的な実装
   - リポジトリ実装（メモリストレージ）
   
4. **ハンドラ層**: HTTPリクエストハンドラ
   - REST APIルート
   - リクエスト/レスポンスのバインディング

## 始め方

### ローカル実行

1. 依存関係のインストール:
```
make deps
```

2. サーバーの起動:
```
make run
```

### Docker実行

1. Dockerイメージのビルドと起動:
```
make up
```

2. 停止:
```
make down
```

### 利用可能なMakeコマンド

```
make build            # アプリケーションをビルド
make run              # アプリケーションを実行
make clean            # ビルド成果物を削除
make d-build          # Dockerイメージをビルド
make d-run            # Dockerコンテナを実行
make d-stop           # 実行中のコンテナを停止
make up               # Docker Composeでサービスを起動
make down             # Docker Composeでサービスを停止
make deps             # 依存関係をインストール
make test             # テストを実行
make lint             # リントを実行
make help             # ヘルプメッセージを表示
```

## APIエンドポイント

サーバーはポート8080で起動し、以下のエンドポイントが利用可能です:
- GET /hello - シンプルな "Hello, world!" メッセージを返します
- POST /users - 名前、メール、パスワードの検証付きでユーザーを作成します

## API使用例

### Hello World
```
curl http://localhost:8080/hello
```

### ユーザー作成
```
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"山田太郎","email":"yamada@example.com","password":"password123"}'
```