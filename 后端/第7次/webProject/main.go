package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"webProject/basicFeatures"
)

func main() {
	db := basicFeatures.OpenDatabase("root", "mima", "chatBord")
	defer db.Close()
	r := gin.Default()

	r.POST("/registe", func(c *gin.Context) {
		basicFeatures.Registe(db, c)
	})
	r.POST("/login", func(c *gin.Context) {
		basicFeatures.Login(db, c)
	})

	r.Run()

}