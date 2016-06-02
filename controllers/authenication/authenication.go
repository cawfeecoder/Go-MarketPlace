package controllerAuth

import (
	"errors"
	"fmt"

	"github.com/labstack/echo"
	"github.com/nfrush/Go-MarketPlace/models/user"
	"github.com/nfrush/Go-MarketPlace/services/authenication"
)

//Login - Login user and get JWT
func Login(c echo.Context) error {
	u := &modelUser.User{}

	if err := c.Bind(u); err != nil {
		return c.JSON(409, err)
	}
	fmt.Println("Bind Successful")
	fmt.Println(u)

	token, err := servicesAuth.Login(u)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(409, err)
	}
	fmt.Println("Login Success")

	return c.JSON(200, token)
}

//Logout User
func Logout(c echo.Context) error {
	u := &modelUser.User{}

	if err := c.Bind(u); err != nil {
		return c.JSON(409, err)
	}
	fmt.Println("Bind Successful")
	fmt.Println(u)

	if err := servicesAuth.Logout(u); err != nil {
		fmt.Println(err.Error())
		return c.JSON(409, err)
	}
	fmt.Println("Logout Successful")

	return c.JSON(200, errors.New("Logout Successful"))
}

//TestKey - obtain test key
func TestKey(c echo.Context) error {
	token, err := servicesAuth.TestKey()
	if err != nil {
		return c.JSON(409, err)
	}
	return c.JSON(200, token)
}

//Refresh - ReIssue Token
func Refresh(c echo.Context) error {
	u := &modelUser.User{}

	if err := c.Bind(u); err != nil {
		return c.JSON(409, err)
	}
	fmt.Println("Bind Successful")
	fmt.Println(u)

	token, err := servicesAuth.Refresh(u)
	if err != nil {
		return c.JSON(409, err)
	}
	return c.JSON(200, token)
}
