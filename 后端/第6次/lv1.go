package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:密码@tcp(127.0.0.1:3306)/student?charset=utf8")
	Check(err)
	//AlterDb(db, "insert into student (name, sno, class) values('haha',123,321)")
	var s student
	SelectDb(db, "select * from student", s)
	//AlterDb(db, "update student set sno = 333 where id = 4")
    AlterDb(db, "delete from student where id = 2")
	defer db.Close()
}

func Check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}


type student struct {
	id    int
	name  string
	sno   int
	class int
}

func SelectDb(db *sql.DB, selectOder string, s student) {
	rows, err := db.Query(selectOder)
	Check(err)
	//defer db.Close()
	for rows.Next() {
		if err := rows.Scan(&s.id, &s.name, &s.sno, &s.class); err != nil {
			panic(err)
		}
		fmt.Println("id:",s.id, "name:",s.name, "class:",s.class, "sno:",s.sno)
	}

}

func AlterDb(db *sql.DB,AlterOder string)  {
	stmt, err := db.Prepare(AlterOder)
	Check(err)
	stmt.Exec()
	//defer stmt.Close()
}
