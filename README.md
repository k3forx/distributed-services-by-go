# distributed-services-by-go

## Requirements

### Go

```bash
go version
go version go1.19 darwin/amd64
```

### Protocol Buffer

```bash
protoc --version
libprotoc 3.21.5
```

## How to run

```bash
cd cmd/server

go run ./main.go
```

## Protocol Buffers

### 1. 一貫性のあるスキーマ

### 2. バージョン管理

### 3. ボイラーテンプレートコードの削減

protobufライブラリがエンコードとデコードを行ってくれるので、そのコードを手書きする必要がない

### 4. 拡張性

### 5. 言語寛容性


### 6. パフォーマンス

データ量が小さく、JSONに比べて最大6倍の速さでシリアライズできる

## Memo

### Protocol Bufferによる構造化データ

- Goのプロジェクトの規約では、protobufをapiディレクトリに置くことになっている
- protobufでは左にフィールドの型、その名前、フィールドIDを書く
- protobufをコンパイルすると、protobufのメッセージがGoの構造体になり、protobufのバイナリワイヤ形式にマーシャルするための構造体のメソッドとフィールドのゲッターが追加される


### コンパイル
```bash
go get google.golang.org/protobuf/...@v1.28.0
protoc api/v1/*.proto --go_out=. --go_opt=paths=source_relative --proto_path=.
```

## 3章 ログパッケージの作成

ログの概念

- レコード: ログに保存されるデータ
- ストア: レコードを保存するファイル
- インデックス: インデックスエントリを保存するファイル
- セグメント: ストアとインデックスをまとめているもの (抽象的概念)
- ログ: セグメントを全てまとめているもの (抽象的な概念)
