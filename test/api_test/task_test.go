package apitest

import (
	"bytes"
	"go-rest-api/controller"
	"go-rest-api/repository"
	"go-rest-api/test/util"
	"go-rest-api/usecase"
	"go-rest-api/validator"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	e  *echo.Echo
	tc controller.ITaskController
	db *gorm.DB
)

// テストのセットアップ
func TestMain(m *testing.M) {
	e = echo.New()

	// テスト用のDBでDI
	db = util.NewTestDB()
	tr := repository.NewTaskRepository(db)
	tv := validator.NewTaskValidator()
	tu := usecase.NewTaskUseCase(tr, tv)
	tc = controller.NewTaskController(tu)

	util.PrepareTestDB()
	db.Exec("INSERT INTO users (id, email, password) VALUES (1, 'user1@text.com', 'password') ON CONFLICT (id) DO NOTHING")

	code := m.Run()

	util.CleanupTestDB(db)
	util.CloseTestDB(db)

	os.Exit(code)
}

func TestCreateTask(t *testing.T) {
	token, tokenString := util.NewJWTToken()

	taskJSON := `{"title": "title"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte(taskJSON)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", token)

	if err := tc.CreateTask(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestGetTaskByID(t *testing.T) {
	token, tokenString := util.NewJWTToken()

	req := httptest.NewRequest(http.MethodGet, "/tasks/1", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", token)
	c.SetParamNames("taskId")
	c.SetParamValues("1")

	db.Exec("INSERT INTO tasks (id, title, user_id) VALUES (1, 'title1', 1) ON CONFLICT (id) DO NOTHING")

	if err := tc.GetTaskByID(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestGetAllTasks(t *testing.T) {
	token, tokenString := util.NewJWTToken()

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", token)

	if err := tc.GetAllTasks(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestUpdateTask(t *testing.T) {
	token, tokenString := util.NewJWTToken()

	taskJSON := `{"title": "updated"}`
	req := httptest.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer([]byte(taskJSON)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", token)
	c.SetParamNames("taskId")
	c.SetParamValues("1")

	if err := tc.UpdateTask(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestDeleteTask(t *testing.T) {
	token, tokenString := util.NewJWTToken()

	req := httptest.NewRequest(http.MethodDelete, "/tasks/1", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", token)
	c.SetParamNames("taskId")
	c.SetParamValues("1")

	if err := tc.DeleteTask(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}
