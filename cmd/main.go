package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jakedansoncyber/portfolio-website/internal"
	"github.com/jakedansoncyber/portfolio-website/internal/repository"
	"log"
)

func main() {
	db := repository.NewDatabase()
	router := gin.Default()
	internal.RegisterHandlers(router, db)
	log.Fatal(router.Run(":8080")) // listen and serve on 0.0.0.0:8080
}
