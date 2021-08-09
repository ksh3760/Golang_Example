// ====================
// 실행 시간 측정
// ====================
package main

import (
	"fmt"
	"time"
)

func main() {
	// 시작 시간
	startTime := time.Now()

	// Task 실행
	for i := 0; i < 100; i++ {
		println(i)
	}

	// 경과 시간
	elapsedTime := time.Since(startTime)

	fmt.Printf("실행시간: %s\n", elapsedTime)
}
