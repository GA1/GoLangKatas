Run server with

```go run HttpServer.go -port=3333```

to run a server on port 3333

return xml from server with random delay

```curl -v localhost:3333/xml ```

Streess test:
```ab -c200 -n100000 http://127.0.0.1:5555/xml```