# Project Context: NFT Music Marketplace

## Purpose

このプロジェクトの目的は、音楽を NFT（Non-Fungible Token）として発行し、売買するためのマーケットプレイスを構築することです。アーティストは自身の楽曲をアップロードして NFT としてミント（発行）でき、ユーザーはウォレットを接続してそれらを購入できます。楽曲のメタデータや音源ファイルは IPFS を利用して分散型で管理されます。

## Tech Stack

### Backend

- **言語:** Go
- **アーキテクチャ:** Clean Architecture
- **データベース:** MySQL
- **DB マイグレーション:** [sql-migrate](https://github.com/rubenv/sql-migrate)
- **コンテナ:** Docker, Docker Compose
- **ビルドツール:** Make
- **主要な責務:**
  - API エンドポイントの提供
  - データベースとの連携
  - EVM 互換ブロックチェーンとの対話
  - IPFS へのファイルアップロードと管理

### Frontend

- **フレームワーク/言語:** Next.js, React, TypeScript
- **パッケージマネージャ:** Yarn
- **スタイリング:** Tailwind CSS
- **フォーム/バリデーション:** React Hook Form, Zod
- **コード品質:** ESLint, Prettier
- **主要な責務:**
  - ユーザーインターフェースの提供
  - ウォレット（例: MetaMask）との連携
  - スマートコントラクトとの対話 (Ethers.js 等を利用)
  - Backend API との通信

### Blockchain

- **言語:** Solidity
- **プラットフォーム:** EVM (Ethereum Virtual Machine) 互換のブロックチェーン
- **スマートコントラクト:** NFT のミント、所有権、取引ロジックを実装

## Project Conventions

### Code Style

- **Frontend:** Prettier, ESLint の規約に従います。リポジトリ内の `frontend/.prettierrc`, `frontend/.eslintrc.json` を参照してください。
- **Backend:** Go の標準的なフォーマット（`gofmt`）に従います。

### Language

- Documentation & Specs: **Japanese (日本語)**
- Code Comments: **Japanese (日本語)**

### Architecture Patterns

- **Backend:** `domain`, `usecases`, `adapters`, `infrastructure` のディレクトリ構成に基づいた Clean Architecture を採用しています。これにより、ビジネスロジックと技術的実装の分離を図ります。
- **Frontend:** Next.js の App Router をベースとしたディレクトリ構成を採用しています。機能やページごとにコンポーネントを整理し、再利用性を高めます。

### Testing Strategy

- **Backend:** Go の標準テスト機能を用いた単体テストを必須とします (`*_test.go`)。詳細は `.openspec/specs/backend-test/spec.md` を参照してください。
- **Frontend:** 単体テストおよび統合テストを必須とします。詳細は `.openspec/specs/frontend/spec.md` を参照してください。E2E テストについては、必要に応じて Cypress や Playwright などの導入を検討します。

### Git Workflow

[Git のブランチ戦略（例: GitHub Flow）や、コミットメッセージの規約（例: Conventional Commits）について記述してください]

## Domain Context

- **NFT:** Non-Fungible Token の略。ブロックチェーン上で所有権が記録される、唯一無二のデジタル資産。
- **ミント (Mint):** デジタルデータ（このプロジェクトでは楽曲）をブロックチェーン上に NFT として新規に発行するプロセス。
- **IPFS:** InterPlanetary File System の略。P2P ネットワークを利用した分散型ファイルシステムで、耐障害性と耐検閲性が高い。楽曲ファイルや NFT メタデータの保存に使用します。
- **EVM:** Ethereum Virtual Machine の略。スマートコントラクトを実行するための環境。

## Important Constraints

[技術的、ビジネス的、あるいは規制上の制約事項（予算、パフォーマンス要件、セキュリティ基準、特定のブロックチェーンネットワークの指定など）をリストアップしてください]

## External Dependencies

[利用している主要な外部サービス、サードパーティ API（例: Infura, Alchemy）、連携システムについて記述してください]
