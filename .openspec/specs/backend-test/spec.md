## ADDED Requirements

### Requirement: Interactor のユニットテスト強化

Interactor のビジネスロジックを検証するためのユニットテストを、Mock を用いて作成しなければならない(SHALL)。

#### Scenario: 正常系：ウォレット情報の取得

- **GIVEN** 有効なウォレットアドレスが存在する
- **WHEN** ウォレット情報の取得を要求する
- **THEN** 対応するウォレットデータが返される
- **TEST_PLAN**:
  - **Method**: Unit Test
  - **Focus**: Interactor から Gateway への呼び出しと戻り値の処理
  - **Assertion**: 取得された ID と Address が Mock で設定した値と一致すること
  - **Mocking**: `MockWalletGateway.EXPECT().Get(...)` を使用して、特定の Address に対して期待する構造体を返すように設定する。

#### Scenario: 異常系：DB エラー時の挙動

- **GIVEN** データベースに一時的な障害が発生している
- **WHEN** ウォレット情報の取得を要求する
- **THEN** 適切なエラーメッセージが返される
- **TEST_PLAN**:
  - **Method**: Unit Test
  - **Focus**: Gateway がエラーを返した際の Interactor のエラーハンドリング
  - **Assertion**: 返されるエラーが期待されるエラー型（例：sql.ErrNoRows ではない予期せぬエラー）であること
  - **Mocking**: `MockWalletGateway.EXPECT().Get(...)` を使用して、エラーを返すように設定する。
