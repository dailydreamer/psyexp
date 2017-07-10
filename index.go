package main

import (
	"net/http"
	"os"
	"io"
	"log"
	"html/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

		"psyexp/config"
		"psyexp/model"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}


func main() {
	config.InitConfig()

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// static files
	e.Static("/static", "static")

	t := &Template{
    templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	tester := model.Tester{}
	log.Println(tester)

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.POST("/intro", func(c echo.Context) error {
		id := c.FormValue("id")
		tester.ID = id
		return c.Render(http.StatusOK, "intro.html", nil)
	})

	e.GET("/pictures/:pid", func(c echo.Context) error {
		pid := c.Param("pid")
		if (pid == "") {
			pid = "0"
		}
		return c.Render(http.StatusOK, "pictures.html", pid)
	})

	e.POST("/finish", func(c echo.Context) error {
		tester.PicturePicked = c.FormValue("pid")
		return c.Render(http.StatusOK, "finish.html", nil)
	})

	// try to get heroku port first
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Port
	}	
	log.Println("Service started at http://127.0.0.1:"+port)
	e.Logger.Fatal(e.Start(":"+port))
}
