package main

import (
	"learning_golang/db"
	"learning_golang/handler"
	"learning_golang/repository/repo_impl"

	"github.com/labstack/echo/v4"
)

func main() {

	db := &db.Database{
		Host:     "cluster0.lidujz8.mongodb.net",
		Username: "huynq3",
		Password: "huynq3Cy",
		AppName:  "Cluster0",
		DbName:   "learning_golang",
	}

	db.Connect()

	defer db.Disconnect()

	e := echo.New()
	UserHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(db),
	}
	api := handler.API{
		Echo:        e,
		UserHandler: UserHandler,
	}
	api.SetupRouter()
	e.Logger.Fatal(e.Start(":8080"))
}

// func home(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }
