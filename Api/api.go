package api

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

func ApiData() {
	http.HandleFunc("/users", func(wr http.ResponseWriter, r *http.Request) {
		fmt.Println(wr)
		switch r.Method {
		case http.MethodGet: // 조회
			json.NewEncoder(wr).Encode(users) // 인코딩
		case http.MethodPost: // 등록
			var user User
			json.NewDecoder(r.Body).Decode(&user) // 디코딩

			users[user.Email] = &user

			json.NewEncoder(wr).Encode(user) // 인코딩
		}
	})
	http.ListenAndServe(":9999", nil)
}
