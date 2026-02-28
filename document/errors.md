# 注意事項

docker-compose.yml と env ファイルが同じ階層にないと env_file は読み込まない

## 以下のエラーが発生した場合

フロントエンドで以下のようなエラーがみられた場合
Error: Cannot find module 'next/dist/server/app-render/work-async-storage.external.js'

原因は node_modules の中身が nextjs のバージョンと合っていないため
対策としてはビルドし直す必要がある

以下のコマンドを実行

```
make front-latest

make down

make up
```

solidity のデプロイ方法はこちらを参照

https://crypto-currency-academy.com/ethereum-smart-contract/
