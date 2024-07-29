package handler

import (
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler UserHandler
}

func (api *API) SetupRouter() {

	api.Echo.GET("/user/signin", api.UserHandler.HandlerSignIn)
	api.Echo.GET("/user/signup", api.UserHandler.HandlerSignUp)

}
