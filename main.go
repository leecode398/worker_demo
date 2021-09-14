package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"worker_demo/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/add_user", AddUser)
	http.ListenAndServe(":8080", nil)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	b, _ := io.ReadAll(r.Body)
	fmt.Println(string(b))
	err := json.Unmarshal(b, &u)
	if err != nil {
		fmt.Println(err)
	}
	u.AddUser()
}
