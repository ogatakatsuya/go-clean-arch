package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
	"go-rest-api/util"
	"os"

	"gorm.io/gorm"
)

func main() {
	var dbConn *gorm.DB
	if os.Getenv("GO_ENV") == "test" {
		dbConn = util.NewTestDB()
	} else {
		dbConn = db.NewDB()
	}
	defer fmt.Println("Successfully migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
