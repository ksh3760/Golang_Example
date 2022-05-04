// ======================================
// Golang tcp client example
// ======================================
package main

import (
	"fmt"
	"net"
	"time"
)

const (
	G_IP   string = "127.0.0.1"
	G_PORT string = ":8081"
)

func main() {
	var (
		sClient net.Conn
		sErr    error  = nil
		sData   []byte = nil
		sNum    int
		sStr    string
	)

	fmt.Println("TCP port : ", G_IP+G_PORT)

	sClient, sErr = net.Dial("tcp", G_IP+G_PORT) // TCP 프로토콜
	if sErr != nil {
		fmt.Println(sErr)
		return
	}

	defer sClient.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

	// 서버에서 데이터를 받은 뒤 출력하는 고루틴
	go func(aCon net.Conn) {
		sData = make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성

		for {
			sNum, sErr = aCon.Read(sData) // 서버에서 받은 데이터를 읽음
			if sErr != nil {
				fmt.Println(sErr)
				return
			}

			fmt.Println("aCon : ", aCon)
			fmt.Println(string(sData[:sNum])) // 데이터 출력

			time.Sleep(1 * time.Second)
		}
	}(sClient)

	// 서버에 데이터를 보내는 고루틴
	go func(aCon net.Conn) {
		cnt := 0
		for {
			fmt.Println("input ur msg")
			fmt.Scanln(&sStr)
			// sStr = "Hello" + strconv.Itoa(cnt)

			_, sErr := aCon.Write([]byte(sStr)) // 서버로 데이터를 보냄
			if sErr != nil {
				fmt.Println(sErr)
				return
			}

			cnt++
			time.Sleep(1 * time.Second)
		}
	}(sClient)

	fmt.Scanln()

} // end func main()
