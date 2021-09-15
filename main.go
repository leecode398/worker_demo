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
	http.HandleFunc("/update_employee", UpdateEmployee)
	http.HandleFunc("/query_user", QueryUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func parsingJson2Struct(r *http.Request) (*model.User, error) {
	var u model.User
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	u, err := parsingJson2Struct(r)
	if err != nil {
		writeResponse(w, "fail")
		log.Panic(err)
	}
	err = model.AddUser(u)
	if err != nil {
		writeResponse(w, "fail")
		log.Panic(err)
	}
	writeResponse(w, "success")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	u, err := parsingJson2Struct(r)
	if err != nil {
		writeResponse(w, "fail")
		log.Panic(err)
	}
	err = model.DeleteUser(u)
	if err != nil {
		writeResponse(w, "fail")
		log.Panic(err)
	}
	writeResponse(w, "success")
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]interface{})
	b, err := io.ReadAll(r.Body)
	if err != nil {
		writeResponse(w, "fail")
		log.Panic(err)
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		writeResponse(w, "fail")
		log.Panic(err)
	}
	err = model.UpdateUser(m)
	if err != nil {
		writeResponse(w, "fail")
		log.Panic(err)
	}
	writeResponse(w, "success")
}
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]interface{})
	b, err := io.ReadAll(r.Body)
	if err != nil {
		writeResponse(w, "fail")
		log.Panic(err)
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		writeResponse(w, "fail")
		log.Panic(err)
	}
	err = model.UpdateEmployee(m)
	if err != nil {
		writeResponse(w, "fail")
		log.Panic(err)
	}
	writeResponse(w, "success")
}
func QueryUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	b, err := io.ReadAll(r.Body)
	if err != nil {
		writeResponse(w, "fail")
		log.Panic(err)
	}
	err = json.Unmarshal(b, &u)
	if err != nil {
		writeResponse(w, "fail")
		log.Panic(err)
	}
	us, err := model.QueryUser(&u)
	if err != nil {
		writeResponse(w, "fail")
		log.Panic(err)
	}
	writeResponse(w, us)
}

func writeResponse(w http.ResponseWriter, res interface{}) {
	resp := make(map[string]interface{})
	resp["data"] = res
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
