// ======================================
// Golang tcp server example
// ======================================
package main

import (
	"fmt"
	"net"
)

// 포트만 설정하면 모든 NIC의 IP 주소에서 연결을 받는다.
// IP 주소와 포트를 설정하면 특정 NIC에서만 TCP 연결을 받는다.
const (
	G_IP   string = "127.0.0.1"
	G_PORT string = ":8081"
)

func main() {
	var (
		sListener net.Listener
		sErr      error = nil
		conn      net.Conn
	)

	// 포트 번호만 설정하면 모든 네트워크 인터페이스(NIC)의 IP주소에서 연결을 받고
	// 192.168.0.1:8081 처럼 IP주소와 함께 설정하면 특정 NIC에서만 TCP 연결을 받는다.
	// 그리고 TCP 연결 대기 ln은 지연 호출을 사용하여 서버가 끝나면 닫아준다.
	sListener, sErr = net.Listen("tcp", G_IP+G_PORT) // TCP 프로토콜에 8000 포트로 연결을 받음
	if sErr != nil {
		fmt.Println(sErr)
		return
	}

	fmt.Println("TCP port : ", G_IP+G_PORT)
	fmt.Println("Server online")
	fmt.Println("---------------------------------------------------------------")

	defer sListener.Close() // main 함수가 끝나기 직전에 연결 대기를 닫음

	// 무한 루프를 돌면서 클라이언트에서 보낸 데이터를 읽어서 다시 클라이언트로 보냄
	for {
		conn, sErr = sListener.Accept() // 클라이언트가 연결되면 TCP 연결을 반환
		if sErr != nil {
			fmt.Println(sErr)
			continue
		}

		defer conn.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

		go requestHandler(conn) // 패킷을 처리할 함수를 고루틴으로 실행
	}

} // end func main()

func requestHandler(aClient net.Conn) {
	var (
		sData []byte = nil
		sErr  error  = nil
		sNum  int    // 데이터
	)

	// 보통 TCP 서버에서 클라이언트와 패킷(데이터)을 주고받을 때는 패킷의 최대 크기(길이)를 약속하고, 각 패킷의 크기와 구조를 정의하여 사용한다.
	// 특히 MMORPG 같은 온라인 게임에서는 패킷의 구조를 독자적으로 만들어서 사용한다.
	sData = make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성

	for {
		sNum, sErr = aClient.Read(sData) // 클라이언트에서 받은 데이터를 읽음
		if sErr != nil {
			fmt.Println(sErr)
			return
		}

		fmt.Println(string(sData[:sNum])) // 데이터 출력

		_, sErr = aClient.Write(sData[:sNum]) // 클라이언트로 데이터를 보냄
		if sErr != nil {
			fmt.Println(sErr)
			return
		}

	}

} // end func requestHandler()
