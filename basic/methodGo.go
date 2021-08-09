package main

// Rect - struct 정의
type Rect struct {
	width, height int
}

// Rect의 area() 메소드
func (r Rect) area() int {
	return r.width * r.height
}

func main() {
	recta := Rect{10, 20}
	areana := recta.area() // 메소드 호출
	println(areana)
}

// 포인터 Receiver
// func (r *Rect) area2() int {
// 	r.width++
// 	return r.width * r.height
// }

// func main() {
// 	rect := Rect{10, 20}
// 	area := rect.area2()      // 메소드 호출
// 	println(rect.width, area) // 11 220 출력
// }
