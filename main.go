package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	err := json.Unmarshal(b, &u)
	if err != nil {
		fmt.Println(err)
	}
	err = model.AddUser(&u)
	if err != nil {
		writeResponse(w, "fail")
		return
	}
	writeResponse(w, "success")
}

func writeResponse(w http.ResponseWriter, res string) {
	resp := make(map[string]string)
	resp["message"] = res
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Panic(err)
	}
	w.Write(jsonResp)
}
