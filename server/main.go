package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"io"
	"net/http"
	"os"
	"song-finder/server/constants"
	"song-finder/server/models"
	"strconv"
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
		models.FindAll(&songs)
		return c.JSON(http.StatusOK, songs)
	})

	e.POST("/songs", func(c echo.Context) error {
		song := new(models.Song)
		song.Init(song)

		if err := c.Bind(song); err != nil {
			return err
		}
		song.Create()
		return c.JSON(http.StatusOK, song)
	})

	e.PUT("/songs/:id", func(c echo.Context) error {
		song := new(models.Song)
		song.Init(song)

		if err := c.Bind(song); err != nil {
			return err
		}

		song.Update()
		return c.JSON(http.StatusOK, song)
	})

	e.DELETE("/songs/:id", func(c echo.Context) error {
		var id uint64
		var err error
		if id, err = strconv.ParseUint(c.Param("id"), 10, 64); err != nil {
			return err
		}

		song := new(models.Song)
		song.Init(song)
		song.ID = uint(id)

		if err = song.Delete(); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, song.ID)
	})

	e.GET("/singers", func(c echo.Context) error {
		var singers []models.Singer
		models.FindAll(&singers)
		return c.JSON(http.StatusOK, singers)
	})

	e.POST("/singers", func(c echo.Context) error {
		singer := new(models.Singer)
		singer.Init(singer)

		if err := c.Bind(singer); err != nil {
			return err
		}

		singer.Create()
		return c.JSON(http.StatusOK, singer)
	})

	e.PUT("/singers/:id", func(c echo.Context) error {
		singer := new(models.Singer)
		singer.Init(singer)

		if err := c.Bind(singer); err != nil {
			return err
		}

		singer.Update()
		return c.JSON(http.StatusOK, singer)
	})

	e.DELETE("/singers/:id", func(c echo.Context) error {
		var id uint64
		var err error
		if id, err = strconv.ParseUint(c.Param("id"), 10, 64); err != nil {
			return err
		}

		singer := new(models.Singer)
		singer.Init(singer)
		singer.ID = uint(id)

		if err = singer.Delete(); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, singer.ID)
	})

	e.GET("/albums", func(c echo.Context) error {
		var albums []models.Album
		models.FindAll(&albums)
		return c.JSON(http.StatusOK, albums)
	})

	e.POST("/albums", func(c echo.Context) error {
		album := new(models.Album)
		album.Init(album)

		if err := c.Bind(album); err != nil {
			return err
		}

		album.Create()
		return c.JSON(http.StatusOK, album)
	})

	e.PUT("/albums/:id", func(c echo.Context) error {
		album := new(models.Album)
		album.Init(album)

		if err := c.Bind(album); err != nil {
			return err
		}

		album.Update()
		return c.JSON(http.StatusOK, album)
	})

	e.DELETE("/albums/:id", func(c echo.Context) error {
		var id uint64
		var err error
		if id, err = strconv.ParseUint(c.Param("id"), 10, 64); err != nil {
			return err
		}

		album := new(models.Album)
		album.Init(album)
		album.ID = uint(id)

		if err = album.Delete(); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, album.ID)
	})

	e.POST("/albums/:id/image-upload", func(c echo.Context) error {
		var id uint64
		var err error
		if id, err = strconv.ParseUint(c.Param("id"), 10, 64); err != nil {
			return err
		}

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create("./static/images/" + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		album := new(models.Album)
		album.Init(album)
		album.ID = uint(id)
		models.First(album)

		album.Photo = file.Filename
		album.Update()
		return c.JSON(http.StatusOK, album)
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
