package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db := getSqlxDb()
	empList, _ := sqlxList(db)
	fmt.Println(empList)
	emp, _ := get(db)
	fmt.Println(emp)
}

func getSqlxDb() *sqlx.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return nil
	}
	return db
}

func sqlxList(db *sqlx.DB) ([]Employees, error) {
	query := "select * from employees where department = ?"
	var employeesList []Employees
	err := db.Select(&employeesList, query, "技术部")
	if err != nil {
		return nil, err
	}
	return employeesList, nil
}
func get(db *sqlx.DB) (Employees, error) {
	query := "SELECT * FROM employees ORDER BY salary desc limit 1"
	var employees Employees
	err := db.Get(&employees, query)
	if err != nil {
		return Employees{}, err
	}
	return employees, nil
}

type Employees struct {
	Id         int     `db:"id" json:"id"`
	Name       string  `db:"name" json:"name"`
	Department string  `db:"department" json:"department"`
	Salary     float64 `db:"salary" json:"salary"`
}
