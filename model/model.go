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
	Employee
}

type Employee struct {
	Induction_time string `json:"induction_time"`
	Dept_name      string `json:"dept_name"`
	Position       string `json:"position"`
	ID             uint32 `json:"id"`
}

var db = &sql.DB{}

func init() {
	//连接数据库
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/worker_demo?charset=utf8")
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(8 * time.Hour)
	db.SetConnMaxIdleTime(8 * time.Hour)
	db.Ping()
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
	user_id int(11) NOT NULL COMMENT,
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

func UpdateUser(m map[string]interface{}) error {
	var sql string
	for key, value := range m {
		sql = fmt.Sprintf("update users set %s = '%v' where user_id = ?", key, value)
		fmt.Println(sql)
		stmt, err := db.Prepare(sql)
		defer stmt.Close()
		if err != nil {
			log.Panic(err)
		}
		_, err = stmt.Exec(m["user_id"])
		if err != nil {
			log.Printf("detele failed, err:%v\n", err)
			return err
		}
	}
	return nil
}

func UpdateEmployee(m map[string]interface{}) error {
	var sql string
	for key, value := range m {
		if key == "user_id" {
			continue
		}
		sql = fmt.Sprintf("update employee set %s = '%v' where id = ?", key, value)
		fmt.Println(sql)
		stmt, err := db.Prepare(sql)
		defer stmt.Close()
		if err != nil {
			log.Panic(err)
		}
		_, err = stmt.Exec(m["user_id"])
		if err != nil {
			log.Printf("detele failed, err:%v\n", err)
			return err
		}
	}
	return nil
}

func QueryUser(u *User) (*User, error) {
	var us User
	sql := "select * from users inner join employee on users.user_id = employee.id where (user_id = ?)"
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(u.User_id)
	for rows.Next() {
		err := rows.Scan(&us.User_id, &us.User_name, &us.Age, &us.Address, &us.Tel,
			&us.Department, &us.ID, &us.Dept_name, &us.Position, &us.Induction_time)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		log.Printf("detele failed, err:%v\n", err)
		return nil, err
	}
	return &us, nil
}
