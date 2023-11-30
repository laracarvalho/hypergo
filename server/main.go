package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Renderer = &TemplateRegistry{
		templates: ListTemplates(),
	}

	e.Use(middleware.Logger())

	// Route => handler
	e.GET("/", IndexHandler)
	e.GET("/profile", ProfileHandler)
	e.GET("/profile/edit", EditingProfileHandler)
	e.PUT("/profile", UpdateProfileHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the Echo server
	e.Start("0.0.0.0:" + port)
}
