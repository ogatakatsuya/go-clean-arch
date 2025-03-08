package validator

import (
	"go-rest-api/model"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskValidator_Success(t *testing.T) {
	tv := NewTaskValidator()
	task := model.Task{
		Title: "title",
	}
	err := tv.TaskValidate(task)
	assert.Nil(t, err)
}

func TestTaskValidator_TitleNil_Failure(t *testing.T) {
	tv := NewTaskValidator()
	task := model.Task{
		Title: "",
	}
	err := tv.TaskValidate(task)
	assert.NotNil(t, err)
	assert.Equal(t, "title: title is requred.", err.Error())
}

func TestTaskValidator_TitleMax_Failure(t *testing.T) {
	tv := NewTaskValidator()
	task := model.Task{
		Title: strings.Repeat("a", 101),
	}
	err := tv.TaskValidate(task)
	assert.NotNil(t, err)
	assert.Equal(t, "title: limited max 100 char.", err.Error())
}
