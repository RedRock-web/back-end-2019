
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	var sm studentTem
	//打开数据库
	db, err := sql.Open("mysql", "root:密码@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	Check(err)
	defer db.Close()
	//创建数据库
	CreateDb(db, "create database student")
    //打开创建的数据库
	db, err = sql.Open("mysql", "root:密码@tcp(127.0.0.1:3306)/student?charset=utf8")
	Check(err)
	//创建student表格
    CreateDb(db, "create table student (id int not null auto_increment, name varchar (100), sno int, class int, primary key (id))")
    //循环获取数据
	for i := 2019211525; i <= 2019211565; i++ {
		//fmt.Println(i)
		url := "http://jwzx.cqupt.edu.cn/data/json_StudentSearch.php?searchKey=" + strconv.Itoa(i)
		//fmt.Println(url)
		resp, _ := http.Get(url)
		body, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(body, &sm)
		Check(err)
		//fmt.Println(sm.ReturnData[0].Xm)
		// 将数据存于表格中
		name := sm.ReturnData[0].Xm
		sno := sm.ReturnData[0].Xh
		class := sm.ReturnData[0].Bj
		InsertDb(db, "insert into student (name, sno, class) values(?,?,?)", name, sno, class)
	}

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
		fmt.Println("id:", s.id, "name:", s.name, "class:", s.class, "sno:", s.sno)
	}

}

func InsertDb(db *sql.DB, AlterOder string, name string, sno string, class string) {
	stmt, err := db.Prepare(AlterOder)
	Check(err)
	stmt.Exec(string(name), string(sno), string(class))
	//defer stmt.Close()
}

func CreateDb(db *sql.DB, createDb string)  {
	stmt, err := db.Prepare(createDb)
	Check(err)
	stmt.Exec()
}


type studentTem struct {
	Code       int    `json:"code"`
	Info       string `json:"info"`
	ReturnData []struct {
		Xh    string      `json:"xh"`
		Xm    string      `json:"xm"`
		XmEn  interface{} `json:"xmEn"`
		Xb    string      `json:"xb"`
		Bj    string      `json:"bj"`
		Zyh   string      `json:"zyh"`
		Zym   string      `json:"zym"`
		Yxh   string      `json:"yxh"`
		Yxm   string      `json:"yxm"`
		Nj    string      `json:"nj"`
		Csrq  string      `json:"csrq"`
		Xjzt  string      `json:"xjzt"`
		Rxrq  string      `json:"rxrq"`
		Yxmen string      `json:"yxmen"`
		ZymEn string      `json:"zymEn"`
		Xz    int         `json:"xz"`
		Mz    string      `json:"mz"`
	} `json:"returnData"`
}
