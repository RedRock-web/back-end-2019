package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	userList := map[string]string {
		"admin": "admin",
		"root": "root",
		"lily": "lily",
		"lucy": "lucy",
	}

	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		//c.Request.Cookie()
		name := c.Query("name")
		if _, ok := userList[name]; ok {   //先判断账号是否存在
			_, err := c.Cookie(name)
			if err == nil {              //若存在,则通过cookie是否登录
				c.JSON(http.StatusOK, gin.H{
					"code": 200,
					"message": "hello  " + name,
				})
			} else {          //若没有登录,则登录并设置cookie
				c.SetCookie(name, "login sucessedly!", 100,
					"/login", "127.0.0.1", false, true)
				c.String(http.StatusOK, "登录成功!")
			}
		} else {           //若账号不存在,则注册账号
			userList[name] = "000000"
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"message": "hello guest",
			})
		}
	})
	r.Run()
}