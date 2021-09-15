package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	User_name  string `json:"user_name"`
	Address    string `json:"address"`
	Department string `json:"department"`
	Tel        string `json:"tel"`
	User_id    uint32 `json:"user_id"`
	Age        uint8  `json:"age"`
}

type Employee struct {
	Induction_time time.Time `json:"induction_time"`
	Dept_name      string    `json:"dept_name"`
	Position       string    `json:"position"`
}

var db = &sql.DB{}

func init() {
	//连接数据库
	var err error
	db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	err = createTable(db)
	fmt.Println(db)
	if err != nil {
		fmt.Println(err)
	}
}

func createTable(db *sql.DB) error {
	sql := `
	CREATE TABLE users(
	user_id int(11) NOT NULL,
	user_name varchar(20) NOT NULL,
	age int(11) NOT NULL,
	address varchar(30) NOT NULL,
	tel varchar(30) NOT NULL,
	department varchar(30) NOT NULL,
	primary key(user_id)
	);
	`
	_, err := db.Exec(sql)
	return err
}

func AddUser(u *User) error {
	sql := "insert into users (user_id, user_name, age, address, tel, department) values (?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(u.User_id, u.User_name, u.Age, u.Address, u.Tel, u.Department)
	if err != nil {
		log.Printf("insert failed, err:%v\n", err)
		return err
	}
	return nil
}
