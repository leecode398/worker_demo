package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type User struct {
	user_name  string `json:"user_name"`
	address    string `json:"address"`
	department string `json:"department"`
	user_id    uint32 `json:"user_id"`
	age        uint8  `json:"age"`
	tel        uint32 `json:"tel"`
}

type Employee struct {
	induction_time time.Time `json:"induction_time"`
	dept_name      string    `json:"dept_name"`
	position       string    `json:"position"`
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
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
	//配置数据库连接的限制.
	// db.SetConnMaxLifetime(config.ConnMaxLifetime)
	// db.SetMaxIdleConns(config.MaxIdleConns)
	// db.SetMaxOpenConns(config.MaxOpenConns)

	//测试是否连接成功.
	if err = db.Ping(); err != nil {
		panic(err)
	}

	//预处理mysql语句
	// createUser = dbPrepare(db, "INSERT INTO users (user_id, user_name, age, address, tel, department) values (?, ?, ?, ?, ?, ?)")
	// createEmployee = dbPrepare(db, "INSERT INTO employee (user_name, nick_name) values (?, ?)")
	// deleteEmployee = dbPrepare(db, "DELETE FROM users WHERE (user_id = ?, nick_name) values (?, ?)")
	// deleteEmployee = dbPrepare(db, "DELETE FROM employee WHERE (user_name, nick_name) values (?, ?)")

	fmt.Println("mysql init done.")
}

func (u User) AddUser() error {
	sqlStr := "INSERT INTO users (user_id, user_name, age, address, tel, department) values (?, ?, ?, ?, ?, ?)"
	ret, err := db.Exec(sqlStr, u.user_id, u.user_name, u.age, u.address, u.tel, u.department)
	if err != nil {
		log.Printf("insert failed, err:%v\n", err)
		return err
	}
	return nil
}
