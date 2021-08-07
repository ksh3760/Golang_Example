// ==================
// golang 배열
// ==================
package main

import (
	"fmt"
)

func main() {
	var a [3]int // 정수형 3개 요소를 갖는 배열 a선언
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a[1])

	// 배열 초기화
	var a1 = [3]int{1, 2, 3}
	var a2 = [...]int{1, 2, 3}

	fmt.Println(a1)
	fmt.Println(a2)

	// 다차원 배열
	var multiArr [3][4]int // 정의
	multiArr[0][1] = 10

	fmt.Println(multiArr)

	// 다차원 배열의 초기화
	var muliArrInit = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, // 끝에 콤마 추가
	}
	fmt.Println(muliArrInit)

	arrTest := a1[len(a1)-1:]
	fmt.Println(len(a1))
	fmt.Println("a1 : ", a1)
	fmt.Println("arrTest : ", arrTest)

}
