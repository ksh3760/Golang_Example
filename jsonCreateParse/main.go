// struct에 값을 입력하여 JSON파일 생성
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Author struct {
	Name  string `json:"name"` 
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"`
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"`
	Content    string    `json:"content"`
	Recommends []string  `json:"recommends"`
	Comments   []Comment `json:"comments"`
}

func main() {
	data := make([]Article, 1) // 값을 저장할 구조체 슬라이스 생성

	data[0].Id = 1
	data[0].Title = "Hello, world!"
	data[0].Author.Name = "JohnDoe"
	data[0].Author.Email = "JohnDoe@gmail.com"
	data[0].Content = "content"
	data[0].Recommends = []string{"Kim", "Park"}
	data[0].Comments = make([]Comment, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Kim"
	data[0].Comments[0].Author.Email = "kim@gmail.com"
	data[0].Comments[0].Content = "hello"

	doc, _ := json.Marshal(data) // data를 JSON 문서로 변환

	err := ioutil.WriteFile("./bbs.json", doc, os.FileMode(0644)) // bbs.json 파일에 JSON 문서 저장
	if err != nil {
		fmt.Println(err)
		return
	}
}