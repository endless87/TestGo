package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* var users = map[string]*User{} */

var data = map[string]*Info{}

/* type User struct {
	User string `json:"user"`
} */

type Info struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

/* var datas = map[string]*Data{}

type Data struct {
	Data string `json:"data"`
} */

func main() {

	fmt.Println("test api start")

	http.HandleFunc("/user", func(wr http.ResponseWriter, r *http.Request) {
		fmt.Println(wr)
		fmt.Println(r)
		switch r.Method {
		case http.MethodGet: // 조회
			fmt.Println("---Get")
			json.NewEncoder(wr).Encode(data) // 인코딩
		case http.MethodPost: // 등록
			fmt.Println("---Post")
			var info Info
			json.NewDecoder(r.Body).Decode(&info) // 디코딩

			data["user"] = &info

			json.NewEncoder(wr).Encode(data) // 인코딩
		}
	})

	/* 	http.HandleFunc("/data", func(wr http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(wr).Encode(datas)
		case http.MethodPost:
			var data Data

			fmt.Println(data)

			json.NewDecoder(r.Body).Decode(&data)

			fmt.Println(&data)

			datas[data.Data1] = &data
			datas[data.Data2] = &data

			json.NewEncoder(wr).Encode(datas)
		}
	}) */

	http.ListenAndServe(":9999", nil)

}

// func testInfo() {

// 	name := "Soo Kwon"
// 	message := "Test Go Go"
// 	fmt.Println(name)
// 	fmt.Println(message)

// 	data1 := 1
// 	data2 := 2

// 	fmt.Println(data1 + data2)

// 	Data.GetMessage()

// 	totalLength, upperName := Data.LenAndUpper(name)
// 	fmt.Println(totalLength, upperName)

// 	Data.RepeatMe("one", "two", "three", "four", "five")

// 	totalLengthP, upperNameP := Data.LenAndUpperNakedReturn(name + " " + name)
// 	fmt.Println(totalLengthP, upperNameP)

// 	Data.SuperAdd(10, 20, 30, 40, 50)

// }
