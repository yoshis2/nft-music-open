# ==============================================================================
# Docker / 全体操作 (Global)
# ==============================================================================

# Dockerイメージをビルドします
.PHONY: build
build:
	docker compose build

# Dockerコンテナを起動します
.PHONY: up
up:
	docker compose up frontend backend mysql mysql-test ganache ipfs

# Dockerコンテナを停止・削除します
.PHONY: down
down:
	docker compose down

# Dockerコンテナの状態を表示します
.PHONY: ps
ps:
	docker compose ps -a

# ==============================================================================
# フロントエンド (Frontend)
# ==============================================================================

# フロントエンドのコンテナ内でbashを起動します
.PHONY: front
front:
	docker compose exec frontend bash

# フロントエンドのコンテナを起動し、bashを実行します（一時的なコンテナ）
.PHONY: front-run
front-run:
	docker compose run --rm frontend bash

# フロントエンドのDockerイメージをビルドします
.PHONY: front-build
front-build:
	docker compose build frontend

# フロントエンドのDockerイメージをビルドします (front-buildのエイリアス)
.PHONY: build-front
build-front:
	docker compose build frontend

# フロントエンドのリントチェックを実行します
.PHONY: flint
flint:
	docker compose exec frontend yarn lint

# フロントエンドのテストを実行します
.PHONY: ftest
ftest:
	docker compose exec frontend yarn test:run

# フロントエンドの開発サーバーを起動します
.PHONY: finst
finst:
	docker compose exec frontend yarn install

# フロントエンドの開発サーバーを起動します
.PHONY: dev
dev:
	docker compose exec frontend yarn dev

# フロントエンドの開発サーバーを起動します
.PHONY: fdev
fdev:
	docker compose exec frontend yarn debug

# フロントエンドの依存関係を最新化し、ビルドします
.PHONY: front-latest
front-latest:
	docker compose run --rm frontend sh -c "rm -rf node_modules/*"
	docker compose run --rm frontend rm -rf .next
	docker compose run --rm frontend rm -rf .yarn
	docker compose run --rm frontend rm yarn.lock
	docker compose run --rm frontend yarn cache clean
	docker compose run --rm frontend yarn set version latest
	docker compose run --rm frontend yarn add
	docker compose run --rm frontend yarn build

# フロントエンドのキャッシュを削除し、再ビルドします
.PHONY: front-clean
front-clean:
	docker compose run --rm frontend sh -c "rm -rf node_modules/*"
	docker compose run --rm frontend rm -rf .next
	docker compose run --rm frontend rm -rf .yarn
	docker compose run --rm frontend rm yarn.lock
	docker compose run --rm frontend yarn cache clean
	docker compose run --rm frontend yarn add
	docker compose run --rm frontend yarn build

# ==============================================================================
# バックエンド (Backend)
# ==============================================================================

# バックエンドのコンテナ内でbashを起動します
.PHONY: back
back:
	docker compose exec backend bash

# バックエンドのDockerイメージをビルドします
.PHONY: back-build
back-build:
	docker compose build backend

# バックエンドのDockerイメージをキャッシュを使わずにビルドします
.PHONY: build-back
build-back:
	docker compose build --no-cache backend

# Goの依存関係を整理します
.PHONY: tidy
tidy:
	docker compose run --rm backend go mod tidy
	docker compose run --rm backend go mod vendor

# バックエンドのコードをリントチェックします
.PHONY: lint
lint:
	docker compose exec backend golangci-lint run

# バックエンドのテストを実行します
.PHONY: test
test:
	docker compose exec backend go test ./...

# モックファイルを生成します
.PHONY: mock
mock:
	docker compose exec backend go generate ./...

# SwaggerのAPIドキュメントを生成します
.PHONY: swag
swag:
	docker compose exec backend swag init

# デバッグセッション（dlv）終了でダウンしたバックエンドコンテナを再起動します
.PHONY: dlv
dlv:
	docker compose restart backend

# ==============================================================================
# データベース / マイグレーション (Database / Migration)
# ==============================================================================

# データベースのマイグレーションを実行します
.PHONY: migrate
migrate:
	docker compose exec backend bash -c "cd ../database && sql-migrate up"

# データベースのマイグレーションをロールバックします
.PHONY: migrate-down
migrate-down:
	docker compose exec backend bash -c "cd ../database && sql-migrate down"

# テスト用データベースのマイグレーションを実行します
.PHONY: migrate-test
migrate-test:
	docker compose exec backend bash -c "cd ../database && sql-migrate up -env='test'"

# 新しいマイグレーションファイルを作成します (例: make table name=create_users)
.PHONY: table
table:
	docker compose exec backend bash -c "cd ../database && sql-migrate new ${name}"

# ==============================================================================
# スマートコントラクト / ブロックチェーン (Smart Contract / Blockchain)
# ==============================================================================

# ganacheコンテナ内でシェルを起動します
.PHONY: ganache
ganache:
	docker compose exec ganache sh

# Solidityのスマートコントラクトをコンパイルし、Goのバインディングを生成します
.PHONY: sol
sol:
	docker compose run --rm backend rm -rf build
	docker compose run --rm solc --optimize --abi --include-path node_modules/ --base-path . -o /usr/src/contract/build contracts/NFTMarketplace.sol
	docker compose run --rm solc --optimize --bin --include-path node_modules/ --base-path . -o /usr/src/contract/build contracts/NFTMarketplace.sol
	docker compose run --rm backend abigen --abi=./build/NFTMarketplace.abi --bin=./build/NFTMarketplace.bin --pkg=contracts --out=./contracts/NFTMarketplace.go
#   docker compose run --rm solc --optimize --abi --include-path node_modules/ --base-path . -o /usr/src/contract/build contracts/Collection.sol --overwrite
#   docker compose run --rm solc --optimize --bin --include-path node_modules/ --base-path . -o /usr/src/contract/build contracts/Collection.sol --overwrite
#   docker compose run --rm backend abigen --abi=./build/Collection.abi --bin=./build/Collection.bin --pkg=contracts --out=./contracts/Collection.go
