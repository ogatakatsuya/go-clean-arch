package repository

import (
	"go-rest-api/model"
	"go-rest-api/util"
	"testing"

	"gorm.io/gorm"
)

func setupTaskTestDB() *gorm.DB {
	db := util.NewTestDB()
	return db
}

func TestCreateTask(t *testing.T) {
	db := setupTaskTestDB()
	db.Exec("INSERT INTO users (id, email, password) VALUES (1, 'user1@testtask.com', 'password') ON CONFLICT (id) DO NOTHING")
	defer util.CloseTestDB(db)
	defer util.CleanupTaskTable(db)

	tr := NewTaskRepository(db)

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
	db := setupTaskTestDB()
	defer util.CloseTestDB(db)
	defer util.CleanupTaskTable(db)

	tr := NewTaskRepository(db)

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
	db := setupTaskTestDB()
	defer util.CloseTestDB(db)
	defer util.CleanupTaskTable(db)

	tr := NewTaskRepository(db)

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
	db := setupTaskTestDB()
	defer util.CloseTestDB(db)
	defer util.CleanupTaskTable(db)

	tr := NewTaskRepository(db)

	expected := model.Task{ID: 1, Title: "Test Title1", UserId: 1}
	db.Create(&expected)

	var actual model.Task
	if err := tr.GetByID(&actual, 1, expected.ID); err != nil {
		t.Fatalf("GetById task failed: %v", err)
	}
	if actual.ID != expected.ID {
		t.Fatalf("Expected id %d got %d", expected.ID, actual.ID)
	}
	if actual.Title != expected.Title {
		t.Fatalf("Expected title %s got %s", expected.Title, actual.Title)
	}
}
