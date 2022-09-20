package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TaskStatus int64

const (
	InProgress TaskStatus = iota
	Completed
)

func (s TaskStatus) String() string {
	switch s {
	case InProgress:
		return "IN_PROGRESS"
	case Completed:
		return "COMPLETED"
	}
	return "unknown"
}

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
	Image       string     `json:"image"`
	Status      TaskStatus `json:"status"`
}

func (t *Task) InsertTask(db *sql.DB) error {
	sql := "INSERT INTO tasks(title,description,created_at,image,status) VALUES($1,$2,transaction_timestamp(),$3,$4) RETURNING id"
	err := db.QueryRow(sql, t.Title, t.Description, t.Image, t.Status.String()).Scan(&t.ID)
	if err != nil {
		return err
	}
	return nil

}
func (t *Task) UpdateTask(db *sql.DB) error {
	sql := "UPDATE tasks SET title = $1 , description=$2,image=$3,status=$4 WHERE id = $5"
	_, err := db.Exec(sql, t.Title, t.Description, t.Image, t.Status.String(), t.ID)
	if err != nil {
		return err
	}
	return nil

}
func (t *Task) DeleteTask(db *sql.DB) error {
	sql := "DELETE FROM tasks WHERE id = $1"
	_, err := db.Exec(sql, t.ID)
	if err != nil {
		return err
	}
	return nil
}
func (t *Task) FindOneTask(db *sql.DB) error {
	sql := "SELECT id,title,description,created_at,image,status FROM tasks WHERE id = $1"
	return db.QueryRow(sql, t.ID).Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.Image, &t.Status)
}
func (t *Task) FindAllTask(db *sql.DB) ([]Task, error) {
	sql := "SELECT id,title,description,created_at,image,status FROM tasks"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	for rows.Next() {
		var task Task
		if rows.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.Image, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil

}
