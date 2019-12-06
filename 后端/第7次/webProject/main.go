package main

import (
	_ "github.com/go-sql-driver/mysql"
	"webProject/basicFeatures"
)

func main() {
	basicFeatures.WebServerStart()
}