package main

import (
	"database/sql"
	"fmt"
	"log"
	"tech-exam-api/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type App struct{}

func (a *App) SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	var tr routes.TaskRoute
	tr.DB = db
	r.GET("/api/v1/tasks", tr.GetAllTask)
	r.GET("/api/v1/tasks/:id", tr.GetOneTask)
	r.POST("/api/v1/tasks", tr.CreateTask)
	r.PUT("/api/v1/tasks/:id", tr.PutTask)
	r.DELETE("/api/v1/tasks/:id", tr.RemoveTask)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
func (a *App) Initialize(user, password, dbname, host string) *sql.DB {
	connectionString :=
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	var err error
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
