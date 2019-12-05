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

func WebServerStart(db *sql.DB) {
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
				c.SetCookie(username, username, 10, "/login", "127.0.0.1", false, true)
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

func Registe(db *sql.DB, c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if IsRegiste(db, "user", username) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "账号已注册！",
		})
	} else {
		InsertField(db, username, password)
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

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
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

	stmt, err := db.Query("select * from message where pid=" + strconv.Itoa(pid))
	CheckError(err)
	for stmt.Next() {
		stmt.Scan(&id, &pid, &username, &message, &time)
		idList = append(idList, id)
	}
	//for _, v := range idList {
	//	fmt.Println(v)
	//}
	return idList
}

func IdToMsg(db *sql.DB, c *gin.Context, id int) {
	var pid int
	var username, message, time string

	stmt, err := db.Query("select * from message where id=" + strconv.Itoa(id))
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

func GetPostPid(c *gin.Context) (int) {
	pidTem := c.DefaultPostForm("pid", "0")
	//fmt.Println(pidTem)
	pid, _ := strconv.Atoi(pidTem)
	//fmt.Println(pid)
	return pid
}

func IdToPidId(db *sql.DB, c *gin.Context, id int) (idList [] int, err error) {
	var pid int
	var username, message, time string

	stmt, err := db.Query("select * from message where pid=" + strconv.Itoa(id))
	CheckError(err)
	for stmt.Next() {
		stmt.Scan(&id, &pid, &username, &message, &time)
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
