package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	User_name  string `json:"user_name"`
	Address    string `json:"address"`
	Department string `json:"department"`
	User_id    uint32 `json:"user_id"`
	Tel        uint32 `json:"tel"`
	Age        uint8  `json:"age"`
}

//type User struct {
//Username string
//Age      int
//Job      string
//Hobby    string
//}

type Employee struct {
	Induction_time time.Time `json:"induction_time"`
	Dept_name      string    `json:"dept_name"`
	Position       string    `json:"position"`
}

// var (
//     createUser     *sql.Stmt
//     createEmployee *sql.Stmt
//     deleteUser     *sql.Stmt
//     deleteEmployee *sql.Stmt
// )

var db *sql.DB

func init() {
	//连接数据库
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		panic(err)
	}
	err = createTable(db)
	if err != nil {
		return
	}
	//insertData(db, User{"zhangsan", 28, "engineer", "play football"})
}

func createTable2(db *sql.DB) error {
	sql := `create table if not exists "users" (
	"id" integer primary key autoincrement,
	"username" text not null,
	"age" integer not null,
	"job" text,
	"hobby" text
	)`
	_, err := db.Exec(sql)
	return err
}

//func insertData(db *sql.DB, u User) error {
//fmt.Println(u)
//sql := `insert into users (username, age, job, hobby) values(?,?,?,?)`
//stmt, err := db.Prepare(sql)
//if err != nil {
//return err
//}
//_, err = stmt.Exec(u.Username, u.Age, u.Job, u.Hobby)
//return err
//}

func createTable(db *sql.DB) error {
	sql := `create table if not exists "users" (
	"user_id" integer primary key not null,
	"user_name" text not null,
	"age" integer not null,
	"address" text not null,
	"tel" integer not null,
	"department" text not null
	)`
	_, err := db.Exec(sql)
	return err
}

func (u User) AddUser() error {
	fmt.Println(u)
	sql := `insert into users (user_id, user_name, age, address, tel, department) values (?, ?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(sql)
	res, err := stmt.Exec(u.User_id, u.User_name, u.Age, u.Address, u.Tel, u.Department)
	fmt.Println(res)
	if err != nil {
		log.Printf("insert failed, err:%v\n", err)
		return err
	}
	return nil
}
