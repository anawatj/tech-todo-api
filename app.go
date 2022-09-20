package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "tech-exam-api/docs"
	"tech-exam-api/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct{}

// @title           Gin Todo Service
// @version         1.0
// @description     A book management service API in Go using Gin framework.
// @termsOfService  https://tos.santoshk.dev

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func (a *App) SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	var tr routes.TaskRoute
	tr.DB = db
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
