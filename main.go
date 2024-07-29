package main

import (
	"learning_golang/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	db := &db.Database{
		Host:     "cluster0.lidujz8.mongodb.net",
		Username: "huynq3",
		Password: "huynq3Cy",
		DbName:   "Cluster0",
	}

	db.Connect()
	defer db.Disconnect()

}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
