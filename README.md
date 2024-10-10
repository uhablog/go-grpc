# GOとgRPCを学ぶためのリポジトリ

[こちら](https://zenn.dev/hsaki/books/golang-grpc-starting/viewer/intro)を参考に進めてみる

## めも

**gPRCサーバーの起動**

/src/cmd/serverディレクトリにて以下コマンドでサーバーの起動ができる

```zsh
go run main.go
```

**gRPCサーバーのメソッド呼び出し**

以下のコマンドでgRPCのコマンドを呼び出すことができる

```zsh
grpcurl -plaintext -d '{"name": "uhablog"}' localhost:8000 myapp.GreetingService.Hello
```