package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
	"song-finder/server/models"
	"time"
)

func main() {
	e := echo.New()

	e.File("/admin/login", "./login.html")
	e.File("/admin", "./admin.html")
	e.Static("/static", "./static")

	e.GET("/songs", func(c echo.Context) error {
		time.Sleep(2000 * time.Millisecond)
		songs := make([]struct {
			Name string
		}, 0)
		return c.JSON(http.StatusOK, songs)
	})

	e.POST("/admin/login", func(c echo.Context) error {
		admin := &models.Admin{}
		if err := c.Bind(admin); err != nil {
			return err
		}

		if admin.IsValid() {
			return c.Redirect(http.StatusMovedPermanently, "/admin")
		}
		//TODO
		return c.Redirect(http.StatusMovedPermanently, "/admin")
	})
	e.Run(standard.New(":7777"))
}
