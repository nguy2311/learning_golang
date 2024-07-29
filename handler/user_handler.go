package handler

import (
	"learning_golang/model"
	"learning_golang/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (u *UserHandler) HandlerSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Ryan",
		"email": "huyn13@gmail.com",
	})
}

func (u *UserHandler) HandlerSignUp(c echo.Context) error {
	req := req.ReqSignup{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
			Data:       nil,
		})
	}
	type User struct {
		Email    string
		Fullname string
		Age      string
	}
	user := User{
		Email:    "huyn13@gmail.com",
		Fullname: "Ryan",
		Age:      "20",
	}
	return c.JSON(http.StatusOK, user)
}
