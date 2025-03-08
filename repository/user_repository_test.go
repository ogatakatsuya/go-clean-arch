package repository

import (
	"go-rest-api/model"
	"go-rest-api/util"
	"testing"

	"gorm.io/gorm"
)

func setupUserTestDB() *gorm.DB {
	db := util.NewTestDB()
	return db
}

func TestCreateUser(t *testing.T) {
	db := setupUserTestDB()
	defer util.CleanupTaskTable(db)
	defer util.CleanupUserTabls(db)

	ur := NewUserRepository(db)

	user := model.User{Email: "user1@test.com", Password: "testpass"}

	if err := ur.Create(&user); err != nil {
		t.Fatalf("Create task failed: %v", err)
	}

	var rec model.User
	db.Select("id, email").Where("email = ?", user.Email).First(&rec)

	if rec.Email != user.Email {
		t.Errorf("Expected Email %s, got %s", user.Email, rec.Email)
	}
}

func TestGetUserByEmail(t *testing.T) {
	db := setupUserTestDB()
	defer util.CleanupTaskTable(db)
	defer util.CleanupUserTabls(db)

	ur := NewUserRepository(db)

	expected := model.User{ID: 100, Email: "user1@testemail.com", Password: "testpass"}
	db.Create(&expected)

	var actual model.User
	if err := ur.GetByEmail(&actual, expected.Email); err != nil {
		t.Fatalf("GetById task failed: %v", err)
	}
	if actual.ID != expected.ID {
		t.Fatalf("Expected id %d got %d", expected.ID, actual.ID)
	}
	if actual.Email != expected.Email {
		t.Fatalf("Expected Email %s got %s", expected.Email, actual.Email)
	}
}
