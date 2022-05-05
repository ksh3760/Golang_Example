// ===================================
// RPC Protocol server example
// ===================================
package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Calc int // RPC 서버에 등록하기 위해 임의의 타입으로 정의

// 매개변수
type Args struct {
	A, B int
}

// 기본값
type Reply struct {
	C int
}

func main() {
	var (
		sErr error = nil
		ln   net.Listener
		conn net.Conn
	)

	// rpc 서버에서 함수가 처리될 수 있도록 등록을 해주고, TCP 연결을 받을 준비를 한다.
	rpc.Register(new(Calc))               // Calc 타입의 인스턴스를 생성하여 RPC 서버에 등록
	ln, sErr = net.Listen("tcp", ":8081") // TCP 프로토콜에 6000번 포트로 연결을 받음
	if sErr != nil {
		fmt.Println(sErr)
		return
	}
	defer ln.Close() // main 함수가 종료되기 직전에 연결 대기를 닫음

	for {
		conn, sErr = ln.Accept() // 클라이언트가 연결되면 TCP 연결(커넥션)을 반환
		if sErr != nil {
			continue
		}
		defer conn.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

		go rpc.ServeConn(conn) // RPC를 처리하는 함수를 고루틴으로 실행
	}
}

func (c *Calc) Sum(args Args, reply *Reply) error {
	reply.C = args.A + args.B // 두 값을 더하여 반환값 구조체에 넣어줌
	return nil
}
