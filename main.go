package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

}

func addUser(w http.ResponseWriter, r *http.ReadRequest) {
	u := model.Users{}
}
