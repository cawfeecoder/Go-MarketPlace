package routers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nfrush/Go-MarketPlace/controllers/user"
	"github.com/nfrush/Go-MarketPlace/services/jwt"
)

//SetUserRoutes - Set Up User Routes
func SetUserRoutes(router *echo.Echo) *echo.Echo {
	//Force Signing Authorization of JWT
	var signingKey = servicesJWT.GetSigningKey()

	//Find All Users
	router.GET("/user", controllerUser.FindAllUser, middleware.JWT([]byte(signingKey)))
	//Create New User
	router.POST("/user", controllerUser.CreateUser)
	//Update User
	router.PUT("/user", controllerUser.UpdateUser, middleware.JWT([]byte(signingKey)))
	//Find Single User
	router.GET("/user/:username", controllerUser.FindOneUser, middleware.JWT([]byte(signingKey)))
	//Delete User
	router.DELETE("/user/:username", controllerUser.DeleteUser, middleware.JWT([]byte(signingKey)))

	return router
}
