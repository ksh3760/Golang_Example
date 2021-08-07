package main

import "time"

func main() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

	EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")

		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}
}


func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

func ChannelGo() {
	// 정수형 채널을 생성한다.
	ch := make(chan int)

	go func() {
		ch <- 123 // 채널에 123을 보낸다.
	}()

	var i int
	i = <-ch // 채널에 123을 보낸다.
	println(i)
}
