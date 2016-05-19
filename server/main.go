package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
	"song-finder/server/constants"
	"song-finder/server/models"
)

func main() {
	db, err := gorm.Open("mysql", constants.DBUsername+":"+constants.DBPassword+"@/"+constants.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Album{}, &models.Singer{}, &models.Song{})

	e := echo.New()

	e.File("/admin/login", "./login.html")
	e.File("/admin", "./admin.html")
	e.Static("/static", "./static")

	e.GET("/songs", func(c echo.Context) error {
		var songs []models.Song
		models.GetAllSong(&songs)
		return c.JSON(http.StatusOK, songs)
	})

	e.POST("/songs", func(c echo.Context) error {
		song := new(models.Song)
		if err := c.Bind(song); err != nil {
			return nil
		}
		song.Create()
		return c.JSON(http.StatusOK, *song)

	})

	e.POST("/admin/login", func(c echo.Context) error {
		admin := new(models.Admin)
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
