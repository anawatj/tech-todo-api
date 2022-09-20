package routes

import (
	"database/sql"
	"net/http"
	"tech-exam-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskRoute struct {
	DB *sql.DB
}

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
func (tr *TaskRoute) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, "Json is invalid")
		return
	}
	if err := task.InsertTask(tr.DB); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, task)
}
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

func (tr *TaskRoute) PutTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, "Json is invalid")
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
