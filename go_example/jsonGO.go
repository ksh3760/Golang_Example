package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Member -
type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {

	// Go data
	mem := Member{"Shawn", 10, true}

	// JSON 인코딩
	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}

	// JSON 바이트를 문자열로 변경
	jsonString := string(jsonBytes)

	fmt.Println(jsonString)

	encoFunc()

}

func encoFunc() {
	// 테스트용 JSON 데이타
	jsonBytes, _ := json.Marshal(Member{"Tim", 1, true})

	// JSON 디코딩
	var mem Member
	err := json.Unmarshal(jsonBytes, &mem)
	if err != nil {
		panic(err)
	}

	// mem 구조체 필드 엑세스
	fmt.Println(mem.Name, mem.Age, mem.Active)
}

// json 파일
// {
//     "id":"admin",
//     "pw":"12345",
//     "hello":"world",
//     "go":"lang"
// }

// ================================
// json test
// ================================
func JsonTestFunc() {
	type Info struct {
		Id string `json:"id"`
		Pw string `json:"pw"`
	}

	type Hello struct {
		Hello string `json:"hello"`
		Go    string `json:"go"`
	}

	data, err := os.Open("logJson/jsonSetting.conf")
	if err != nil {
		fmt.Printf("ERROR : %v", err)
	}
	defer data.Close()

	byteValue, _ := ioutil.ReadAll(data)
	fmt.Println("byteValue : ", string(byteValue)) // 확인용

	var info Info
	json.Unmarshal(byteValue, &info)
	fmt.Println("Info : ", info)

	// byteValue1, _ := ioutil.ReadAll(data)

	var hello Hello
	json.Unmarshal(byteValue, &hello)
	fmt.Println("Hello : ", hello)

}
