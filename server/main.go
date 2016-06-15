package main

import (
	// "fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"song-finder/server/models"
	"strconv"
)

type searchResult struct {
	SongId     uint
	SongName   string
	Lyricist   string
	Composer   string
	SingerName string
	AlbumName  string
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var store = sessions.NewFilesystemStore("./session/", securecookie.GenerateRandomKey(20))

func authenticateAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := c.Request().(*standard.Request).Request

		session, _ := store.Get(request, "user")

		if is_admin, ok := session.Values["is_admin"].(bool); !ok || !is_admin {
			return c.Redirect(http.StatusMovedPermanently, "/admin/login")
		}

		return next(c)
	}
}

func main() {
	db, err := models.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Album{}, &models.Singer{}, &models.Song{})
	db.Close()

	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseFiles("./search.html", "./song-detail.html", "./song.html")),
	}
	e.SetRenderer(t)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				return err
			}
			request := c.Request().(*standard.Request).Request
			context.Clear(request)
			return nil
		}
	})

	e.File("/", "./index.html")
	e.File("/singer", "./singer.html")
	e.File("/album", "./album.html")
	e.File("/admin/login", "./login.html")

	e.GET("/song", func(c echo.Context) error {
		var songs []models.Song
		db, err := models.ConnectDB()
		defer db.Close()
		if err != nil {
			return err
		}

		db.Preload("Singer").Preload("Album").Find(&songs)
		return c.Render(http.StatusOK, "song", songs)
	})

	e.GET("/search", func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		searchWord := "%" + keyword + "%"

		results := make([]searchResult, 0)
		if keyword != "" {
			db, err := models.ConnectDB()
			defer db.Close()
			if err != nil {
				return err
			}

			rows, err := db.Raw(`SELECT songs.id, songs.name, songs.lyricist, songs.composer, singers.name, albums.name 
				                 FROM songs, singers, albums 
				                 WHERE (songs.singer_id = singers.id AND songs.album_id = albums.id) AND 
				                       (songs.lyricist LIKE ? OR songs.composer LIKE ? OR songs.name LIKE ? OR singers.name LIKE ? OR albums.name LIKE ?)`, searchWord, searchWord, searchWord, searchWord, searchWord).Rows()
			defer rows.Close()

			if err != nil {
				return err
			}

			var songId uint
			var songName, lyricist, composer, singerName, albumName string

			for rows.Next() {
				rows.Scan(&songId, &songName, &lyricist, &composer, &singerName, &albumName)
				results = append(results, searchResult{
					SongId:     songId,
					Lyricist:   lyricist,
					Composer:   composer,
					SongName:   songName,
					SingerName: singerName,
					AlbumName:  albumName,
				})
			}
		}
		return c.Render(http.StatusOK, "search", results)
	})

	e.GET("/song/:id", func(c echo.Context) error {
		var id uint64
		var err error
		if id, err = strconv.ParseUint(c.Param("id"), 10, 64); err != nil {
			return err
		}

		db, err := models.ConnectDB()

		defer db.Close()
		if err != nil {
			return err
		}

		song := models.Song{
			Singer: models.Singer{},
			Album:  models.Album{},
		}

		db.First(&song, id)
		db.Model(&song).Related(&song.Singer)
		db.Model(&song).Related(&song.Album)

		return c.Render(http.StatusOK, "song-detail", song)
	})

	e.GET("/admin", func(c echo.Context) error {
		file, err := ioutil.ReadFile("admin.html")
		if err != nil {
			return err
		}
		return c.HTML(http.StatusOK, string(file))
	}, authenticateAdmin)

	e.Static("/static", "./static")

	e.GET("/songs", func(c echo.Context) error {
		var songs []models.Song
		models.FindAll(&songs)
		return c.JSON(http.StatusOK, songs)
	}, authenticateAdmin)

	e.POST("/songs", func(c echo.Context) error {
		song := new(models.Song)
		song.Init(song)

		if err := c.Bind(song); err != nil {
			return err
		}
		song.Create()
		return c.JSON(http.StatusOK, song)
	}, authenticateAdmin)

	e.PUT("/songs/:id", func(c echo.Context) error {
		song := new(models.Song)
		song.Init(song)

		if err := c.Bind(song); err != nil {
			return err
		}

		song.Update()
		return c.JSON(http.StatusOK, song)
	}, authenticateAdmin)

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
	}, authenticateAdmin)

	e.GET("/singers", func(c echo.Context) error {
		var singers []models.Singer
		models.FindAll(&singers)
		return c.JSON(http.StatusOK, singers)
	}, authenticateAdmin)

	e.POST("/singers", func(c echo.Context) error {
		singer := new(models.Singer)
		singer.Init(singer)

		if err := c.Bind(singer); err != nil {
			return err
		}

		singer.Create()
		return c.JSON(http.StatusOK, singer)
	}, authenticateAdmin)

	e.PUT("/singers/:id", func(c echo.Context) error {
		singer := new(models.Singer)
		singer.Init(singer)

		if err := c.Bind(singer); err != nil {
			return err
		}

		singer.Update()
		return c.JSON(http.StatusOK, singer)
	}, authenticateAdmin)

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
	}, authenticateAdmin)

	e.GET("/albums", func(c echo.Context) error {
		var albums []models.Album
		models.FindAll(&albums)
		return c.JSON(http.StatusOK, albums)
	}, authenticateAdmin)

	e.POST("/albums", func(c echo.Context) error {
		album := new(models.Album)
		album.Init(album)

		if err := c.Bind(album); err != nil {
			return err
		}

		album.Create()
		return c.JSON(http.StatusOK, album)
	}, authenticateAdmin)

	e.PUT("/albums/:id", func(c echo.Context) error {
		album := new(models.Album)
		album.Init(album)

		if err := c.Bind(album); err != nil {
			return err
		}

		album.Update()
		return c.JSON(http.StatusOK, album)
	}, authenticateAdmin)

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
	}, authenticateAdmin)

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
	}, authenticateAdmin)

	e.POST("/admin/login", func(c echo.Context) error {
		admin := new(models.Admin)
		if err := c.Bind(admin); err != nil {
			return c.Redirect(http.StatusMovedPermanently, "/admin/login")
		}

		if !admin.IsValid() {
			return c.Redirect(http.StatusMovedPermanently, "/admin/login")
		}

		request := c.Request().(*standard.Request).Request
		responseWriter := c.Response().(*standard.Response).ResponseWriter
		session, _ := store.Get(request, "user")

		session.Values["is_admin"] = true
		session.Save(request, responseWriter)

		return c.Redirect(http.StatusMovedPermanently, "/admin")
	})

	e.Run(standard.New(":7777"))
}
