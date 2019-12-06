package basicFeatures

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
	"time"
)

func WebServerStart() {
	var username string

	db1 := OpenDatabase("root", "mima", "mysql")
	defer db1.Close()
	CreateDatabase(db1, "messageBoard")
	db := OpenDatabase("root", "mima", "messageBoard")
	defer db.Close()
	CreateTable(db, "user", "username varchar(100), password varchar(100), authority int not null default 0")
	CreateTable(db, "message", "pid int not null default 0, username varchar(100), message varchar(100),"+
		"likeNum int not null default 0, time varchar(100)")

	r := gin.Default()

	r.POST("/registe", func(c *gin.Context) {
		Registe(db, c, "user")
	})
	r.POST("/login", func(c *gin.Context) {
		username = Login(db, c)
	})
	r.POST("/login/msg/sendMsg", func(c *gin.Context) {
		SendMsg(db, c, username)
	})
	r.POST("/login/msg/getMsg", func(c *gin.Context) {
		pid := GetPostPid(c)
		GetMsg(db, c, pid)
	})
	r.POST("/login/logout", func(c *gin.Context) {
		Logout(db, c, username)
	})
	r.POST("/login/root/deleteMsg", func(c *gin.Context) {
		if IsAdminAuthority(db, c) {
			DeleteMsg(db, c)
		}
	})
	r.POST("/login/msg/like", func(c *gin.Context) {
		Like(db, c, username, )
	})
	r.Run()
}

func Login(db *sql.DB, c *gin.Context) string {
	username := c.PostForm("username")
	password := c.DefaultPostForm("password", "admin")
	if IsRegiste(db, "user", username) {
		_, err := c.Cookie(username)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "欢迎回来！" + username,
			})
		} else {
			_, passwd := DatabaseSearch(db, "user", username)
			if passwd == password {
				//暂时不支持中文cookie，如果用户名是中文，则会报错
				c.SetCookie(username, username, 100, "/login", "127.0.0.1", false, true)
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

//暂时出了点问题，无法成功删除cookie
func Logout(db *sql.DB, c *gin.Context, username string) {
	_, err := c.Cookie(username)
	if err == nil {
		c.SetCookie(username, username, -1, "/login", "127.0.0.1", false, true)
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "账号已登出！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "您并未登录，无需登出！",
		})
	}
}

func Registe(db *sql.DB, c *gin.Context, tableName string) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if IsRegiste(db, "user", username) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "账号已注册！",
		})
	} else {
		InsertField(db, username, password, tableName)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "注册成功！",
		})
	}
}

func IsRegiste(db *sql.DB, tableName string, username string) bool {
	var id string
	var name string
	var passwd string
	var judge bool
	var authority int

	selectOder := "select * from " + tableName + " where username= \"" + username + "\""
	//fmt.Println(a)
	stmt, err := db.Query(selectOder)
	CheckError(err)
	for stmt.Next() {
		stmt.Scan(&id, &name, &passwd, &authority)
	}
	if name != "" {
		judge = true
	} else {
		judge = false
	}
	return judge
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

		id := TimeToId(db, timeNow)
		//fmt.Println(id)
		tableName := "messageLikeUser" + id
		CreateTable(db, tableName, "username varchar(100)")
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

func GetMsgByPid(db *sql.DB, c *gin.Context) (int) {
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
	return pid
}

func PidToId(db *sql.DB, c *gin.Context, pid int) (idList [] int) {
	var id int
	var username, message, time string
	var likeNum int

	stmt, err := db.Query("select * from message where pid=" + strconv.Itoa(pid))
	CheckError(err)
	for stmt.Next() {
		stmt.Scan(&id, &pid, &username, &message, &time, &likeNum)
		idList = append(idList, id)
	}
	//for _, v := range idList {
	//	fmt.Println(v)
	//}
	return idList
}

func IdToMsg(db *sql.DB, c *gin.Context, id int) {
	var pid int
	var likeNum int
	var username, message, time string

	stmt, err := db.Query("select * from message where id=" + strconv.Itoa(id))
	CheckError(err)
	for stmt.Next() {
		stmt.Scan(&id, &pid, &username, &message, &time, &likeNum)
		_, err := c.Cookie(username)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"time":     time,
				"username": username,
				"message":  message,
				"like":     likeNum,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"time":     time,
				"username": username,
				"message":  message,
			})
		}

	}
}

func GetPostPid(c *gin.Context) (int) {
	pidTem := c.DefaultPostForm("pid", "0")
	//fmt.Println(pidTem)
	pid, _ := strconv.Atoi(pidTem)
	//fmt.Println(pid)
	return pid
}

func IdToPidId(db *sql.DB, c *gin.Context, id int) (idList [] int, err error) {
	var pid int
	var likeNum int
	var username, message, time string

	stmt, err := db.Query("select * from message where pid=" + strconv.Itoa(id))
	CheckError(err)
	for stmt.Next() {
		stmt.Scan(&id, &pid, &username, &message, &time, &likeNum)
		idList = append(idList, id)
	}

	return idList, err
}

func GetMsg(db *sql.DB, c *gin.Context, pid int) {
	fmt.Println(pid)
	idList := PidToId(db, c, pid)
	for k, id1 := range idList {
		c.String(200, strconv.Itoa(k+1)+"楼\n")
		IdToAllMsg(db, c, id1)
	}
}

func IdToAllMsg(db *sql.DB, c *gin.Context, id int) {
	IdToMsg(db, c, id)
	idList, err := IdToPidId(db, c, id)
	if err == nil {
		for _, id1 := range (idList) {
			IdToAllMsg(db, c, id1)
		}
	}
}

func IsAdminAuthority(db *sql.DB, c *gin.Context) (flag bool) {
	_, err := c.Cookie("root")
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "管理员,你好！",
		})
		flag = true
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "请登录管理员账号！",
		})
		flag = false
	}
	return flag
}

func DeleteMsg(db *sql.DB, c *gin.Context) {
	id := c.PostForm("id")
	stmt, err := db.Prepare("delete from message where id = ?")
	CheckError(err)
	stmt.Exec(id)
	c.JSON(200, gin.H{
		"status":  200,
		"message": "删除成功！",
	})
}

func Like(db *sql.DB, c *gin.Context, username string) {
	_, err := c.Cookie(username)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "登录后才能点赞！",
		})
	} else {
		var tableName, message, time1, id string
		var pid, likeNum int
		tableName = "message"

		id = c.PostForm("id")
		messageLikeUserName := "messageLikeUser" + id

		stmt, err1 := db.Query("select * from " + tableName + " where id = " + id)
		CheckError(err1)
		for stmt.Next() {
			stmt.Scan(&id, &pid, &username, &message, &likeNum, &time1)
		}
		if IsLike(db, username, messageLikeUserName) {
			likeNum--
			oder := "update " + tableName + " set likeNum = " + strconv.Itoa(likeNum) + " where id = " + id
			//fmt.Println(oder)
			stmt1, err2 := db.Prepare(oder)
			stmt1.Exec()
			oder2 := "delete from " + messageLikeUserName + " where username =  \"" + username + " \""
			fmt.Println(oder2)
			stmt3, err3 := db.Prepare(oder2)
			CheckError(err3)
			stmt3.Exec()
			if err2 == nil {
				c.JSON(200, gin.H{
					"message": "取消点赞成功！",
				})
			}
		} else {
			likeNum++
			oder := "update " + tableName + " set likeNum = " + strconv.Itoa(likeNum) + " where id = " + id
			//fmt.Println(oder)
			stmt1, err2 := db.Prepare(oder)
			stmt1.Exec()
			oder2 := "insert into " + messageLikeUserName + " (username) values( \"" + username + " \")"
			fmt.Println(oder2)
			stmt3, err3 := db.Prepare(oder2)
			CheckError(err3)
			stmt3.Exec()
			if err2 == nil {
				c.JSON(200, gin.H{
					"message": "点赞成功！",
				})
			}
		}
	}
}
func IsLike(db *sql.DB, username string, messageLikeUserName string) bool {
	var id, username1 string
	oder := "select * from " + messageLikeUserName + " where username = \"" + username + "\""
	stmt, err := db.Query(oder)
	CheckError(err)
	for stmt.Next() {
		stmt.Scan(&id, &username1)
	}
	if id == "" {
		return false
	} else {
		return true
	}
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
