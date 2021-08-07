package main

func main() {
	s := make([]int, 5, 10)
	println(len(s), cap(s)) // len 5, cap 10
}
