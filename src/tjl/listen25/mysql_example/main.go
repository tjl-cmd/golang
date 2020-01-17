package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDb() error {
	var err error
	dsn := "root:root@tcp(localhost:3306)/golang"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//最大连接池
	db.SetMaxOpenConns(100)

	//最小连接池
	db.SetMaxIdleConns(16)
	return nil
}

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"string"`
	Age  int    `db:"age"`
}

//查询操作
func testQueryData() {
	for i := 0; i < 101; i++ {
		fmt.Printf("query %d times\n", i)
		sqlstr := "select id ,name,age from user where id=?"

		row := db.QueryRow(sqlstr, 3)
		/*if row != nil {
			continue
		}*/
		var user User
		err := row.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("Scan failed err:%v\n", err)
			return
		}
		fmt.Printf("id:%d,anme:%s age :%d\n", user.Id, user.Name, user.Age)
	}
}
func testQqeryMultilRow() {
	//sqlstr := "select id ,name,age from user where id=?"
	sqlstr := "select id ,name,age from user where id>?"

	row, err := db.Query(sqlstr, 0)
	//rows对象一定要关闭
	defer func() {
		if row != nil {
			row.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Scan failed err:%v\n", err)
		return
	}
	for row.Next() {
		var user User
		err := row.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("Scan failed err:%v\n", err)
			return
		}
		//fmt.Printf("id:%d,anme:%s age :%d\n", user.Id, user.Name, user.Age)
		fmt.Printf("user:%#v\n", user)
	}

}

//插入操作
func testInsertData() {
	sqlstr := "insert into user(name,age) values(?,?)"
	result, err := db.Exec(sqlstr, "tom", 19)
	if err != nil {
		fmt.Printf("insert failed err :%v\n", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id failed ,err :%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)
}

//更新
func testUpdateData() {
	sqlstr := "update user set name=? where id =?"
	result, err := db.Exec(sqlstr, "jim", 3)
	if err != nil {
		fmt.Printf("insert failed err :%v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows failed err :%v\n", err)
	}
	fmt.Printf("update db success agected rows:%d \n", affected)
}

//删除操作
func testDeleteData() {
	sqlstr := "delete from user where id =?"
	result, err := db.Exec(sqlstr, 3)
	if err != nil {
		fmt.Printf("insert failed err :%v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows failed err :%v\n", err)
	}
	fmt.Printf("update db success agected rows:%d \n", affected)
}
func testPrepareData() {
	sqlstr := "select id,name age from user where id > ?"
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed ,err :%v\n", err)
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	rows, err := stmt.Query(0)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Scan failed err:%v\n", err)
		return
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("Scan failed err:%v\n", err)
			return
		}
		//fmt.Printf("id:%d,anme:%s age :%d\n", user.Id, user.Name, user.Age)
		fmt.Printf("user:%#v\n", user)
	}

}
func testPrepareInsertData() {
	sqlstr := "insert into user(name,age) values(?,?)"
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("insert failed err :%v\n", err)
		return
	}
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	result, err := stmt.Exec("jim", 100)
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id failed ,err :%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)
}
func testTrans() {
	conn, err := db.Begin()
	if err != nil {
		if conn != nil {
			conn.Rollback()
		}
		fmt.Printf("begin failed ,err :%v\n", err)
		return
	}
	sqlstr := "update user set age = 22 where id = ?"
	_, err = conn.Exec(sqlstr, 1)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec sql:%s failed err :%v\n", sqlstr, err)
		return
	}
	sqlstr = "update user set age  = 102 where id = ? "
	_, err = conn.Exec(sqlstr, 2)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec sql:%s failed err :%v\n", sqlstr, err)
		return
	}
	err = conn.Commit()
	if err != nil {
		fmt.Printf("commit faile err:%v\n", err)
		conn.Rollback()
		return
	}
}
func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init db failed err :%v\n", err)
		return
	}
	//testQueryData()
	//testQqeryMultilRow()
	//testInsertData()
	//testUpdateData()
	//testDeleteData()
	//testPrepareData()
	//testPrepareInsertData()
	testTrans()
}
