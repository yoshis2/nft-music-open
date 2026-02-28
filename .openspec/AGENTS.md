# OpenSpec 利用ガイド

OpenSpec を利用してスペック駆動開発を行う AI コーディングアシスタントのためのガイドです。

## クイックチェックリスト

- 既存の作業を検索: `openspec spec list --long`, `openspec list` (`rg` は全文検索にのみ使用)
- スコープの決定: 新機能 vs 既存機能の修正
- ユニークな `change-id` を選択: ケバブケース、動詞で始める (`add-`, `update-`, `remove-`, `refactor-`)
- 足場作り: `proposal.md`, `tasks.md`, `design.md` (必要な場合のみ)、および影響を受ける機能ごとの差分スペック
- 差分を記述: `## ADDED|MODIFIED|REMOVED|RENAMED Requirements` を使用し、各要件に最低 1 つの `#### Scenario:` を含める
- バリデーション: `openspec validate [change-id] --strict` を実行し、問題を修正する
- 承認依頼: 提案が承認されるまで実装を開始しない

## 言語ルール

- すべての思考、提案、仕様、説明には必ず**日本語**を使用してください。
- 出力ファイル名とコードは英語でも構いませんが、マークダウンファイルの内容は日本語でなければなりません。

## 3 段階のワークフロー

### ステージ 1: 変更の作成

以下の場合に提案を作成します:

- 機能の追加
- 破壊的変更 (API, スキーマ)
- アーキテクチャやパターンの変更
- パフォーマンスの最適化 (動作が変更される場合)
- セキュリティパターンの更新

トリガー (例):

- "変更提案の作成を手伝って"
- "変更の計画を手伝って"
- "提案を作成して"
- "スペック提案を作成したい"
- "スペックを作成したい"

緩やかなマッチングガイダンス:

- `proposal`, `change`, `spec` のいずれかを含む
- `create`, `plan`, `make`, `start`, `help` のいずれかを含む

以下の場合、提案は不要です:

- バグ修正 (意図した動作の復元)
- タイポ、フォーマット、コメントの修正
- 依存関係の更新 (非破壊的)
- 設定の変更
- 既存の振る舞いに対するテスト

**ワークフロー**

1. `.openspec/project.md`, `openspec list`, `openspec list --specs` を確認し、現在のコンテキストを理解します。
2. ユニークで動詞から始まる `change-id` を選び、`.openspec/changes/<id>/` 以下に `proposal.md`, `tasks.md`, オプションの `design.md`、スペック差分を生成します。
3. `## ADDED|MODIFIED|REMOVED Requirements` を用い、各要件に最低 1 つの `#### Scenario:` を含めてスペック差分を記述します。
4. `openspec validate <id> --strict` を実行し、提案を共有する前にすべての問題を解決します。

### ステージ 2: 変更の実装

これらのステップを TODO として追跡し、一つずつ完了させます。

1. **proposal.md を読む** - 何を構築するのかを理解する
2. **design.md を読む** (存在する場合) - 技術的な決定事項を確認する
3. **tasks.md を読む** - 実装のチェックリストを確認する
4. **タスクを順番に実装する** - 順序通りに完了させる
5. **完了の確認** - ステータスを更新する前に `tasks.md` のすべての項目が完了していることを確認する
6. **チェックリストを更新する** - すべての作業が完了したら、各タスクを `- [x]` に設定し、リストが現実を反映するようにする
7. **承認ゲート** - 提案がレビューされ承認されるまで、実装を開始しない

### ステージ 3: 変更のアーカイブ

デプロイ後、別の PR を作成して以下を行います:

- `changes/[name]/` → `changes/archive/YYYY-MM-DD-[name]/` へ移動
- 機能が変更された場合は `specs/` を更新
- ツールのみの変更の場合は `openspec archive <change-id> --skip-specs --yes` を使用 (常に change ID を明示的に渡す)
- `openspec validate --strict` を実行し、アーカイブされた変更がチェックをパスすることを確認する

## すべてのタスクの前に

**コンテキストチェックリスト:**

- [ ] `specs/[capability]/spec.md` の関連スペックを読む
- [ ] `changes/` の保留中の変更で競合がないか確認する
- [ ] `.openspec/project.md` でプロジェクトの規約を読む
- [ ] `openspec list` を実行してアクティブな変更を確認する
- [ ] `openspec list --specs` を実行して既存の機能を確認する

**スペック作成の前に:**

- 機能がすでに存在しないか常に確認する
- 重複するスペックを作成するより、既存のスペックを修正することを優先する
- `openspec show [spec]` を使って現在の状態を確認する
- 要求が曖昧な場合は、足場を作る前に 1〜2 個の明確化のための質問をする

### 検索ガイダンス

- スペックの列挙: `openspec spec list --long` (スクリプトの場合は `--json`)
- 変更の列挙: `openspec list` (または `openspec change list --json` - 非推奨だが利用可能)
- 詳細の表示:
  - スペック: `openspec show <spec-id> --type spec` (フィルタには `--json` を使用)
  - 変更: `openspec show <change-id> --json --deltas-only`
- 全文検索 (ripgrep を使用): `rg -n "Requirement:|"Scenario:" .openspec/specs`

## クイックスタート

### CLI コマンド

```bash
# 基本コマンド
openspec list                  # アクティブな変更を一覧表示
openspec list --specs          # 仕様を一覧表示
openspec show [item]           # 変更または仕様を表示
openspec validate [item]       # 変更または仕様を検証
openspec archive <change-id> [--yes|-y]   # デプロイ後にアーカイブ (非対話的実行には --yes を追加)

# プロジェクト管理
openspec init [path]           # OpenSpec を初期化
openspec update [path]         # 指示ファイルを更新

# 対話モード
openspec show                  # 選択を求めるプロンプト
openspec validate              # 一括検証モード

# デバッグ
openspec show [change] --json --deltas-only
openspec validate [change] --strict
```

### コマンドフラグ

- `--json` - 機械可読な出力
- `--type change|spec` - アイテムを明確化
- `--strict` - 包括的な検証
- `--no-interactive` - プロンプトを無効化
- `--skip-specs` - スペックを更新せずにアーカイブ
- `--yes`/`-y` - 確認プロンプトをスキップ (非対話的アーカイブ)

## ディレクトリ構造

```
.openspec/
├── project.md              # プロジェクト規約
├── specs/                  # 現在の真実 - 何が構築されたか
│   └── [capability]/       # 単一の焦点を絞った機能
│       ├── spec.md         # 要件とシナリオ
│       └── design.md       # 技術的パターン
├── changes/                # 提案 - 何を変更すべきか
│   ├── [change-name]/
│   │   ├── proposal.md     # なぜ、何を、影響
│   │   ├── tasks.md        # 実装チェックリスト
│   │   ├── design.md       # 技術的決定 (任意; 基準を参照)
│   │   └── specs/          # 差分変更
│   │       └── [capability]/
│   │           └── spec.md # ADDED/MODIFIED/REMOVED
│   └── archive/            # 完了した変更
```

## 変更提案の作成

### 決定木

```
新しいリクエスト？
├─ スペックの動作を復元するバグ修正？ → 直接修正
├─ タイポ/フォーマット/コメント？ → 直接修正
├─ 新機能？ → 提案を作成
├─ 破壊的変更？ → 提案を作成
├─ アーキテクチャの変更？ → 提案を作成
└─ 不明瞭？ → 提案を作成 (安全策)
```

### 提案の構造

1. **ディレクトリ作成:** `changes/[change-id]/` (ケバブケース、動詞で始め、ユニークに)

2. **proposal.md の作成:**

```markdown
# 変更: [変更の簡単な説明]

## なぜ

[問題/機会に関する 1-2 文]

## 何が変わるか

- [変更点の箇条書き]
- [破壊的変更は **BREAKING** とマーク]

## 影響

- 影響を受けるスペック: [機能のリスト]
- 影響を受けるコード: [主要なファイル/システム]
```

3. **スペック差分の作成:** `specs/[capability]/spec.md`

```markdown
## ADDED Requirements

### Requirement: 新機能

システムは...を提供しなければならない(SHALL)。

#### Scenario: 成功ケース

- **WHEN** ユーザーがアクションを実行する
- **THEN** 期待される結果

## MODIFIED Requirements

### Requirement: 既存機能

[修正された要件の全文]

## REMOVED Requirements

### Requirement: 古い機能

**理由**: [なぜ削除するのか]
**移行**: [どのように対処するか]
```

複数の機能が影響を受ける場合、`changes/[change-id]/specs/<capability>/spec.md` の下に複数の差分ファイルを作成します — 機能ごとに 1 つ。

4. **tasks.md の作成:**

```markdown
## 1. 実装

- [ ] 1.1 データベーススキーマを作成
- [ ] 1.2 API エンドポイントを実装
- [ ] 1.3 フロントエンドコンポーネントを追加
- [ ] 1.4 テストを記述
```

5. **必要な場合に design.md を作成:**
   以下のいずれかに該当する場合に `design.md` を作成し、それ以外の場合は省略します:

- 横断的な変更 (複数のサービス/モジュール) または新しいアーキテクチャパターン
- 新しい外部依存関係または大幅なデータモデルの変更
- セキュリティ、パフォーマンス、または移行の複雑さ
- コーディング前に技術的な決定から恩恵を受ける曖昧さ

最小限の `design.md` の骨格:

```markdown
## 背景

[背景、制約、関係者]

## ゴール / 非ゴール

- ゴール: [...]
- 非ゴール: [...]

## 決定事項

- 決定: [内容と理由]
- 検討された代替案: [選択肢 + 根拠]

## リスク / トレードオフ

- [リスク] → 緩和策

## 移行計画

[ステップ、ロールバック]

## 未解決の質問

- [...]
```

## スペックファイルのフォーマット

### 重要: シナリオのフォーマット

**正しい** (#### ヘッダーを使用):

```markdown
#### Scenario: ユーザーログイン成功

- **WHEN** 有効な認証情報が提供された
- **THEN** JWT トークンを返す
```

**間違い** (箇条書きや太字は使用しない):

```markdown
- **Scenario: ユーザーログイン** ❌
  **Scenario**: ユーザーログイン ❌

### Scenario: ユーザーログイン ❌
```

すべての要件には、少なくとも 1 つのシナリオが必要です。

### 要件の言葉遣い

- 規範的な要件には SHALL/MUST を使用します (意図的に非規範的でない限り should/may は避ける)

### 差分操作

- `## ADDED Requirements` - 新機能
- `## MODIFIED Requirements` - 変更された振る舞い
- `## REMOVED Requirements` - 非推奨の機能
- `## RENAMED Requirements` - 名前の変更

ヘッダーは `trim(header)` で照合され、空白は無視されます。

#### ADDED と MODIFIED の使い分け

- ADDED: 要件として独立できる新しい機能またはサブ機能を追加する場合。変更が直交的である場合（例：「スラッシュコマンド設定」の追加）、既存の要件の意味を変更するのではなく、ADDED を優先します。
- MODIFIED: 既存の要件の振る舞い、スコープ、または受け入れ基準を変更する場合。常に、更新された要件の全文（ヘッダー + すべてのシナリオ）を貼り付けます。アーカイバーは、提供された内容で要件全体を置き換えます。部分的な差分では、以前の詳細が失われます。
- RENAMED: 名前のみが変更される場合に使用します。振る舞いも変更する場合は、RENAMED（名前）と MODIFIED（内容）を新しい名前を参照して使用します。

よくある落とし穴: 以前のテキストを含めずに、新しい関心事を追加するために MODIFIED を使用すること。これにより、アーカイブ時に詳細が失われます。既存の要件を明示的に変更していない場合は、代わりに ADDED の下に新しい要件を追加します。

MODIFIED 要件を正しく作成する方法:

1. `.openspec/specs/<capability>/spec.md`で既存の要件を見つけます。
2. `### Requirement: ...`からそのシナリオまで、要件ブロック全体をコピーします。
3. `## MODIFIED Requirements`の下に貼り付け、新しい振る舞いを反映するように編集します。
4. ヘッダーテキストが（空白を無視して）完全に一致することを確認し、少なくとも 1 つの`#### Scenario:`を保持します。

RENAMED の例:

```markdown
## RENAMED Requirements

- FROM: `### Requirement: Login`
- TO: `### Requirement: User Authentication`
```

## トラブルシューティング

### 一般的なエラー

**"変更には少なくとも 1 つの差分が必要です"**

- `changes/[name]/specs/` に .md ファイルが存在することを確認
- ファイルに操作プレフィックス (## ADDED Requirements) があることを確認

**"要件には少なくとも 1 つのシナリオが必要です"**

- シナリオが `#### Scenario:` 形式 (ハッシュタグ 4 つ) を使用していることを確認
- シナリオヘッダーに箇条書きや太字を使用しない

**シナリオのサイレント解析失敗**

- 正確な形式が必要: `#### Scenario: Name`
- デバッグ: `openspec show [change] --json --deltas-only`

### 検証のヒント

```bash
# 包括的なチェックには常に strict モードを使用
openspec validate [change] --strict

# 差分解析のデバッグ
openspec show [change] --json | jq '.deltas'

# 特定の要件をチェック
openspec show [spec] --json -r 1
```

## ハッピーパススクリプト

```bash
# 1) 現在の状態を調査
openspec spec list --long
openspec list
# オプションの全文検索:
# rg -n "Requirement:|"Scenario:" .openspec/specs
# rg -n "^#|Requirement:" .openspec/changes

# 2) 変更IDを選択し、足場を作成
CHANGE=add-two-factor-auth
mkdir -p .openspec/changes/$CHANGE/{specs/auth}
printf "## なぜ\n...\n\n## 何が変わるか\n- ...\n\n## 影響\n- ...\n" > .openspec/changes/$CHANGE/proposal.md
printf "## 1. 実装\n- [ ] 1.1 ...\n" > .openspec/changes/$CHANGE/tasks.md

# 3) 差分を追加 (例)
cat > .openspec/changes/$CHANGE/specs/auth/spec.md << 'EOF'
## ADDED Requirements
### Requirement: 二要素認証
ユーザーはログイン時に第二要素を提供しなければならない(MUST)。

#### Scenario: OTPが必要
- **WHEN** 有効な認証情報が提供された
- **THEN** OTPチャレンジが要求される
EOF

# 4) 検証
openspec validate $CHANGE --strict
```

## 複数機能の例

```
.openspec/changes/add-2fa-notify/
├── proposal.md
├── tasks.md
└── specs/
    ├── auth/
    │   └── spec.md   # ADDED: 二要素認証
    └── notifications/
        └── spec.md   # ADDED: OTPメール通知
```

auth/spec.md

```markdown
## ADDED Requirements

### Requirement: 二要素認証

...
```

notifications/spec.md

```markdown
## ADDED Requirements

### Requirement: OTP メール通知

...
```

## ベストプラクティス

### シンプルさ第一

- デフォルトで 100 行未満の新規コード
- 十分であることが証明されるまで単一ファイルでの実装
- 明確な正当化なしにフレームワークを避ける
- 退屈で実績のあるパターンを選択

### 複雑さのトリガー

以下の場合にのみ複雑さを追加します:

- 現在の解決策が遅すぎることを示すパフォーマンスデータ
- 具体的な規模要件 (>1000 ユーザー、>100MB データ)
- 抽象化を必要とする複数の実績のあるユースケース

### 明確な参照

- コードの場所には `file.ts:42` 形式を使用
- `specs/auth/spec.md`としてスペックを参照
- 関連する変更や PR をリンク

### 機能の命名

- 動詞-名詞を使用: `user-auth`, `payment-capture`
- 機能ごとに単一の目的
- 10 分で理解できるルール
- 説明に "AND" が必要な場合は分割

### 変更 ID の命名

- ケバブケースを使用し、短く記述的に: `add-two-factor-auth`
- 動詞で始まるプレフィックスを優先: `add-`, `update-`, `remove-`, `refactor-`
- ユニークさを確保。もし使用済みなら `-2`, `-3` などを追加

## ツール選択ガイド

| タスク                   | ツール | 理由                     |
| ------------------------ | ------ | ------------------------ |
| パターンでファイルを検索 | Glob   | 高速なパターンマッチング |
| コード内容を検索         | Grep   | 最適化された正規表現検索 |
| 特定のファイルを読み込み | Read   | 直接ファイルアクセス     |
| 未知のスコープを調査     | Task   | 複数ステップの調査       |

## エラー回復

### 変更の競合

1. `openspec list` を実行してアクティブな変更を確認
2. 重複するスペックがないかチェック
3. 変更の所有者と調整
4. 提案の結合を検討

### 検証の失敗

1. `--strict` フラグ付きで実行
2. JSON 出力で詳細を確認
3. スペックファイルの形式を検証
4. シナリオが正しくフォーマットされていることを確認

### コンテキストの欠落

1. まず project.md を読む
2. 関連するスペックをチェック
3. 最近のアーカイブを確認
4. 明確化を求める

## クイックリファレンス

### ステージインジケーター

- `changes/` - 提案済み、未構築
- `specs/` - 構築済み、デプロイ済み
- `archive/` - 完了した変更

### ファイルの目的

- `proposal.md` - なぜ、何を
- `tasks.md` - 実装ステップ
- `design.md` - 技術的決定
- `spec.md` - 要件と振る舞い

### CLI の要点

```bash
openspec list              # 何が進行中か？
openspec show [item]       # 詳細を表示
openspec validate --strict # 正しいか？
openspec archive <change-id> [--yes|-y]  # 完了とマーク (自動化には --yes を追加)
```

忘れないでください: スペックは真実です。変更は提案です。同期を保ちましょう。
