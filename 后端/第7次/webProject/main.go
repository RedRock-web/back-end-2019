package main

import (
	_ "github.com/go-sql-driver/mysql"
	"webProject/basicFeatures"
)

func main() {
	db := basicFeatures.OpenDatabase("root", "mima", "chatBord")
	defer db.Close()
	basicFeatures.WebServerStart(db)
}