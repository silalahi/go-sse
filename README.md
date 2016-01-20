# Golang SSE
This repository is modification from [Kyle's Golang HTML5 SSE Example](https://github.com/kljensen/golang-html5-sse-example). SSE is nearly identical to long polling. The client makes a GET request that establishes a TCP connection. The server keeps this connection open and sends events to the client when they are available.

## Installing
```
go get github.com/silalahi/go-sse
```

## Example of SSE
Check out example folder and try to running server.go
