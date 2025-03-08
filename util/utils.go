package util

import (
	"fmt"
	"go-rest-api/model"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PrepareTestDB() {
	conn := NewTestDB()
	defer fmt.Println("Test database migration succeded.")
	defer CloseTestDB(conn)
	conn.AutoMigrate(&model.User{}, &model.Task{})
}

func NewTestDB() *gorm.DB {
	url := "postgres://test_user:test_password@localhost:5434/test_db"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func CloseTestDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}

func CleanupTestDB(db *gorm.DB) {
	tables := []string{"tasks", "users"}

	for _, table := range tables {
		db.Exec("TRUNCATE TABLE " + table + " CASCADE")
	}
}

func CleanupTaskTable(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE tasks CASCADE")
}

func CleanupUserTabls(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE users CASCADE")
}

func NewJWTToken() (*jwt.Token, string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": 1.0,
		"exp":    time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("test-secret"))
	return token, tokenString
}
