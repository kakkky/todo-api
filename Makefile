help: # コマンド確認
	@echo "\033[32mAvailable targets:\033[0m"
	@grep "^[a-zA-Z\-]*:" Makefile | grep -v "grep" | sed -e 's/^/make /' | sed -e 's/://'


##############
### テスト ####
##############

# テスト処理の共通化
# パラメータ:
# - path=<テスト対象のパス> (デフォルト: ./...)
# - opts=<追加オプション> (デフォルト: なし)
# - tags=<ビルドタグ> (デフォルト: なし)
define tests
	$(if $(TEST_TAGS),\
		go test -v -timeout 10m -tags=$(TEST_TAGS) $(TEST_PATH) $(TEST_OPTIONS),\
		go test -v -timeout 10m $(TEST_PATH) $(TEST_OPTIONS)\
	)
endef

# appディレクトリの全体テスト
# コマンド例: $ make test-app opts="-run TestXxx"
test-app:
	$(eval TEST_PATH=./...)
	$(eval TEST_TAGS=$(tags))
	$(eval TEST_OPTIONS=${opts})
	@echo "Running all tests in app..."
	cd ./app && $(call tests)

# ドメイン層のテスト
# コマンド例: $ make test-domain path=./... opts="-run TestXxx"
test-domain:
	$(eval TEST_PATH=$(or $(path),./...))
	$(eval TEST_TAGS=$(tags))
	$(eval TEST_OPTIONS=${opts})
	@echo "Running tests in domain..."
	cd ./app/domain && $(call tests)

# ユースケースのテスト
# コマンド例: $ make test-usecase path=./... opts="-run TestXxx"
test-usecase:
	$(eval TEST_PATH=$(or $(path),./...))
	$(eval TEST_TAGS=$(tags))
	$(eval TEST_OPTIONS=${opts})
	@echo "Running tests in usecase..."
	cd ./app/application/usecase && $(call tests)

# リポジトリのテスト
# コマンド例: $ make test-repo path=./... opts="-run TestXxx"
test-repo:
	$(eval TEST_PATH=$(or $(path),./...))
	$(eval TEST_TAGS=$(tags))
	$(eval TEST_OPTIONS=${opts})
	@echo "Running tests in repository..."
	cd ./app/infrastructure/repository_test && $(call tests)

# クエリサービスのテスト
# コマンド例: $ make test-query path=./... opts="-run TestXxx"
test-query:
	$(eval TEST_PATH=$(or $(path),./...))
	$(eval TEST_TAGS=$(tags))
	$(eval TEST_OPTIONS=${opts})
	@echo "Running tests in queryservice..."
	cd ./app/infrastructure/queryservice_test && $(call tests)

# pkgディレクトリのテスト
# コマンド例: $ make test-pkg path=./... opts="-run TestXxx"
test-pkg:
	$(eval TEST_PATH=$(or $(path),./...))
	$(eval TEST_TAGS=$(tags))
	$(eval TEST_OPTIONS=${opts})
	@echo "Running tests in pkg..."
	cd ./pkg && $(call tests)

test-integration:
	$(eval TEST_PATH=$(or $(path),./...))
	$(eval TEST_TAGS=$(tags))
	$(eval TEST_OPTIONS=${opts})
	@echo "Running tests in integration"
	cd ./app/infrastructure/api_test/integration && $(call tests) -tags=integration

#####################
##### コンテナ操作 ####
#####################

# docker-composeにおけるDockerfileのビルド
build:
	@echo "Building Docker images..."
	docker compose build 

# docker compose up
up:
	@echo "Starting containers with docker-compose up..."
	docker compose up

# docker compose down
down:
	@echo "Stopping containers with docker-compose down..."
	docker compose down

# docker compose logs -f
logs:
	@echo "Fetching logs with docker-compose logs..."
	docker compose logs -f

# docker compose ps
ps:
	@echo "Viewing running containers with docker-compose ps..."
	docker compose ps

ls:
	@echo "Viewing running containers with docker-compose ls..."
	docker container ls

exec-db:
	docker compose exec db /bin/bash

exec-kvs:
	docker compose exec kvs /bin/bash

########################
### DBマイグレーション ####
########################

MIGRATE_PATH = infrastructure/db/migrations
DB_URL = mysql://user:pswd@tcp(db:3306)/todo-db?parseTime=true

# マイグレーションファイルを作成
# コマンド例: $ make migrate-create name=<migration-name>
migrate-create:
	$(eval NAME=$(or $(name),$(error "Error: Please specify a migration name using name=<name>")))
	@echo "Creating migration file..."
	cd app/ && migrate create -ext sql -dir $(MIGRATE_PATH) -seq $(NAME)

# マイグレーションを適用
# コマンド例: $ make migrate-up
migrate-up:
	@echo "Applying migrations..."
	docker compose run --rm app migrate --path $(MIGRATE_PATH) --database "$(DB_URL)" -verbose up

# マイグレーションをロールバック
# コマンド例: $ make migrate-down
migrate-down:
	@echo "Rolling back migrations..."
	docker compose run --rm app migrate --path $(MIGRATE_PATH) --database "$(DB_URL)" -verbose down

# マイグレーションのバージョンをセット
# migration-upが失敗した時は、最後のバージョンにforceしなおしてから再度upするようにする
# コマンド例: $ make migrate-force version=<version_number>
migrate-force:
	@echo "Forcing database migration version to $(version)..."
	docker compose run --rm app migrate --path $(MIGRATE_PATH) --database "$(DB_URL)" force $(version)


#############
### sqlc ####
#############

# sqlcでコードを生成
sqlc-gen:
	@echo "Generating query in sql by sqlc..."
	cd ./app/infrastructure/db/sqlc && sqlc generate

#############
## swagger ##
#############

# コメントをパースしてドキュメント生成
swag:
	@echo "Generating document by swagger..."
	cd ./app && swag fmt && swag init

###################
### パッケージ管理 ###
###################

# github.com/kakkky/appにおいてgo getする
# コマンド例: $ make get-app name=github.com/xxxx/xxx
get-app:
	cd ./app && go get $(name)

get-pkg:
	cd ./pkg && go get $(name)

###################
### go generate ###
###################
go-gen:
	cd ./app && go generate ./...