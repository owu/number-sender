# number sender

Number-senderは、事前設定されたルールに基づいてインクリメンタルな数字を分類管理するツールで、実体に一意の数字を割り当てる必要があるシナリオ（例：ユーザー登録時のユーザーID割り当て、配信者へのライブルームID割り当て、チャットグループへのグループ番号割り当てなど）に設計されています。実際の業務シナリオでは、開発者はタイプを指定して番号を取得できます（例：一般ユーザーに通常の番号を割り当て、有料ユーザーに専用の特別な番号を割り当てるなど）。

---

[English](./README.md) | [中文](./README_cn.md) | 日本語

## 1. コンパイル
```
go mod vendor && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o number-sender-go main.go

```

## 2. 実行
```
./number-sender-go --config=config/config-test.toml

設定ファイル内の app.env 変数は、開発環境では test 、本番環境では prod です。
インターフェース認証ミドルウェア /internal/pkg/mware/auth.go では、test 環境では認証を行わないように判定されており、テストが容易になっています。
```


## 3. APIインターフェース
### 3.1 サービスステータスチェック
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

### 3.2 キャッシュ内の各プランの数量を問い合わせる
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

### 3.3 指定されたプランの番号を取得する
- リクエスト例:

```
# {:plan} は starter, standard, premium, ultimate に置き換えることができます。

curl \
-H "Token:f98f84f7939a56f6ec8c42ef088139f5" \
-H "Milli:1746460799000" \
'http://localhost:8080/api/pop/{:plan}'
```
- レスポンス例:
```
{"data":{"starter":0,"standard":0,"premium":0,"ultimate":10101},"error":0,"msg":"success"}
```

### 3.4 API認証プロトコル
```
リクエストヘッダ: Milli , タイムスタンプ（ミリ秒形式）
リクエストヘッダ: Token , md5({Milli},{Encrypt}) によって生成

---

Token生成例:
{Milli} は現在のミリ秒タイムスタンプで、1746460799000 と仮定します
{Encrypt} は config/config-*.toml ファイルで設定され、api.encrypt の値が 9b0b7484bc65e241804ce8eeb014f247 と仮定します

Token = MD5(1746460799000,9b0b7484bc65e241804ce8eeb014f247) = f98f84f7939a56f6ec8c42ef088139f5
```


## 4. Docker デプロイ

### 4.0 プロジェクト準備

Dockerデプロイを行う前に、プロジェクトコードをクローンして指定のブランチに切り替える必要があります：

```bash
git clone https://github.com/owu/number-sender.git && cd number-sender && git checkout main && git pull && chmod 777 -R ./docker && chmod +x docker-tools.sh
```

### 4.1 Docker ツールスクリプト

プロジェクトでは、Docker イメージのビルド、コンテナ管理、イメージ移行などの操作を簡素化するための統合 Docker 管理スクリプト `docker-tools.sh` を提供しています。

#### 機能特徴
- デリバリーイメージのビルド、コンテナの起動停止、ログ表示、イメージ移行機能を統合
- インタラクティブな選択とコマンドラインパラメータの両方の使用方法をサポート
- 各操作で実際に実行される Docker コマンドを表示し、透明度を高める
- モジュラー設計により、後の保守と拡張が容易

#### 使用方法

##### インタラクティブモード
```bash
./docker-tools.sh
```

実行後、以下のメニューが表示されます。対応する文字を入力して操作を選択します：
```
========================================
Docker ツールスクリプト
========================================
実行する操作を選択してください:
b) Dockerイメージを構築 (docker build)
d) Dockerコンテナを停止 (docker compose down)
l) コンテナログを表示 (docker logs)
m) イメージ移行 (docker save/load)
u) Dockerコンテナを起動 (docker compose up -d)
直接Enterキーを押してスクリプトを終了
```

##### コマンドラインパラメータモード
```bash
./docker-tools.sh [オプション]
```

サポートされているオプション:
- `b` : Dockerイメージを構築
- `d` : Dockerコンテナを停止
- `l` : コンテナログを表示
- `m` : イメージ移行機能を実行（サブメニューに入る）
- `u` : Dockerコンテナを起動
- `h` : ヘルプ情報を表示

#### 各機能の説明

1. **Dockerイメージを構築 (b)**
   - イメージが既に存在するかどうかを確認し、重複したビルドを回避
   - イメージが実行中かどうかを確認し、ビルドの安全性を確保
   - docker-compose.yml のイメージバージョンを自動的に更新

2. **Dockerコンテナを起動 (u)**
   - `docker compose up -d` を使用してコンテナを起動
   - コンテナの起動状態を表示

3. **Dockerコンテナを停止 (d)**
   - `docker compose down` を使用してコンテナを停止
   - 現在実行中のコンテナのリストを表示

4. **コンテナログを表示 (l)**
   - `docker logs number-sender` を使用してコンテナログを表示

5. **イメージ移行 (m)**
   - イメージをローカルファイルにエクスポート（`docker save` 使用）
   - ローカルファイルからイメージをインポート（`docker load` 使用）
   - エクスポートまたはインポート操作のインタラクティブな選択をサポート

#### イメージ移行サブメニュー
`m` オプションを選択すると、イメージ移行サブメニューに入ります：
```
========================================
Dockerイメージ移行機能
現在の設定:
  イメージ名: owu/number-sender
  イメージバージョン: 0.0.1
  完全なイメージ名: owu/number-sender:0.0.1
  エクスポートファイル名: owu.number-sender.v0.0.1.tar
========================================
実行する操作を選択してください:
1) イメージをエクスポート (docker save)
2) イメージをインポート (docker load)
直接Enterキーを押してイメージ移行機能を終了
```

### 4.2 従来のDockerデプロイ方法（オプション）

従来の方法でデプロイする必要がある場合は、以下のコマンドを手動で実行することもできます：

1. イメージを構築：
```bash
docker build -t owu/number-sender:0.0.1 .
```

2. コンテナを起動：
```bash
docker compose up -d
```

3. コンテナを停止：
```bash
docker compose down
```

4. ログを表示：
```bash
docker logs number-sender
```

5. イメージ移行：
```bash
# イメージをエクスポート
docker save -o owu.number-sender.v0.0.1.tar owu/number-sender:0.0.1

# イメージをインポート
docker load -i owu.number-sender.v0.0.1.tar
```