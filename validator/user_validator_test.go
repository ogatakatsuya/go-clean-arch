package validator

import (
	"go-rest-api/model"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidator_Success(t *testing.T) {
	uv := NewUserValidator()
	user := model.User{
		Email:    "user@test.com",
		Password: "password",
	}
	err := uv.UserValidate(user)

	assert.Nil(t, err)
}

func TestUserValidator_EmailNil_Fsilure(t *testing.T) {
	uv := NewUserValidator()
	user := model.User{
		Email:    "",
		Password: "password",
	}
	err := uv.UserValidate(user)

	assert.NotNil(t, err)
	assert.Equal(t, "email: email is required.", err.Error())
}

func TestUserValidator_EmailMax_Failure(t *testing.T) {
	uv := NewUserValidator()
	user := model.User{
		Email:    strings.Repeat("a", 31),
		Password: "password",
	}

	err := uv.UserValidate(user)

	assert.NotNil(t, err)
	assert.Equal(t, "email: limited max 30 char.", err.Error())
}

func TestUserValidator_InvalidEmail_Failure(t *testing.T) {
	uv := NewUserValidator()
	user := model.User{
		Email:    "invalid_email",
		Password: "password",
	}

	err := uv.UserValidate(user)

	assert.NotNil(t, err)
	assert.Equal(t, "email: is not valid email format.", err.Error())
}

func TestUserValidator_PasswordNil_Failure(t *testing.T) {
	uv := NewUserValidator()
	user := model.User{
		Email:    "user@test.com",
		Password: "",
	}

	err := uv.UserValidate(user)

	assert.NotNil(t, err)
	assert.Equal(t, "password: password is required.", err.Error())
}

func TestUserValidator_PasswordMin_Failure(t *testing.T) {
	uv := NewUserValidator()
	user := model.User{
		Email:    "user@test.com",
		Password: "pass",
	}

	err := uv.UserValidate(user)

	assert.NotNil(t, err)
	assert.Equal(t, "password: limited min 6 max 30 char.", err.Error())
}

func TestUserValidator_PasswordMax_Failure(t *testing.T) {
	uv := NewUserValidator()
	user := model.User{
		Email:    "user@test.com",
		Password: strings.Repeat("a", 31),
	}

	err := uv.UserValidate(user)

	assert.NotNil(t, err)
	assert.Equal(t, "password: limited min 6 max 30 char.", err.Error())
}
