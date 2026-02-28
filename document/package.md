# パッケージ説明

このドキュメントは `package.json` に記載されている主要なパッケージについての簡単な説明を提供します。

## `package.json` スクリプト解説

- `yarn dev`: Turbopack を使用して開発サーバーを起動します。変更はリアルタイムでブラウザに反映されます。
- `yarn build`: 本番環境向けのアプリケーションをビルドします。
- `yarn start`: ビルドされた本番アプリケーションを起動します。
- `yarn lint`: ESLint を使用してコードの静的解析とフォーマットチェックを行います。

## Dependencies (アプリケーションの依存関係)

- **`ethers`**: Ethereum ブロックチェーンとの対話を容易にするための、完全でコンパクトなライブラリです。ウォレット管理、スマートコントラクトの実行、トランザクションの署名などが可能です。
- **`next`**: React ベースのフレームワークで、サーバーサイドレンダリング（SSR）、静的サイト生成（SSG）、ルーティング、API ルートなどの機能を提供します。
- **`nodemailer`**: Node.js 環境（主に Next.js の API ルート）からメールを送信するためのモジュールです。
- **`react`**, **`react-dom`**: ユーザーインターフェースを構築するための JavaScript ライブラリと、それをブラウザ DOM に描画するためのパッケージです。
- **`react-hook-form`**: React でフォームを扱うためのパフォーマンスが高く、柔軟で拡張可能なライブラリです。バリデーションや状態管理を簡素化します。
- **`zod`**: TypeScript ファーストのスキーマ宣言および検証ライブラリです。静的型推論を活用して、データの型安全性を保証します。

## DevDependencies (開発用の依存関係)

- **`@hookform/resolvers`**: `react-hook-form` と `zod` のようなバリデーションライブラリを連携させるためのアダプターです。
- **`@typescript-eslint/eslint-plugin`**: TypeScript 固有の構文や型情報を利用した ESLint ルールを提供します。
- **`@typescript-eslint/parser`**: ESLint が TypeScript コードを解析できるようにするためのパーサーです。`@typescript-eslint/eslint-plugin` と共に使用されます。
- **`@types/*`**: (`@types/node`, `@types/react`など) 各種ライブラリの TypeScript 型定義ファイルです。型チェックとエディタのコード補完を可能にします。
- **`eslint`**: JavaScript および JSX のコードを静的解析し、問題点を報告したり自動修正したりするリンターツールです。
- **`eslint-config-next`**: Next.js プロジェクト向けの ESLint 推奨設定を提供します。Next.js 特有のルールやベストプラクティスが含まれています。
- **`eslint-config-prettier`**: ESLint と Prettier の競合するルールを無効化するための設定です。Prettier にフォーマットを任せ、ESLint はコード品質のチェックに専念させます。
- **`eslint-plugin-prettier`**: ESLint のルールとして Prettier を実行し、フォーマットの差異を ESLint のエラーとして報告します。
- **`prettier`**: コードフォーマッターです。コードスタイルを統一し、可読性を向上させます。
- **`tailwindcss`**: ユーティリティファーストの CSS フレームワークです。クラス名を組み合わせることで迅速に UI を構築できます。
- **`typescript`**: JavaScript に静的型付けを追加したスーパーセット言語です。大規模開発でのコードの堅牢性や保守性を高めます。

## Package Manager

- **`yarn@4.9.2`**: このプロジェクトで使用されているパッケージマネージャーのバージョンです。

---

_この説明は一般的なものであり、プロジェクト内での具体的な使用方法や目的を完全に網羅するものではありません。_
_各パッケージの詳細については、公式ドキュメントを参照してください。_
