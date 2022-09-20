package routes

import (
	"database/sql"
	"net/http"
	"strings"
	"tech-exam-api/models"

	_ "tech-exam-api/docs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskRoute struct {
	DB *sql.DB
}

// GetAllTask godoc
// @summary Get List Task
// @description get list Task
// @tags tasks
// @accept json
// @produce json
// @response 200 {object} models.Task "Ok"
// @Router /api/v1/tasks [get]
func (tr *TaskRoute) GetAllTask(c *gin.Context) {
	orderBy := c.Query("orderBy")
	orderType := c.Query("orderType")
	title := c.Query("title")
	description := c.Query("description")
	var task models.Task

	tasks, err := task.FindAllTask(tr.DB, orderBy, orderType, title, description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(tasks) == 0 {
		c.JSON(http.StatusNotFound, "Task Not Found")
		return
	}
	c.JSON(http.StatusOK, tasks)

}

// createTask godoc
// @summary Create Task
// @description Create Task
// @tags tasks
// @accept json
// @produce json
// @param Task body models.Task true "Task data to be created"
// @response 201 {object} models.Task "Created"
// @Router /api/v1/tasks [post]
func (tr *TaskRoute) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, "Json is invalid")
		return
	}
	errors := []string{}
	if task.Title == "" {
		errors = append(errors, "Title is required")
	}
	if task.Description == "" {
		errors = append(errors, "Description is required")
	}
	if task.Status == "" {
		errors = append(errors, "Status is required")
	}
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, strings.Join(errors, ","))
		return
	}
	if err := task.InsertTask(tr.DB); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, task)
}

// getTask godoc
// @summary Get Task
// @description Get task by id
// @tags tasks
// @id getTask
// @accept json
// @produce json
// @param id path string true "id of task to be updated"
// @response 200 {object} models.Task "OK"
// @Router /api/v1/tasks/:id [get]
func (tr *TaskRoute) GetOneTask(c *gin.Context) {
	id := c.Param("id")
	var task = models.Task{
		ID: uuid.MustParse(id),
	}
	err := task.FindOneTask(tr.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if task.Title == "" && task.Description == "" && task.Image == "" {
		c.JSON(http.StatusNotFound, "Task Not Found")
		return
	}
	c.JSON(http.StatusOK, task)
}

// putTask godoc
// @summary Update Task
// @description Update task by id
// @tags tasks
// @id updateTask
// @accept json
// @produce json
// @param id path string true "id of task to be updated"
// @param Task body models.Task true "Task data to be updated"
// @response 200 {object} models.Task "OK"
// @Router /api/v1/tasks/:id [put]
func (tr *TaskRoute) PutTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, "Json is invalid")
		return
	}
	errors := []string{}
	if task.Title == "" {
		errors = append(errors, "Title is required")
	}
	if task.Description == "" {
		errors = append(errors, "Description is required")
	}
	if task.Status == "" {
		errors = append(errors, "Status is required")
	}
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, strings.Join(errors, ","))
		return
	}
	task.ID = uuid.MustParse(id)
	err := task.UpdateTask(tr.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, task)

}

// deleteTask godoc
// @summary Delete Task
// @description Delete task by id
// @tags tasks
// @id deleteTask
// @accept json
// @produce json
// @param id path string true "id of task to be deleted"
// @response 200 {object} string "OK"
// @Router /api/v1/tasks/:id [delete]
func (tr *TaskRoute) RemoveTask(c *gin.Context) {
	id := c.Param("id")
	var task = models.Task{
		ID: uuid.MustParse(id),
	}
	err := task.DeleteTask(tr.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Task is deleted")
}
