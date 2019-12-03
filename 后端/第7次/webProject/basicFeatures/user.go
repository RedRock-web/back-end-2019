package basicFeatures

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

func Login(db *sql.DB, c *gin.Context) string {
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
				//暂时不支持中文cookie，如果用户名是中文，则会报错
				c.SetCookie(username, "loginSucessedly!", 100, "/", "127.0.0.1", false, true)
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
	return username
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

func SendMsg(db *sql.DB, c *gin.Context, username string) {
	_, err := c.Cookie(username)
	//CheckError(err)
	if err == nil {
		message := c.PostForm("message")
		pid1 := c.DefaultPostForm("pid", "0")
		timeNow := time.Now().Format("2006-01-02/03:04:05")

		stmt, err := db.Prepare("insert into message(pid,username,message,time) values(?, ?, ?, ?)")
		CheckError(err)
		stmt.Exec(pid1, username, message, timeNow)

		c.JSON(http.StatusOK, gin.H{
			"message": "发送成功！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"error":  "没有登录，不能发送消息！",
		})
	}
}

func GetMsg(db *sql.DB, c *gin.Context) {
	var pid, id int
	var username, message, time string

	//_, err := c.Cookie(username)
	toId := c.DefaultPostForm("pid", "0")
	stmt, err := db.Query("select * from message where pid=" + toId)
	CheckError(err)
	for stmt.Next() {
		stmt.Scan(&id, &pid, &username, &message, &time)
		c.JSON(http.StatusOK, gin.H{
			"time":     time,
			"username": username,
			"message":  message,
		})
	}
}

func WebServerStart(db *sql.DB)  {
	var username string
	r := gin.Default()

	r.POST("/registe", func(c *gin.Context) {
		Registe(db, c)
	})
	r.POST("/login", func(c *gin.Context) {
		username = Login(db, c)
	})
	r.POST("/login/sendMsg", func(c *gin.Context) {
		SendMsg(db, c, username)
	})
	r.POST("/login/getMsg", func(c *gin.Context) {
		GetMsg(db, c)
	})
	r.Run()
}