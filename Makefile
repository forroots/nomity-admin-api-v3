DB_HOST = localhost
DB_PORT = 5430
DB_NAME_DEFAULT = postgres
DB_USER_SUPER = takeo

DB_USER_APP = nomity_dev
DB_USER_PASSWORD_APP = nomity_dev
DB_NAME_APP = nomity_dev
DB_NAME_ATLAS_DEV = nomity_dev_atlas
DB_URL_APP = postgres://$(DB_USER_APP):$(DB_USER_PASSWORD_APP)@$(DB_HOST):$(DB_PORT)/$(DB_NAME_APP)?sslmode=disable
DB_URL_ATLAS_DEV = postgres://$(DB_USER_APP):$(DB_USER_PASSWORD_APP)@$(DB_HOST):$(DB_PORT)/$(DB_NAME_ATLAS_DEV)?sslmode=disable

DB_SCHEMA_FILE = ./db/schema.sql

###################### DB初期化関連 ######################
# アプリ用のDBを削除
.PHONY: drop-db-app
drop-db-app:
	psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER_SUPER) -d $(DB_NAME_DEFAULT) -c "DROP DATABASE $(DB_NAME_APP);"

# atlasの検証用データベースを削除
.PHONY: drop-db-atlas-dev
drop-db-atlas-dev:
	psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER_SUPER) -d $(DB_NAME_DEFAULT) -c "DROP DATABASE $(DB_NAME_ATLAS_DEV);"

# アプリ用のDBロールを削除
.PHONY: drop-db-user-app
drop-db-user-app:
	psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER_SUPER) -d $(DB_NAME_DEFAULT) -c "DROP ROLE $(DB_USER_APP);"

# アプリ用のDBロールを作成
.PHONY: create-db-user-app
create-db-user-app:
	psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER_SUPER) -d $(DB_NAME_DEFAULT) -c "CREATE ROLE $(DB_USER_APP) LOGIN PASSWORD '$(DB_USER_PASSWORD_APP)';"

# アプリ用のDBを作成
.PHONY: create-db-app
create-db-app:
	psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER_SUPER) -d $(DB_NAME_DEFAULT) -c "CREATE DATABASE $(DB_NAME_APP) OWNER $(DB_USER_APP);"

# atlasの検証用データベースの作成（スキーマのオーナーも変更する）
.PHONY: create-db-atlas-dev
create-db-atlas-dev:
	psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER_SUPER) -d $(DB_NAME_DEFAULT) -c "CREATE DATABASE $(DB_NAME_ATLAS_DEV) OWNER $(DB_USER_APP);"
	psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER_SUPER) -d $(DB_NAME_DEFAULT) -c "ALTER SCHEMA public OWNER TO $(DB_USER_APP);"


########################## DBバックアップ関連 ##########################
# アプリDBをバックアップ
.PHONY: backup-db-app
backup-db-app:
	mkdir -p ./tmp/db_backup
	pg_dump -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER_APP) -d $(DB_NAME_APP) -f ./tmp/db_backup/$(DB_NAME_APP)_backup.sql

# アプリDBのバックアップをリストア
.PHONY: restore-db-app
restore-db-app:
	psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER_APP) -d $(DB_NAME_APP) -f ./tmp/db_backup/$(DB_NAME_APP)_backup.sql

# アプリDBのスキーマをダンプ
.PHONY: dump-schema-db-app
dump-schema-db-app:
	mkdir -p ./tmp/db_backup
	pg_dump -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER_APP) -d $(DB_NAME_APP) --schema-only -f ./tmp/db_backup/$(DB_NAME_APP)_schema.sql

######################## データベースマイグレーション関連 ######################
# atlasで差分を取得（need:brew install pgformatter）
.PHONY: atlas-diff
atlas-diff:
	atlas schema diff --from "${DB_URL_APP}" --to file://${DB_SCHEMA_FILE} --dev-url "${DB_URL_ATLAS_DEV}" > ./tmp/atlas_diff.sql
	pg_format -i ./tmp/atlas_diff.sql

# atlasでマイグレーションを実行
.PHONY: atlas-apply
atlas-apply:
	atlas schema apply --url "${DB_URL_APP}" --file file://${DB_SCHEMA_FILE}  --dev-url "${DB_URL_ATLAS_DEV}"

# atlasでスキーマをinspect（sql形式で出力）
.PHONY: atlas-inspect-sql
atlas-inspect-sql:
	atlas schema inspect --url "${DB_URL_APP}" --format "{{ sql . }}" > ./tmp/atlas_inspect.sql
	pg_format -i ./tmp/atlas_inspect.sql

# atlasでスキーマをinspect（hcl形式で出力）
.PHONY: atlas-inspect-hcl
atlas-inspect-hcl:
	atlas schema inspect --url "${DB_URL_APP}" > ./tmp/atlas_inspect.hcl

###################### sqlc 関連 ######################
# sqlcのコード生成
.PHONY: sqlc
sqlc:
	sqlc generate -f ./tools/sqlc/sqlc.yaml
	go mod tidy

###################### air 関連 ######################
# air でサーバを起動
.PHONY: air
air:
	air -c ./tools/air/.air.toml


###################### docker向けビルド関連 ######################
# ビルド(dev & for linux)
# ※dockerの中からホストOS上のDBにアクセスするには、アプリ内のDSNのホスト情報を localhost から host.docker.internal に変更する必要があります。
.PHONY: build-dev
build-dev:
	mkdir -p ./docker/dev/app
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./docker/dev/app/dev_nomity_admin_api ./cmd/server/main.go
	cp ./config.dev.yaml ./docker/dev/app/config.dev.yaml
	echo "./docker/dev/app/config.dev.yamlのdatabase.hostをlocalhostからhost.docker.internalに変更してください"

# ビルド(test & for linux)
.PHONY: build-test
build-test:
	mkdir -p ./docker/test/app
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./docker/test/app/test_nomity_admin_api ./cmd/server/main.go
	cp ./config.test.yaml ./docker/test/app/config.test.yaml

# ビルド(prod & for linux)
.PHONY: build-prod
build-prod:
	mkdir -p ./build/prod_$$(date +%Y%m%d)/app
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/prod_$$(date +%Y%m%d)/app/prod_nomity_admin_api ./cmd/server/main.go
	cp ./configs/prod/* ./build/prod_$$(date +%Y%m%d)/app
	cp ./docker/prod/* ./build/prod_$$(date +%Y%m%d)


###################### サーバ作業関連 ######################
# サーバログイン（~/.ssh/configに「kensakit-ec-web-prod-01」が設定されていること）
.PHONY: ssh-web-server-with-ap
ssh-web-server-with-ap:
	ssh kensakit-ec-web-prod-01 -L 20022:163.43.158.30:10022

# ローカルのポート55432を本番DBに転送する
.PHONY: ssh-web-server-with-db
ssh-web-server-with-db:
	ssh kensakit-ec-web-prod-01 -L 55432:163.43.158.30:15432

# サーバアップロード（~/.ssh/configに「kensakit-ec-web-prod-01」が設定されていること）
.PHONY: upload-build-prod
upload-build-prod:
	rsync -avz ./build/prod_$$(date +%Y%m%d) kensakit-ec-web-prod-01:/home/ubuntu/apps/nomity-admin-api

########################## 本番DB関連 ##########################
PROD_DB_HOST = localhost
PROD_DB_PORT = 55432
PROD_DB_USER = nomity_admin
PROD_DB_USER_PASSWORD = chikuyo-diver-1110
PROD_DB_NAME = nomity_admin
PROD_DB_URL = postgres://$(PROD_DB_USER):$(PROD_DB_USER_PASSWORD)@$(PROD_DB_HOST):$(PROD_DB_PORT)/$(PROD_DB_NAME)?sslmode=disable

# 本番DBのバックアップ
.PHONY: prod-backup-db
prod-backup-db:
	mkdir -p ./tmp/prod_db_backup
	pg_dump -h $(PROD_DB_HOST) -p $(PROD_DB_PORT) -U $(PROD_DB_USER) -d $(PROD_DB_NAME) -f ./tmp/prod_db_backup/$(PROD_DB_NAME)_backup.sql

# 本番DBのバックアップをリストア
.PHONY: prod-restore-db
prod-restore-db:
	psql -h $(PROD_DB_HOST) -p $(PROD_DB_PORT) -U $(PROD_DB_USER) -d $(PROD_DB_NAME) -f ./tmp/prod_db_backup/$(PROD_DB_NAME)_backup.sql

# 本番DBのスキーマをダンプ
.PHONY: prod-dump-schema-db
prod-dump-schema-db:
	mkdir -p ./tmp/prod
	pg_dump -h $(PROD_DB_HOST) -p $(PROD_DB_PORT) -U $(PROD_DB_USER) -d $(PROD_DB_NAME) --schema-only -f ./tmp/prod/$(PROD_DB_NAME)_schema.sql

# 本番DBのスキーマをinspect（sql形式で出力）
.PHONY: prod-atlas-inspect-sql
prod-atlas-inspect-sql:
	atlas schema inspect --url "${PROD_DB_URL}" --format "{{ sql . }}" > ./tmp/prod/atlas_inspect_prod.sql
	pg_format -i ./tmp/prod/atlas_inspect_prod.sql

# 本番DBのスキーマをinspect（hcl形式で出力）
.PHONY: prod-atlas-inspect-hcl
prod-atlas-inspect-hcl:
	atlas schema inspect --url "${PROD_DB_URL}" > ./tmp/prod/atlas_inspect_prod.hcl

# 本番DBのスキーマをdiff（atlasの検証用データベースと比較）
.PHONY: prod-atlas-diff
prod-atlas-diff:
	mkdir -p ./tmp/prod
	atlas schema diff --from "${PROD_DB_URL}" --to file://${DB_SCHEMA_FILE} --dev-url "${DB_URL_ATLAS_DEV}" > ./tmp/prod/atlas_diff_prod.sql
	pg_format -i ./tmp/prod/atlas_diff_prod.sql

# 本番DBのスキーマをapply（atlasの検証用データベースに適用）
.PHONY: prod-atlas-apply
prod-atlas-apply:
	atlas schema apply --url "${PROD_DB_URL}" --file file://${DB_SCHEMA_FILE} --dev-url "${DB_URL_ATLAS_DEV}"
