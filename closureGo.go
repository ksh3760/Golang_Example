package main

import "fmt"

func main() {
	next := nextValue()

	println(next()) // 1
	println(next()) // 2
	println(next()) // 3

	anotherNext := nextValue()
	println(anotherNext()) // 1 다시 시작
	println(anotherNext()) // 2
}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func closureGo() {
	a, b := 3, 5

	f := func(x int) int {
		return a*x + b // 함수 바깥의 변수 a, b 사용
	}

	y := f(5)
	fmt.Println(y) // 20
}
