package repository

import (
	"go-rest-api/model"
	"go-rest-api/util"
	"os"
	"testing"

	"gorm.io/gorm"
)

var (
	tr ITaskRepository
	db *gorm.DB
)

// テストのセットアップ
func TestMain(m *testing.M) {
	db = util.NewTestDB()
	tr = NewTaskRepository(db)

	util.PrepareTestDB()
	db.Exec("INSERT INTO users (id, email, password) VALUES (1, 'user1@text.com', 'password') ON CONFLICT (id) DO NOTHING")

	code := m.Run()

	util.CloseTestDB(db)

	os.Exit(code)
}

func TestCreateTask(t *testing.T) {
	defer util.CleanupTaskTable(db)

	task := model.Task{Title: "Test Task", UserId: 1}

	if err := tr.Create(&task); err != nil {
		t.Fatalf("Create task failed: %v", err)
	}

	var rec model.Task
	db.First(&rec)

	if rec.UserId != task.UserId {
		t.Errorf("Expected UserID %d, got %d", task.UserId, rec.UserId)
	}
	if rec.Title != task.Title {
		t.Errorf("Expected Title %s, got %s", task.Title, rec.Title)
	}
}

func TestUpdateTask(t *testing.T) {
	defer util.CleanupTaskTable(db)

	task := model.Task{Title: "Test Task", UserId: 1}
	db.Create(&task)

	updatedTask := model.Task{Title: "Updated Title", UserId: 1}
	if err := tr.Update(&updatedTask, 1, task.ID); err != nil {
		t.Fatalf("Update task failed: %v", err)
	}

	var rec model.Task
	db.First(&rec)

	if rec.Title != updatedTask.Title {
		t.Errorf("Expected Title %s, got %s", updatedTask.Title, rec.Title)
	}
}

func TestGetAllTasks(t *testing.T) {
	defer util.CleanupTaskTable(db)

	db.Create(&model.Task{Title: "Test Title1", UserId: 1})
	db.Create(&model.Task{Title: "Test Title2", UserId: 1})

	var tasks []model.Task
	if err := tr.GetAll(&tasks, 1); err != nil {
		t.Fatalf("GetAll task failed: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestGetTaskById(t *testing.T) {
	defer util.CleanupTaskTable(db)

	db.Create(&model.Task{ID: 1, Title: "Test Title1", UserId: 1})

	var task model.Task
	if err := tr.GetByID(&task, 1, 1); err != nil {
		t.Fatalf("GetById task failed: %v", err)
	}
	if task.ID != 1 {
		t.Fatalf("Expected id 1 got %d", task.ID)
	}
	if task.Title != "Test Title1" {
		t.Fatalf("Expected title Test Title1 got %s", task.Title)
	}
}
