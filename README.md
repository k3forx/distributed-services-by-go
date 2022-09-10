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
