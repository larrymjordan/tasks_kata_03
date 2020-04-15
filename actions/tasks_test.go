package actions

import (
	"net/http"
	"tasks/models"
)

func (as *ActionSuite) Test_Tasks_Create() {
	taskStorage.Clear()
	task := models.Task{Description: "Do this"}
	resp := as.JSON("/tasks").Post(task)
	as.Equal(http.StatusCreated, resp.Code)

	createdTask := models.Task{}
	resp.Bind(&createdTask)
	as.Equal(task.Description, createdTask.Description)
	as.NotZero(createdTask.ID)

	as.Len(taskStorage.List(), 1)
}

func (as *ActionSuite) Test_Tasks_List() {
	taskStorage.Clear()
	resp := as.JSON("/tasks").Get()
	as.Equal(http.StatusOK, resp.Code)

	tasks := []models.Task{}
	resp.Bind(&tasks)
	as.Len(tasks, 0)

	taskStorage.Add(&models.Task{Description: "Do this"})
	resp = as.JSON("/tasks").Get()
	tasks = []models.Task{}
	resp.Bind(&tasks)
	as.Len(tasks, 1)
}
