package routers

import (
	"github.com/labstack/echo"
	"github.com/nfrush/Go-MarketPlace/controllers/authenication"
)

//SetAuthRoutes - Authenication
func SetAuthRoutes(router *echo.Echo) *echo.Echo {

	router.POST("/login", controllerAuth.Login)

	router.POST("/login/test", controllerAuth.TestKey)

	router.POST("/logout", controllerAuth.Logout)

	router.POST("/refresh", controllerAuth.Refresh)

	return router
}
