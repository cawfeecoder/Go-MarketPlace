package routers

import "github.com/labstack/echo"

//InitRoutes - Initalize The Router With Routes
func InitRoutes() *echo.Echo {
	router := echo.New()
	router = SetUserRoutes(router)
	router = SetAuthRoutes(router)
	return router
}
