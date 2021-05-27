package main

import (
	"first-api/Config"
	"first-api/Models"
	"first-api/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.GormUser{})
	r := Routes.SetupRouter()
	r.Run("121.126.223.225:8080")
}
