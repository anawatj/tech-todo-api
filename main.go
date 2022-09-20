package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"tech-exam-api/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Initialize(user, password, dbname, host string) *sql.DB {
	connectionString :=
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	var err error
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func main() {
	r := gin.Default()
	db := Initialize(os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"))
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
	r.Run()
}
