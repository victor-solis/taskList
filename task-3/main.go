package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := getDb()

	if db == nil {
		return
	}

	defer func(db *sql.DB) {
		db.Close()
	}(db)

	//插入
	//student := students{
	//	"张三",
	//	20,
	//	"三年级",
	//}

	//insert(db, student)
	//list := []students{
	//	students{
	//		"李四",
	//		13,
	//		"一年级",
	//	},
	//	students{
	//		"王五",
	//		15,
	//		"二年级",
	//	},
	//}
	//batchInsert(db, &list)

	//update(db)
	//
	//queryList, _ := list(db)
	//for i := range queryList {
	//	fmt.Println(queryList[i])
	//}
	//
	//delete(db)

	err := transaction(db, 1, 2, 100)
	if err != nil {
		println(err.Error())
	} else {
		println("ok")
	}
}

func insert(db *sql.DB, student students) {
	result, err := db.Exec("insert into students(name,age,grade) values(?,?,?)", student.name, student.age, student.grade)
	if err != nil {
		println("插入失败")
	} else {
		id, _ := result.LastInsertId()
		println("插入成功id=", id)
	}
}

func batchInsert(db *sql.DB, studentList *[]students) error {
	tx, _ := db.Begin()

	stmt, err := db.Prepare("insert into students(name,age,grade) values(?,?,?)")
	if err != nil {
		tx.Rollback()
		println(err.Error())
		return err
	}
	defer stmt.Close()

	for _, student := range *studentList {
		_, stmtErr := stmt.Exec(student.name, student.age, student.grade)
		if stmtErr != nil {
			tx.Rollback()
			return stmtErr
		}
	}
	tx.Commit()
	return nil

}

func list(db *sql.DB) ([]students, error) {

	rows, err := db.Query("select name,age,grade from students where age > ?", 18)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var studentList []students
	for rows.Next() {
		var student *students = new(students)
		rowErr := rows.Scan(&student.name, &student.age, &student.grade)
		if rowErr != nil {
			println(rowErr.Error())
			return nil, rowErr
		}
		studentList = append(studentList, *student)
	}
	return studentList, nil
}

func update(db *sql.DB) error {
	_, err := db.Exec("update students set grade = ? where name = ? ", "四年级", "张三")
	if err != nil {
	}
	return err
}

func delete(db *sql.DB) error {
	_, err := db.Exec("delete from students where age < ? ", 15)
	return err
}

func transaction(db *sql.DB, fromId int, toId int, balance float64) error {
	tx, _ := db.Begin()

	_, err := tx.Exec("update accounts set balance = balance - ? where id = ?", balance, fromId)
	if err != nil {

		return err
	}

	// 确保在函数退出时处理事务
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // 重新抛出panic
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var fromBalance float64
	err = tx.QueryRow("select balance from accounts where id = ?", fromId).Scan(&fromBalance)
	if err != nil {

		return err
	}
	if fromBalance < 0 {
		err = errors.New("转账余额不足")
		return err
	}
	_, err = tx.Exec("update accounts set balance = balance + ? where id = ?", balance, toId)
	if err != nil {

		return err
	}
	_, err = tx.Exec("insert into transactions(from_account_id,to_account_id,amount) values(?,?,?)", fromId, toId, balance)
	if err != nil {

		return err
	}
	tx.Commit()
	return nil
}

func getDb() *sql.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return nil
	}
	return db
}

type students struct {
	name  string
	age   int
	grade string
}
