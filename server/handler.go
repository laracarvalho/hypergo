package main

import (
	"fmt"
	"net/http"

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

	fmt.Print("PUT: ", username, email)

	return c.Redirect(303, "/profile?email="+email+"&username="+username)
}
