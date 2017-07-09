package main

import (
	"net/http"
	"log"
	"os"
	"html/template"

	"psyexp/config"
	"time"
	"github.com/labstack/echo"
)

var templates = template.Must(template.ParseFiles(
	"templates/index.html", 
	"templates/intro.html", 
	"templates/pictures.html", 
	"templates/finish.html", 
	"templates/header.html",
))

type Data struct {
	ID string
	DecisionTime time.Time
	PicturePicked int
}

type Page struct {
	Title string
}


func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	config.InitConfig()

	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// root router for test
  // r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	p := Page{"title"}
	// 	renderTemplate(w, "index", p)
  // })


	// try to get heroku port first
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Port
	}	
	log.Println("Service started at http://127.0.0.1:"+port)
	e.Logger.Fatal(e.Start(":"+port))
}
