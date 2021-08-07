package main

import "fmt"

func main() {
	// person 객체 생성
	p := person{}

	// 필드값 설정
	p.name = "Kim"
	p.age = 10

	fmt.Println(p)

	rect := Rect{10, 20}
	area := rect.area() // 메소드 호출
	fmt.Println(area)

	// rect := Rect{10, 20}
	// area := rect.area2() // 메소드 호출
	// fmt.Println(rect.width, area)

}

type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

// 생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}

	return &d
}

// Rect - struct 정의
type Rect struct {
	width, height int
}

// Rect의 area() 메소드
func (r Rect) area() int {
	return r.width * r.height
}

// 포인터 Receiver
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}
