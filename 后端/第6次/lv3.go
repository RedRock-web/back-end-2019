//查询选修课共有多少人选
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	var sm studentTem
	var sumList [100] string
	//打开数据库
	db, err := sql.Open("mysql", "root:mima@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	Check(err)
	defer db.Close()
	//创建数据库
	CreateDb(db, "create database student")
	//打开创建的数据库
	db, err = sql.Open("mysql", "root:mima@tcp(127.0.0.1:3306)/student?charset=utf8")
	Check(err)
	//创建student表格
	CreateDb(db, "create table student (id int not null auto_increment, name varchar (100), sno int, class int, elective varchar(20), primary key (id))")
	//循环获取数据
	for i := 2019211525; i <= 2019211565; i++ {
		//获取基础信息
		url1 := "http://jwzx.cqupt.edu.cn/data/json_StudentSearch.php?searchKey=" + strconv.Itoa(i)
		resp1, _ := http.Get(url1)
		body1, _ := ioutil.ReadAll(resp1.Body)
		err := json.Unmarshal(body1, &sm)
		Check(err)
		//获取选修课
		url2 := "http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + strconv.Itoa(i)
		resp2, _ := http.Get(url2)
		body2, _ := ioutil.ReadAll(resp2.Body)
		reg := regexp.MustCompile(`(大学体育1-).*(</td)`)
		temp := string(reg.Find(body2))
		elective := string(temp)[14 : len(temp)-4]

		//fmt.Println(sm.ReturnData[0].Xm)
		// 将数据存于表格中
		name := sm.ReturnData[0].Xm
		sno := sm.ReturnData[0].Xh
		class := sm.ReturnData[0].Bj
		InsertDb(db, "insert into student (name, sno, class, elective) values(?,?,?,?)", name, sno, class, elective)
	}

	var sum string
	i := 1

	//创建选修课表格
	CreateDb(db, "create table elective (id int not null auto_increment, elective varchar(20), num int, primary key (id))")
    //获取去重后的选修课名单
	electiveList := SelectDb(db, "select distinct elective from student")

	for _, value := range electiveList {
		//插入选修课名字
		oder1 := "insert into elective (elective) value(\"" + value + "\")"
		//fmt.Println(oder1)
		stmt1, err := db.Prepare(oder1)
		Check(err)
		stmt1.Exec()
		//获取选修课次数
		oder2 := "select sum(elective = \"" + value + "\") from student;"
		//fmt.Println(oder2)
		rows, err := db.Query(oder2)
		Check(err)
        //更新选修课次数
		for rows.Next() {
			rows.Scan(&sum)
			sumList[i] = sum
			//fmt.Println(sum)
			oder3 := "update elective set num = " + sum + " where id = " + strconv.Itoa(i)
			//fmt.Println(oder3)
			temp1, err := db.Prepare(oder3)
			Check(err)
			temp1.Exec()
			i ++
		}

	}

	var input string
    fmt.Println("请输入想要查询的选修课")
	fmt.Scanf("%s", &input)
	//fmt.Println(input)
	for k, v :=  range electiveList {
		if v == input {
			fmt.Println("选择" + input + "选修课的一共有" + string(sumList[k+1]) + "人")
		}
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

func SelectDb(db *sql.DB, selectOder string) (elective []string) {
	var temp [100]string

	rows, err := db.Query(selectOder)
	Check(err)
	for i := 0; rows.Next(); i++ {
		err := rows.Scan(&temp[i])
		Check(err)
		elective = append(elective, temp[i])
	}
	return elective
}

func InsertDb(db *sql.DB, AlterOder string, name string, sno string, class string, elective string) {
	stmt, err := db.Prepare(AlterOder)
	Check(err)
	stmt.Exec(string(name), string(sno), string(class), string(elective))
	//defer stmt.Close()
}

func CreateDb(db *sql.DB, createDb string) {
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
