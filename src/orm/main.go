package main

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var db *sql.DB

func init() {
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	db, _ = sql.Open("mysql", "root:qwer.1234@tcp(10.116.20.58:3306)/kds3?charset=utf8")
}

func main() {
	//增加数据
	stmt, err := db.Prepare(`INSERT student (name,age) values (?,?)`)
	res, err := stmt.Exec("wangwu", 26)
	id, err := res.LastInsertId()
	fmt.Println("自增id=", id)
	//修改数据
	stmt, err = db.Prepare(`UPDATE student SET age=? WHERE id=?`)
	res, err = stmt.Exec(21, 5)
	num, err := res.RowsAffected() //影响行数
	fmt.Println(num)
	//删除数据
	stmt, err = db.Prepare(`DELETE FROM student WHERE id=?`)
	res, err = stmt.Exec(5)
	num, err = res.RowsAffected()
	fmt.Println(num)



	rows,err:=db.Query(
		`select c.CompanyID,c.CompanyName,c.IsValid,c.ModifyTime 
				from company c WHERE c.IsValid=1  `)
	if err!=nil{
		panic(err)
	}
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}



}
