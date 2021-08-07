// ================
// 스캐너 예제
// ================
package main

import "fmt"

func main() {
	var numId int
	var name string
	var dpt string

	fmt.Print("insert your number id, name, department.")
	fmt.Scanln(&numId, &name, &dpt)

	fmt.Println("numId :", numId, "/ name :", name, "/ dpt :", dpt)
}
