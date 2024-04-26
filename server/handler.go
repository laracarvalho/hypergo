package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func IndexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"name": "HOME",
		"msg":  "Hello, World!",
	})
}

type ProfilePage struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Username string
	Email    string
}

func ProfileHandler(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")

	return c.Render(http.StatusOK, "profile", ProfilePage{
		Name: "EDIT PROFILE",
		Profile: Profile{
			Username: username,
			Email:    email,
		},
	})
}

func EditingProfileHandler(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")

	return c.Render(http.StatusOK, "editprofile", ProfilePage{
		Name: "EDITING PROFILE",
		Profile: Profile{
			Username: username,
			Email:    email,
		},
	})
}

func UpdateProfileHandler(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")

	// d1 := []byte(msg)
	// err := os.WriteFile("../msg.txt", d1, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	return c.Redirect(303, "/profile?email="+email+"&username="+username)
}

type GetPage struct {
	Name string
	Msg  string
}

func GetHandler(c echo.Context) error {
	//msg := c.QueryParam("msg")

	data, err := os.ReadFile("../msg.txt")
	if err != nil {
		fmt.Print(err)
	}

	msg := string(data[:])

	return c.Render(http.StatusOK, "get", GetPage{
		Name: "GET PAGE",
		Msg:  msg,
	})
}

func PostHandler(c echo.Context) error {
	msg := c.FormValue("msg")

	d1 := []byte(msg)
	err := os.WriteFile("../msg.txt", d1, 0644)
	if err != nil {
		panic(err)
	}

	return c.Redirect(303, "/get?msg"+msg)
}
