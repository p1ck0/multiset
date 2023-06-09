# Example

Run multiple server for example.

```bash
$ go run main.go -p 8080
$ go run main.go -p 8081
```

Before launch multiset and check terminal.


```bash
$ multiset -f file.json

2023/05/20 15:04:36 | DEBUG | START PARSE
2023/05/20 15:04:36 | DEBUG | PARSE DONE
2023/05/20 15:04:36 | INFO | START SEND MULTI REQ
2023/05/20 15:04:36 | DEBUG | METHOD: POST | URL: http://localhost:8080/api/test | HEADERS: map[Authorization:[Token 124] Content-Type:[application/json]] | BODY: {"enable":true,"name":"test1","type":"auto"}
2023/05/20 15:04:36 | INFO | SENDING REQUEST ON http://localhost:8080/api/test
2023/05/20 15:04:36 | INFO | STATUS CODE FROM http://localhost:8080/api/test: 200
2023/05/20 15:04:36 | INFO | HEADERS FROM http://localhost:8080/api/test: map[Content-Length:[2] Content-Type:[text/plain; charset=utf-8] Date:[Sat, 20 May 2023 13:04:36 GMT]]
2023/05/20 15:04:36 | INFO | BODY FROM http://localhost:8080/api/test: ok
2023/05/20 15:04:36 | DEBUG | METHOD: PUT | URL: http://localhost:8081/api/test | HEADERS: map[Authorization:[Token 123] Content-Type:[application/json]] | BODY: {"enable":false,"name":"test2","type":"auto"}
2023/05/20 15:04:36 | INFO | SENDING REQUEST ON http://localhost:8081/api/test
2023/05/20 15:04:36 | INFO | STATUS CODE FROM http://localhost:8081/api/test: 200
2023/05/20 15:04:36 | INFO | HEADERS FROM http://localhost:8081/api/test: map[Content-Length:[2] Content-Type:[text/plain; charset=utf-8] Date:[Sat, 20 May 2023 13:04:36 GMT]]
2023/05/20 15:04:36 | INFO | BODY FROM http://localhost:8081/api/test: ok
2023/05/20 15:04:36 | INFO | DONE
```

Server terminal.

```bash
$ go run server.go -p 8080

Listening on port 8080
BODY: {"enable":true,"name":"test1","type":"auto"}
HEADERS: map[Accept-Encoding:[gzip] Authorization:[Token 124] Content-Length:[44] Content-Type:[application/json] User-Agent:[Go-http-client/1.1]]
METHOD: POST
```

```bash
$ go run server.go -p 8081

Listening on port 8081
BODY: {"enable":false,"name":"test2","type":"auto"}
HEADERS: map[Accept-Encoding:[gzip] Authorization:[Token 123] Content-Length:[45] Content-Type:[application/json] User-Agent:[Go-http-client/1.1]]
METHOD: PUT
```