package basicFeatures

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"

)

func Login(db *sql.DB, c *gin.Context) {
	username := c.PostForm("username")
	password := c.DefaultPostForm("password", "admin")
	if IsRegiste(db, "student", username) {
		_, err := c.Cookie(username)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "欢迎回来！" + username,
			})
		} else {
			_, passwd := SelectDatabase(db, "student", username)
			if passwd == password {
				c.SetCookie(username, "loginSucessedly!", 2, "/", "127.0.0.1", false, true)
				c.JSON(http.StatusOK, gin.H{
					"status":  http.StatusOK,
					"message": "登录成功！",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  http.StatusOK,
					"message": "密码错误！",
				})
			}
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "账号未注册！",
		})
	}
}

func Registe(db *sql.DB, c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if IsRegiste(db, "student", username) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "账号已注册！",
		})
	} else {
		InsertDatabase(db, username, password)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "注册成功！",
		})
	}
}

func InsertDatabase(db *sql.DB, username string, password string) {
	stmt, err := db.Prepare("insert into student(username, password) values(?,?)")
	CheckError(err)
	stmt.Exec(username, password)
}

func IsRegiste(db *sql.DB, tableName string, username string) bool {
	var id string
	var name string
	var passwd string
	var judge bool

	selectOder := "select * from " + tableName + " where username= \"" + username + "\""
	//fmt.Println(a)
	stmt, err := db.Query(selectOder)
	CheckError(err)
	for stmt.Next() {
		stmt.Scan(&id, &name, &passwd)
	}
	if name != "" {
		judge = true
	} else {
		judge = false
	}
	return judge
}

func OpenDatabase(username string, password string, databaseName string) (db *sql.DB) {
	//fmt.Println("mysql", username + ":" + password + "@tcp(127.0.0.1:3306)/" + databaseName + "?charset=utf8")
	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+databaseName+"?charset=utf8")
	CheckError(err)
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func SelectDatabase(db *sql.DB, tableName string, username string) (name string, passwd string) {
	var id string

	selectOder := "select * from " + tableName + " where username= \"" + username + "\""
	//fmt.Println(a)
	stmt, err := db.Query(selectOder)
	CheckError(err)
	for stmt.Next() {
		stmt.Scan(&id, &name, &passwd)
	}
	return name, passwd
}
