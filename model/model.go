package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	User_name  string `json:"user_name"`
	Address    string `json:"address"`
	Department string `json:"department"`
	Tel        string `json:"tel"`
	User_id    uint32 `json:"user_id"`
	Age        uint8  `json:"age"`
	Employee
}

type Employee struct {
	Induction_time string `json:"induction_time"`
	Dept_name      string `json:"dept_name"`
	Position       string `json:"position"`
}

var db = &sql.DB{}

func init() {
	//连接数据库
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/worker_demo?charset=utf8")
	if err != nil {
		panic(err)
	}

	err = createTable(db)
	if err != nil {
		fmt.Println(err)
	}
}

func createTable(db *sql.DB) error {
	users := `
	CREATE TABLE IF NOT EXISTS users(
	user_id int(11) NOT NULL,
	user_name varchar(20) NOT NULL,
	age int(11) NOT NULL,
	address varchar(30) NOT NULL,
	tel varchar(30) NOT NULL,
	department varchar(30) NOT NULL,
	primary key(user_id)
	);`
	employee := `
	CREATE TABLE IF NOT EXISTS employee(
	id int(11) NOT NULL,
	dept_name varchar(30) NOT NULL,
	position varchar(20) NOT NULL,
	induction_time varchar(20) NOT NULL,
	foreign key(id) references users(user_id) on delete cascade on update cascade
	);
	`
	_, err := db.Exec(users)
	if err != nil {
		return err
	}
	_, err = db.Exec(employee)
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
	sql = "insert into employee (id, dept_name, position, induction_time) values (?, ?, ?, ?)"
	stmt, err = db.Prepare(sql)
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(u.User_id, u.Dept_name, u.Position, u.Induction_time)
	if err != nil {
		log.Printf("insert failed, err:%v\n", err)
		return err
	}
	return nil
}

func DeleteUser(u *User) error {
	sql := "delete from users where user_id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(u.User_id)
	if err != nil {
		log.Printf("detele failed, err:%v\n", err)
		return err
	}
	return nil
}

func UpdateUser(u *User) error {
	sql := "delete from users where (user_id = ?)"
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(u.User_id)
	if err != nil {
		log.Printf("detele failed, err:%v\n", err)
		return err
	}
	return nil
}

func QueryUser(u *User) error {
	sql := "select * from users where (user_id = ?)"
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Panic(err)
	}
	ans, err := stmt.Exec(u.User_id)
	fmt.Println(ans)
	if err != nil {
		log.Printf("detele failed, err:%v\n", err)
		return err
	}
	return nil
}
