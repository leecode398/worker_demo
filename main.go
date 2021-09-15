package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"worker_demo/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/add_user", AddUser)
	http.HandleFunc("/delete_user", DeleteUser)
	http.HandleFunc("/update_user", UpdateUser)
	http.HandleFunc("/query_user", QueryUser)
	http.ListenAndServe(":8080", nil)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	b, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(b, &u)
	if err != nil {
		log.Println(err)
		return
	}
	err = model.AddUser(&u)
	if err != nil {
		writeResponse(w, "fail")
		return
	}
	writeResponse(w, "success")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	b, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(b, &u)
	if err != nil {
		log.Println(err)
		return
	}
	err = model.DeleteUser(&u)
	if err != nil {
		writeResponse(w, "fail")
		return
	}
	writeResponse(w, "success")
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	b, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(b, &u)
	if err != nil {
		log.Println(err)
		return
	}
	err = model.UpdateUser(&u)
	if err != nil {
		writeResponse(w, "fail")
		return
	}
	writeResponse(w, "success")
}
func QueryUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	b, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(b, &u)
	if err != nil {
		log.Println(err)
		return
	}
	err = model.QueryUser(&u)
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
