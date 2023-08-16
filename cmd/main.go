package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jakedansoncyber/portfolio-website/internal"
	"github.com/jakedansoncyber/portfolio-website/internal/repository"
	"github.com/jakedansoncyber/portfolio-website/pkg"
	"log"
)

func main() {
	db := repository.NewDatabase()
	db.Migrate(&pkg.Film{})
	router := gin.Default()
	internal.RegisterHandlers(router, db)
	log.Fatal(router.Run(":8080")) // listen and serve on 0.0.0.0:8080
}
