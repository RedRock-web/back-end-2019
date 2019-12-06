package basicFeatures

import (
	"database/sql"
	"fmt"
)

func InsertField(db *sql.DB, username string, password string, tableName string) {
	stmt, err := db.Prepare("insert into " + tableName +
		"(username, password) values(?,?)")
	CheckError(err)
	stmt.Exec(username, password)
}

func CreateDatabase(db *sql.DB, NewDataBaseName string) {
	stmt, err := db.Prepare("create database " + NewDataBaseName)
	CheckError(err)
	stmt.Exec()
}

func OpenDatabase(username string, password string, databaseName string) (db *sql.DB) {
	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+databaseName+"?charset=utf8")
	CheckError(err)
	return db
}

func CreateTable(db *sql.DB, tableName string, keysAndValues string) {
	//oder := "create table " + tableName + " (id int NOT NULL auto_increment," +
	//	keysAndValues + ", primary key(id))"
	//fmt.Println(oder)
	stmt, err := db.Prepare("create table " + tableName + " (id int NOT NULL auto_increment," +
		keysAndValues + ", primary key(id)) character set = utf8")
	CheckError(err)
	stmt.Exec()
}
func DatabaseSearch(db *sql.DB, tableName string, username string) (name string, passwd string) {
	var id string
	var authority int

	selectOder := "select * from " + tableName + " where username= \"" + username + "\""
	//fmt.Println(a)
	stmt, err := db.Query(selectOder)
	CheckError(err)
	for stmt.Next() {
		stmt.Scan(&id, &name, &passwd, &authority)
	}
	return name, passwd
}
func TimeToId(db *sql.DB, Time string) (id string) {
	var pid, username, message, likeNum string

	selectOder := "select * from message where time= \"" + Time + "\""
	stmt, err := db.Query(selectOder)
	CheckError(err)
	for stmt.Next() {
		stmt.Scan(&id, &pid, &username, &message, &likeNum, &Time)
	}
	fmt.Println(Time)
	return id
}

func UserId(db *sql.DB, username string) (id string) {
	var username1, password, authority string

	selectOder := "select * from user where username= \"" + username + "\""
	stmt, err := db.Query(selectOder)
	CheckError(err)
	for stmt.Next() {
		stmt.Scan(&id, &username1, &password, &authority)
	}
	return id
}
