Run server with

```go run HttpServer.go -port=5555```

to run a server on port 5555

return xml from server with random delay

```curl -v localhost:5555/xml ```

Streess test:
```ab -c200 -n100000 http://127.0.0.1:5555/xml```