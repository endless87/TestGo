package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var users = map[string]*User{}

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

var datas = map[string]*Data{}

type Data struct {
	Data1 string `json:"data1"`
	Data2 string `json:"data2"`
}

func main() {

	fmt.Println("test api")

	http.HandleFunc("/users", func(wr http.ResponseWriter, r *http.Request) {
		fmt.Println(wr)
		fmt.Println(r)
		switch r.Method {
		case http.MethodGet: // 조회
			fmt.Println("---Get")
			json.NewEncoder(wr).Encode(users) // 인코딩
		case http.MethodPost: // 등록
			fmt.Println("---Post")
			var user User

			fmt.Println(user)

			json.NewDecoder(r.Body).Decode(&user) // 디코딩

			users[user.Email] = &user

			json.NewEncoder(wr).Encode(users) // 인코딩
		}
	})

	http.HandleFunc("/data", func(wr http.ResponseWriter, r *http.Request) {
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
	})

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
