package main

import "os"

func main() {
	a := App{}
	db := a.Initialize(os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"))
	r := a.SetupRouter(db)
	r.Run()
}
