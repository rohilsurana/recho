# recho

HTTP Server to respond back with a dump of request

## Usage

### Start the server
```bash
$ go run main.go
```

### Call the server
```bash
$ curl http://localhost:8080/any/path
```
Output
```
GET /any/path HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/7.64.1

```
