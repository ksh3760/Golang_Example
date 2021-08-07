package main

import "fmt"

func main() {
	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}

	// for range 문을 사용하여 모든 맵 요소 출력
	// Map은 unordered 이므로 순서는 무작위
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}

// func mapGo() {
// 	var m map[int]string
// 	m = make(map[int]string)
// 	// 추가 혹은 갱신
// 	m[101] = "Apple"
// 	m[102] = "Google"
// 	m[103] = "FaceBook"

// 	//리터럴을 사용한 초기화
// 	tickers := map[string]string{
// 		"GOOG": "Google Inc",
// 		"MSFT": "Microsoft",
// 		"FB":   "FaceBook",
// 	}

// 	// 키에 대한 값 읽기
// 	str := m[102]
// 	println(str)

// 	noData := m[104] // 값이 없으면 nil 혹은 zero 리턴
// 	println(noData)

// 	// 삭제
// 	delete(m, 104)

// 	// Map 키 체크
// 	// val, exists := tickers["MSFT"]
// 	// if !exists {
// 	// 	fmt.Println("No MSFT ticker")
// 	// }

// 	// for loop를 사용한 Map 열거
// 	myMap := map[string]string{
// 		"A": "Apple",
// 		"B": "Banana",
// 		"C": "Charlie",
// 	}

// 	// for range 문을 사용하여 모든 맵 요소 출력
// 	// Map은 underdered 이므로 순서는 무작위
// 	for key, val := range myMap {
// 		fmt.Println(key, val)
// 	}

// }
