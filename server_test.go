package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"tech-exam-api/models"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAllTaskSuccess(t *testing.T) {
	app := App{}
	db := app.Initialize(os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"))
	var task models.Task
	task.Title = "test"
	task.Description = "test test"
	task.Image = ""
	task.Status = models.InProgress
	db.QueryRow("INSERT INTO tasks(title,description,created_at,image,status) VALUES($1,$2,transaction_timestamp(),$3,$4) RETURNING id", task.Title, task.Description, task.Image, task.Status).Scan(&task.ID)

	r := app.SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
	db.Exec("DELETE FROM tasks")
}
func TestGetAllTaskNotFound(t *testing.T) {
	app := App{}
	db := app.Initialize(os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"))

	r := app.SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.NotNil(t, w.Body)
}

func TestCreateTaskSuccess(t *testing.T) {
	app := App{}
	db := app.Initialize(os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"))

	r := app.SetupRouter(db)

	body := []byte(`{
		"title": "tao",
		"description": "tao",
		"image": "",
		"status":"IN_PROGRESS"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/tasks", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotNil(t, w.Body)

	var resp map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &resp)

	// assert field of response body
	assert.Nil(t, err)
	assert.Equal(t, "tao", resp["title"])
	assert.Equal(t, "tao", resp["description"])
	assert.Equal(t, "IN_PROGRESS", resp["status"])
	db.Exec("DELETE FROM tasks")
}

func TestGetTaskByIdSuccess(t *testing.T) {
	app := App{}
	db := app.Initialize(os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"))
	var task models.Task
	task.Title = "test"
	task.Description = "test test"
	task.Image = ""
	task.Status = models.InProgress
	db.QueryRow("INSERT INTO tasks(title,description,created_at,image,status) VALUES($1,$2,transaction_timestamp(),$3,$4) RETURNING id", task.Title, task.Description, task.Image, task.Status).Scan(&task.ID)

	r := app.SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/"+task.ID.String(), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
	db.Exec("DELETE FROM tasks")
}

func TestGetTaskByIdNotFound(t *testing.T) {
	app := App{}
	db := app.Initialize(os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"))

	r := app.SetupRouter(db)
	id := uuid.New()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/"+id.String(), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
}

func TestUpdateTaskByIdSuccess(t *testing.T) {
	app := App{}
	db := app.Initialize(os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"))
	var task models.Task
	task.Title = "test"
	task.Description = "test test"
	task.Image = ""
	task.Status = models.InProgress
	db.QueryRow("INSERT INTO tasks(title,description,created_at,image,status) VALUES($1,$2,transaction_timestamp(),$3,$4) RETURNING id", task.Title, task.Description, task.Image, task.Status).Scan(&task.ID)
	body := []byte(`{
		"title": "tao",
		"description": "tao",
		"image": "",
		"status":"COMPLETE"
	}`)
	r := app.SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/tasks/"+task.ID.String(), bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)

	db.Exec("DELETE FROM tasks")
}

func TestDeleteTaskByIdSuccess(t *testing.T) {
	app := App{}
	db := app.Initialize(os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"))
	var task models.Task
	task.Title = "test"
	task.Description = "test test"
	task.Image = ""
	task.Status = models.InProgress
	db.QueryRow("INSERT INTO tasks(title,description,created_at,image,status) VALUES($1,$2,transaction_timestamp(),$3,$4) RETURNING id", task.Title, task.Description, task.Image, task.Status).Scan(&task.ID)

	r := app.SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/tasks/"+task.ID.String(), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)

	db.Exec("DELETE FROM tasks")
}
