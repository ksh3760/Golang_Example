module goChatServer

replace chat => ./chat

go 1.16

require (
	chat v0.0.0-00010101000000-000000000000
	github.com/gomodule/redigo v1.8.8 // indirect
	github.com/googollee/go-socket.io v1.6.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
)
