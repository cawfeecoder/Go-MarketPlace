package controllerUser

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/labstack/echo"
	"github.com/nfrush/Go-MarketPlace/models/user"
	"github.com/nfrush/Go-MarketPlace/services/jwt"
	"github.com/nfrush/Go-MarketPlace/services/user"
)

//CreateUser creates a new user
func CreateUser(c echo.Context) error {
	u := &modelUser.User{}

	if err := c.Bind(u); err != nil {
		return c.JSON(409, err)
	}
	fmt.Println("Bind Successful")

	if err := servicesUser.CreateUser(u); err != nil {
		return c.JSON(409, err)
	}
	fmt.Println("Create User Successful")

	return c.JSON(200, u)
}

//FindAllUser find all application users
func FindAllUser(c echo.Context) error {
	if c.Request().Header().Get("Authorization") != "" {
		r, _ := regexp.Compile("((Bearer )*)")
		var token = r.ReplaceAllString(c.Request().Header().Get("Authorization"), "")
		auth, err := servicesJWT.RequiresAuth(token)
		if err != nil {
			return err
		}
		if auth {
			users := servicesUser.FindAllUser()
			return c.JSON(200, users)
		}
		return c.JSON(409, "Error")
	}
	return c.JSON(409, "Error")
}

//FindOneUser finds one application user
func FindOneUser(c echo.Context) error {
	user := servicesUser.FindOneUser(c.Param("username"))
	return c.JSON(200, user)
}

//UpdateUser updates a single user
func UpdateUser(c echo.Context) error {
	u := &modelUser.User{}

	if err := c.Bind(u); err != nil {
		return err
	}
	fmt.Println("Bind Successful")

	if err := servicesUser.UpdateUser(u); err != nil {
		return err
	}
	fmt.Println("User Update Successful")

	return errors.New("User has been updated.")
}

//DeleteUser deletes the specified user
func DeleteUser(c echo.Context) error {
	user := servicesUser.DeleteUser(c.Param("username"))
	return c.JSON(200, user)
}
