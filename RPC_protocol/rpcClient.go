// ===================================
// RPC Protocol client example
// ===================================
package main

import (
	"fmt"
	"net/rpc"
)

// 매개변수
type Args struct {
	A, B int
}

// 리턴값
type Reply struct {
	C int
}

const (
	// G_IP   string = "127.0.0.1"
	G_PORT string = ":8081"
)

func main() {
	client, err := rpc.Dial("tcp", G_PORT) // RPC 서버에 연결
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close() // main 함수가 끝나기 직전에 RPC 연결을 닫음

	// 동기 호출
	args := &Args{1, 2}
	reply := new(Reply)
	err = client.Call("calc.Sum", args, reply) // Calc.Sum 함수 호출
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply.C)

	// 비동기 호출
	args.A = 4
	args.B = 9
	sumCall := client.Go("Calc.Sum", args, reply, nil) // Calc.Sum 함수를 고루틴으로 호출
	<-sumCall.Done                                     // 함수가 끝날 때까지 대기
	fmt.Println(reply.C)
}
