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
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// static files
	e.Static("/static", "static")

	t := &Template{
    templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	// TODO put tester to session
	var tester *model.Tester

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.POST("/intro", func(c echo.Context) error {
		tester = model.New()
		id := c.FormValue("id")
		tester.ID = id
		return c.Render(http.StatusOK, "intro.html", nil)
	})

	e.GET("/start", func(c echo.Context) error {
		pid := tester.Start()
		return c.Render(http.StatusOK, "pictures.html", pid)
	})

	e.GET("/finish", func(c echo.Context) error {
		pid := tester.Finish()
		return c.Render(http.StatusOK, "finish.html", pid)
	})

	e.GET("/keep", func(c echo.Context) error {
		pid, isRoundOver := tester.Keep()
		if isRoundOver {
			return c.Render(http.StatusOK, "roundover.html", nil)
		}
		return c.Render(http.StatusOK, "pictures.html", pid)
	})

	e.GET("/giveup", func(c echo.Context) error {
		pid, isAllOver, isRoundOver := tester.Giveup()
		if isAllOver {
			pid = tester.Finish()
			return c.Render(http.StatusOK, "finish.html", pid)
		}
		if isRoundOver {
			return c.Render(http.StatusOK, "roundover.html", nil)
		}
		return c.Render(http.StatusOK, "pictures.html", pid)
	})

	e.GET("/roundover", func(c echo.Context) error {
		pid := tester.CurrentPicture.Value
		return c.Render(http.StatusOK, "pictures.html", pid)
	})

	e.GET("/questions", func(c echo.Context) error {
		q1 := []string{
			"请基于你在刚才图片选择任务中的感受，回答下面的问题:",
			"你对你的最终选择是否满意？",
			"1（非常不满意）…2…3…4（一般）…5…6…7（非常满意）",
		}
		return c.Render(http.StatusOK, "questions.html", q1)
	})

	// try to get heroku port first
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Port
	}	
	log.Println("Service started at http://127.0.0.1:"+port)
	e.Logger.Fatal(e.Start(":"+port))
}
