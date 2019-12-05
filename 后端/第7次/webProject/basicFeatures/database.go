package basicFeatures

import "database/sql"

func InsertField(db *sql.DB, username string, password string) {
	stmt, err := db.Prepare("insert into user(username, password) values(?,?)")
	CheckError(err)
	stmt.Exec(username, password)
}

func CreateDatabase(db *sql.DB)  {
    
}

func OpenDatabase(username string, password string, databaseName string) (db *sql.DB) {
	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+databaseName+"?charset=utf8")
	CheckError(err)
	return db
}

func CreateTable()  {
	
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
