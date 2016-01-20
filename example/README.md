# Go SSE Example

For example SSE program, you can execute
```
go run server.go
```
Then open http://localhost:8000 from your browser, you should get output similar like this:
```
Server Sent Event (EventSource) example:
1 - the time is 2016-01-20 11:25:48.422806026 +0700 WIB
2 - the time is 2016-01-20 11:25:53.424385546 +0700 WIB
3 - the time is 2016-01-20 11:25:58.42790376 +0700 WIB
4 - the time is 2016-01-20 11:26:03.432141024 +0700 WIB
```
And, you should see output in the terminal like the following
```
2016/01/20 11:25:33 Sent message 0
2016/01/20 11:25:33 Broadcase message to 0 clients
2016/01/20 11:25:33 New client added. 1 registered clients
2016/01/20 11:25:38 Sent message 1
2016/01/20 11:25:38 Broadcase message to 1 clients
2016/01/20 11:25:43 Sent message 2
2016/01/20 11:25:43 Broadcase message to 1 clients
2016/01/20 11:25:43 Finished HTTP request at  /
2016/01/20 11:25:43 HTTP connection just closed
2016/01/20 11:25:43 Removed client. 0 registered clients
2016/01/20 11:25:43 Finished HTTP request at  /event/
2016/01/20 11:25:43 New client added. 1 registered clients
2016/01/20 11:25:43 Finished HTTP request at  /favicon.ico
2016/01/20 11:25:48 Sent message 3
2016/01/20 11:25:48 Broadcase message to 1 clients
2016/01/20 11:25:53 Sent message 4
2016/01/20 11:25:53 Broadcase message to 1 clients
2016/01/20 11:25:58 Sent message 5
2016/01/20 11:25:58 Broadcase message to 1 clients
2016/01/20 11:26:03 Sent message 6
2016/01/20 11:26:03 Broadcase message to 1 clients
```
