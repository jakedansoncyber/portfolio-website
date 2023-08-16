package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/jakedansoncyber/portfolio-website/internal/repository"
	"github.com/jakedansoncyber/portfolio-website/pkg"
	"html/template"
	"time"
)

const (
	IndexPath = "web/static/index.html"
)

func RegisterHandlers(r *gin.Engine, database *repository.Database) {
	r.GET("/ping", handleHealth)
	r.GET("/", func(c *gin.Context) {
		handleBase(c, database)
	})
	r.POST("/add-film/", func(c *gin.Context) {
		handleRewrite(c, database)
	})
}

func handleHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// handlerBase - returns the index.html template, with film data
func handleBase(c *gin.Context, database *repository.Database) {

	var films []pkg.Film

	database.GetAll(&films)

	tmpl := template.Must(template.ParseFiles(IndexPath))

	filmsMap := &map[string][]pkg.Film{
		"Films": films,
	}

	err := tmpl.Execute(c.Writer, filmsMap)
	if err != nil {
		return
	}
}

// handlerBase - returns the index.html template, with film data
// updates database
func handleRewrite(c *gin.Context, database *repository.Database) {
	time.Sleep(1 * time.Second)
	title, _ := c.GetPostForm("title")
	director, _ := c.GetPostForm("director")
	film := &pkg.Film{
		Director: director,
		Title:    title,
	}
	database.Create(film)
	tmpl := template.Must(template.ParseFiles(IndexPath))
	err := tmpl.ExecuteTemplate(c.Writer, "film-list-element", film)
	if err != nil {
		return
	}
}
