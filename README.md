# NFT Music Marketplace

## 概要

このプロジェクトは「NFT Music Marketplace」という名称で、音楽に特化した NFT の売買ができるマーケットプレイスです。

## ✨ 技術スタック

| カテゴリ             | 技術                                  |
| :------------------- | :------------------------------------ |
| **フロントエンド**   | Next.js, React, TypeScript, ethers.js |
| **バックエンド**     | Go, Echo V4, Solidity, go-ethereum    |
| **インフラ**         | Docker, IPFS, Ganache                 |
| **データベース**     | MySQL, sql-migrate                    |
| **API ドキュメント** | Swagger                               |

## 🚀 開発環境のセットアップ

開発を始めるための手順です。

### 1. 前提条件

- [Docker](https://www.docker.com/get-started/) と `docker-compose`
- `make` コマンド
- GitHub CLI (任意)

### 2. セットアップ手順

1.  **プロジェクトのクローン:**
    リポジトリをローカルにクローンします。

    ```bash
    # GitHub CLI を使う場合
    gh repo clone yoshis2/nft-music

    # または SSH を使う場合
    git clone git@github.com:yoshis2/nft-music.git

    cd nft-music
    ```

2.  **環境変数の設定:**
    サンプルファイルをコピーして、`.env` ファイルを作成・編集します。

    ```bash
    cp .env.sample .env
    ```

    > **Note** > `.env` ファイル内の各変数の詳細については、frontend/README.md を参照してください。

3.  **ビルドと起動:**
    以下のコマンドで、Docker コンテナをビルドして起動します。

    ```bash
    make build && make up
    ```

    > **Note**
    > 初回起動時や依存関係の更新後は `make build` が必要です。それ以降は `make up` だけで起動できます。

4.  **GO フォーマット GO テスト コマンド**

    ```go
    make lint

    make test
    ```

## 画面確認

| カテゴリ             | URL/サービス                             |
| :------------------- | :--------------------------------------- |
| フロントエンド       | http://localhost:3000                    |
| バックエンド Swagger | http://localhost:1323/swagger/index.html |
| IPFS ノード          | http://localhost:5001/webui              |

| **登録が必要なサービス** |                       |
| :----------------------- | :-------------------- |
| メール設定               | SMTP や IMAP 設定     |
| Etherscan                | https://etherscan.io/ |

| **利用しているサービス** |          |
| :----------------------- | :------- |
| IPFS ノード              | ローカル |
| IPFS ノード              | pinata   |

## Swagger の情報

[公式ドキュメント](https://github.com/swaggo/swag#general-api-info)

### Swagger 最新化

```
make swag
```

http://localhost:1323/swagger/index.html

## フロントエンドのデプロイ

```bash
make front-latest
```

## 🗃️ データベース管理

データベースのテーブル定義（スキーマ）の変更や管理については、専用のドキュメントにまとめられています。
詳細は [データベース管理ドキュメント](./document/database.md) を参照してください。

## contract でエラーが発生した時の対処方法

開発中にトランザクションが失敗した場合、以下の手順で解決できることがあります。

### Brave Wallet の Nance をリセット

開発中にトランザクションが失敗し、「Nonce too low」などのエラーが表示されることがあります。これは、Ganache のようなローカルネットワークをリセットした際に、ウォレットが保持している nonce (トランザクションの連番) とブロックチェーン上の nonce が一致しなくなるために発生します。

この問題を解決するには、以下の手順で Brave Wallet のアカウントをリセットしてください。

**手順:**

1.  Brave Wallet を開きます。
2.  「設定」 (歯車アイコン) > 「高度な設定」に移動します。
3.  「アカウントをリセット」をクリックします。

> **Note**
> この操作は、ウォレット内のトランザクション履歴をクリアしますが、資産や秘密鍵には影響しません。主に開発環境での利用を想定しています。

### typechain-types と artifacts と ignition/development を削除して再デプロイ

スマートコントラクトのコードを変更した後などに、フロントエンドとの連携で型定義やデプロイ情報が古くなり、予期せぬエラーが発生することがあります。この場合、古いコンパイル結果や型定義、デプロイ情報を一度クリーンアップして、再デプロイすることで解決できます。

**手順:**

1.  **古い情報の削除:**
    以下のコマンドを実行して、古いコントラクト関連のディレクトリを削除します。
    ```bash
    rm -rf hardhat/artifacts hardhat/cache hardhat/typechain-types hardhat/ignition/deployments
    ```
2.  **再ビルドと再デプロイ:**
    `make` コマンドでコンテナを再ビルドし、起動します。これにより、コントラクトのコンパイルとデプロイが再度実行されます。
    ```bash
    make build && make up
    ```

### Web プロジェクトフォルダに格納

/OneDrive-スリーネクスト/001-Web プロジェクト/007-NFT-MUSIC/101-設定ファイル

### IPFS の画像 URL

```
### ローカル
${IPFS_HOST}${IPFS_GATEWAY_PORT}/ipfs/QmeDmr92XJX1RYBNmBujhWTqhQwuogfo8buv4SL9ARfrPJ
http://localhost:5001/ipfs/QmeDmr92XJX1RYBNmBujhWTqhQwuogfo8buv4SL9ARfrPJ

### グローバル
"https://ipfs.io/ipfs/QmeDmr92XJX1RYBNmBujhWTqhQwuogfo8buv4SL9ARfrPJ?filename=deer.jpg";
```

ブラウザで http://localhost:3000 を開き、結果を確認してください。

## デプロイ完了後の対応

デプロイ後に出てくるデプロイアドレスをコピーし,以下のファイルの LocalMarketAddress 項目のアドレスを上書きする

```ts
/frontend/crs / components / lib / constants.ts;
```

## package.json の解説

詳細は[パッケージ説明](./document/package.md)をご参照ください。

## エラー対応

詳細は[エラー対応](./document/errors.md)をご参照ください。
