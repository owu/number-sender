# number sender

Number-sender は事前定義されたルールに基づいて自動インクリメント数値を分類管理するツールです（例：ユーザー登録時のユーザーID発行、配信者へのライブルームID割り当て、チャットグループへのグループID付与など）。実際のビジネスシナリオでは、タイプを指定して番号を取得可能です（例：一般ユーザーには標準番号、有料ユーザーには特別な番号を割り当てる）。

---

[English](./README.md) | [中文](./README_cn.md) | 日本語

## ビルド
```
go mod vendor && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o owu-number-sender-go main.go

```

## 実行
```
./owu-number-sender-go --config=config/config-dev.toml

```


## apiインターフェース
#### サービスステータスチェック
- リクエスト例:
```
curl \
-H "Token:f98f84f7939a56f6ec8c42ef088139f5" \
-H "Milli:1746460799000" \
'http://localhost:8080/ping'
```

- レスポンス例:
```
{"data":{"time":1747067149074},"error":0,"msg":"success"}
```
---

#### キャッシュ内の各プラン数量照会
- リクエスト例:
```
curl \
-H "Token:f98f84f7939a56f6ec8c42ef088139f5" \
-H "Milli:1746460799000" \
'http://localhost:8080/api/len'

```
- レスポンス例:
```
{"data":{"starter":3472008,"standard":2237018,"premium":82317,"ultimate":5657},"error":0,"msg":"success"}
```

#### 指定プランの数値取得
- リクエスト例:

```
# {:plan} は starter, standard, premium, ultimate のいずれか

curl \
-H "Token:f98f84f7939a56f6ec8c42ef088139f5" \
-H "Milli:1746460799000" \
'http://localhost:8080/api/pop/{:plan}'
```
- レスポンス例:
```
{"data":{"starter":0,"standard":0,"premium":0,"ultimate":10101},"error":0,"msg":"success"}
```

#### API認証プロトコル
```
Header: Milli , ミリ秒単位のタイムスタンプ
Header: Token , md5({Milli},{Encrypt}) で生成

---

Token生成例:
{Milli} 現在のミリ秒タイムスタンプ（例: 1746460799000）
{Encrypt} config/config-*.toml ファイルに設定されるapi.encrypt値（例: 9b0b7484bc65e241804ce8eeb014f247）

Token = MD5(1746460799000,9b0b7484bc65e241804ce8eeb014f247) = f98f84f7939a56f6ec8c42ef088139f5
```

