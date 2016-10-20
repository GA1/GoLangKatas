Run first the server:
```go run HttpServer.go -port=5555```

Then, run proxy with:

```go run proxy.go -port=6666```

Call proxy:

```curl -v localhost:6666/json ```