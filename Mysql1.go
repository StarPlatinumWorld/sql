package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}else {
		fmt.Print("连接成功")
		fmt.Println("")
	}
	stmt, err := db.Prepare("CREATE TABLE user (id int NOT NULL AUTO_INCREMENT,name int(11),age varchar(40),PRIMARY KEY(id));")
	if err != nil {
		fmt.Println(err. Error)
	}else{
		fmt.Print("创建成功")
		fmt.Println("")
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err. Error)
	}
	insertDB(db)
	selectDB(db)
}
func insertDB(db *sql.DB){
	stmt,err:=db.Prepare("insert into user(name) value(?)")
	if err!=nil{
		log.Fatal(err)
	}else{
		fmt.Print("添加成功")
		fmt.Println("")
	}
	stmt.Exec("a,123")
}
func deleteDB(db *sql.DB){
	stmt,err:=db.Prepare("delete from user where name='?'")
	if err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println("删除成功")
		fmt.Println("")
	}
	stmt.Exec("a")
}
func updateDB(db *sql.DB){
	stmt,err:=db.Prepare("update user set name ='?' where id='1'")
	if err!=nil{
		log.Fatal(err)
	}else{
		fmt.Print("修改完成")
	}
	stmt.Exec("a,132")
}
func selectDB(db *sql.DB)  {
	stmt, err := db.Query("select * from user;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for stmt.Next(){
		var id int
		var name string
		err := stmt.Scan(&id,&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id,name)
	}

}