# 【Go】 WEBフレームワークなしで開発する TODO API ハンズオン
このリポジトリーは、Zenn書籍で扱っているハンズオンのソースコードです。

URL ▶︎ https://zenn.dev/yuta_kakiki/books/08227a2ca0d290

# バージョン
```
go version go1.23.1 darwin/arm64
```

# エンドポイント
| HTTPメソッド | エンドポイント       | 説明                        |
|--------------|----------------------|-----------------------------|
| POST         | /users               | 新しいユーザーを作成する      |
| DELETE       | /users/me               | カレントユーザーを削除する           |
| GET          | /users              | 全てのユーザーを取得する      |
| PATCH        | /users/me               | カレントユーザーのプロフィールを更新する |
| POST         | /tasks              | 新しいタスクを作成する        |
| DELETE       | /tasks/{id}         | 指定したIDのタスクを削除する   |
| PATCH        | /tasks/{id}/state         | 指定したIDのタスクの状態を更新する |
| GET          | /tasks/{id}         | 指定したIDのタスクを取得する   |
| GET          | /tasks              | 全てのタスクを取得する        |
| GET          | /users/me/tasks         | ユーザーに紐づくタスクを取得する |
| POST         | /login              | ログインを行う        |
| DELETE       | /logout             | ログアウトを行う      |

# コマンド（一部）
Makefileにタスクを定義しています。

## Docker関連

### コンテナビルド
```
make build
```

### コンテナ起動
同時に、プログラムが実行されます。
```
make up
```

### コンテナ削除
```
make down
```

## パッケージのインストール

### appパッケージにインストール
```
make get-app name="xxxxxxxx"
```

### pkgパッケージにインストール
```
make get-pkg name="xxxxxxxx"
```

## DBマイグレーション

### マイグレーションファイルの生成
```
make migrate-create name="xxxxxxxx"
```
### マイグレーションを適用
```
make migrate-up
```

### マイグレーションをロールバック
```
make migrate-down
```

## テスト
### appパッケージ全体の単体テストを実行
```
make test-app
```
### pkgパッケージ全体の単体テストを実行
```
make test-pkg
```

### リポジトリのテスト
```
make test-repo
```

### エンドポイントの統合テストを実行
```
make test-integration
```
ゴールデンファイルを生成する場合は`-update`フラグをつけます
```
make test-integration -update
```



