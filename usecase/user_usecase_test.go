package usecase

import (
	"errors"
	"go-rest-api/model"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct {
	mock.Mock
}

func (mr *MockUserRepository) GetByEmail(user *model.User, email string) error {
	args := mr.Called(user, email)
	return args.Error(0)
}

func (mr *MockUserRepository) Create(user *model.User) error {
	args := mr.Called(user)
	return args.Error(0)
}

func newMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

type MockUserValidator struct {
	mock.Mock
}

func (mv *MockUserValidator) UserValidate(user model.User) error {
	args := mv.Called(user)
	return args.Error(0)
}

func newMockUserValidator() *MockUserValidator {
	return &MockUserValidator{}
}

func TestSignUp_Success(t *testing.T) {
	mr := newMockUserRepository()
	mv := newMockUserValidator()
	mr.On("Create", mock.Anything).Return(nil)
	mv.On("UserValidate", mock.Anything).Return(nil)

	uu := NewUserUsecase(mr, mv)

	res, err := uu.SignUp(model.User{Email: "user@test.com", Password: "password"})

	assert.NoError(t, err)
	assert.Equal(t, "user@test.com", res.Email, "Email should match the input email")
	mv.AssertCalled(t, "UserValidate", mock.Anything)
	mr.AssertCalled(t, "Create", mock.Anything)
}

func TestSignUp_Respository_Failed(t *testing.T) {
	mr := newMockUserRepository()
	mv := newMockUserValidator()
	mr.On("Create", mock.Anything).Return(nil)
	mv.On("UserValidate", mock.Anything).Return(errors.New("error"))

	uu := NewUserUsecase(mr, mv)

	_, err := uu.SignUp(model.User{Email: "user@test.com", Password: "password"})
	assert.Error(t, err)
}

func TestSignUp_Validator_Failed(t *testing.T) {
	mr := newMockUserRepository()
	mv := newMockUserValidator()
	mr.On("Create", mock.Anything).Return(errors.New("error"))
	mv.On("UserValidate", mock.Anything).Return(nil)

	uu := NewUserUsecase(mr, mv)

	_, err := uu.SignUp(model.User{Email: "user@test.com", Password: "password"})
	assert.Error(t, err)
}

func TestLogin_Success(t *testing.T) {
	mr := newMockUserRepository()
	mv := newMockUserValidator()

	mv.On("UserValidate", mock.Anything).Return(nil)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	mr.On("GetByEmail", mock.AnythingOfType("*model.User"), "user@test.com").
		Run(func(args mock.Arguments) {
			user := args.Get(0).(*model.User)
			*user = model.User{
				ID:       1,
				Email:    "user@test.com",
				Password: string(hashedPassword),
			}
		}).
		Return(nil)

	uu := NewUserUsecase(mr, mv)

	os.Setenv("SECRET", "testsecret")
	defer os.Unsetenv("SECRET")

	token, err := uu.Login(model.User{Email: "user@test.com", Password: "password"})

	assert.NoError(t, err)
	assert.NotEmpty(t, token, "Token should not be empty")
	mv.AssertCalled(t, "UserValidate", mock.Anything)
	mr.AssertCalled(t, "GetByEmail", mock.AnythingOfType("*model.User"), "user@test.com")
}

func TestLogin_Fail_UserValidateError(t *testing.T) {
	mr := newMockUserRepository()
	mv := newMockUserValidator()
	mv.On("UserValidate", mock.Anything).Return(errors.New("validation error"))

	uu := NewUserUsecase(mr, mv)

	_, err := uu.Login(model.User{Email: "user@test.com", Password: "password"})

	assert.Error(t, err)
	assert.Equal(t, "validation error", err.Error())

	mv.AssertCalled(t, "UserValidate", mock.Anything)
	mr.AssertNotCalled(t, "GetByEmail", mock.Anything)
}

func TestLogin_Fail_UserNotFound(t *testing.T) {
	mr := newMockUserRepository()
	mv := newMockUserValidator()
	mv.On("UserValidate", mock.Anything).Return(nil)
	mr.On("GetByEmail", mock.Anything, "user@test.com").Return(errors.New("user not found"))

	uu := NewUserUsecase(mr, mv)

	_, err := uu.Login(model.User{Email: "user@test.com", Password: "password"})

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())

	mv.AssertCalled(t, "UserValidate", mock.Anything)
	mr.AssertCalled(t, "GetByEmail", mock.Anything, "user@test.com")
}

func TestLogin_PassMisMatch_Failure(t *testing.T) {
	mr := newMockUserRepository()
	mv := newMockUserValidator()

	mv.On("UserValidate", mock.Anything).Return(nil)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	mr.On("GetByEmail", mock.AnythingOfType("*model.User"), "user@test.com").
		Run(func(args mock.Arguments) {
			user := args.Get(0).(*model.User)
			*user = model.User{
				ID:       1,
				Email:    "user@test.com",
				Password: string(hashedPassword),
			}
		}).
		Return(nil)

	uu := NewUserUsecase(mr, mv)

	os.Setenv("SECRET", "testsecret")
	defer os.Unsetenv("SECRET")

	_, err := uu.Login(model.User{Email: "user@test.com", Password: "wrongpass"})

	assert.Error(t, err)
}
