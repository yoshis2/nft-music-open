# Test Strategy Template (Draft)

openspec の提案において、テストコードの作成を強化するためのガイドラインです。

## 1. Scenario へのテスト方針の記述

各 `#### Scenario:` には、以下の形式でテスト方針を含めることを推奨します。

```markdown
#### Scenario: [シナリオ名]

- **GIVEN** [前提条件]
- **WHEN** [操作]
- **THEN** [期待される結果]
- **TEST_PLAN**:
  - **Method**: [Unit | E2E | Integration]
  - **Focus**: [何を検証するか]
  - **Assertion**: [期待されるアサーションの内容]
  - **Mocking**: [backend: 依存関係のモック（Gateway 等）の必要性] (Backend の場合)
```

## 2. tasks.md へのテストタスクの記述

各実装タスクに対応するテストタスクを含めることを推奨します。

```markdown
- [ ] [実装タスク名]
- [ ] [実装タスク名] に対するテストコードの作成とパスの確認
```
