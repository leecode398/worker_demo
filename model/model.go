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
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/worker_demo?charset=utf8")
	// db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/worker_demo?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	err = createTable(db)
	fmt.Println(db)
	if err != nil {
		fmt.Println(err)
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
	sql := `
	CREATE TABLE users(
	user_id int(11) NOT NULL,
	user_name varchar(20) NOT NULL,
	age int(11) NOT NULL,
	address varchar(30) NOT NULL,
	tel int(15) NOT NULL,
	department varchar(30) NOT NULL,
	primary key(user_id)
	);
	`
	_, err := db.Exec(sql)
	return err
}

func (u User) AddUser() error {
	// fmt.Println(u)
	// sql := "insert into users (user_id, user_name, age, address, tel, department) values (?, ?, ?, ?, ?, ?)"
	// stmt, err := db.Prepare(sql)
	// defer stmt.Close()
	// if err != nil {
	//     log.Panic(err)
	// }
	// res, err := stmt.Exec(u.User_id, u.User_name, u.Age, u.Address, u.Tel, u.Department)
	// fmt.Println(res)
	// if err != nil {
	//     log.Printf("insert failed, err:%v\n", err)
	//     fmt.Println("err1")
	//     return err
	// }
	// fmt.Println("err2")
	// return nil
	fmt.Println("AddUser")
	fmt.Println(db)
	stmt, err := db.Prepare("insert into t(name,ts) values(?,?)")
	fmt.Println("AddUser2")
	defer stmt.Close()
	if err != nil {
		log.Println(err)
	}
	ts, _ := time.Parse("2006-01-02 15:04:05", "2014-08-28 15:04:00")
	stmt.Exec("edmond", ts)
	return nil
}
