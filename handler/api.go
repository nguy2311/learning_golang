package handler

import (
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler UserHandler
}

func (api *API) SetupRouter() {

	api.Echo.POST("/user/signin", api.UserHandler.HandlerSignIn)
	api.Echo.POST("/user/signup", api.UserHandler.HandlerSignUp)

}
