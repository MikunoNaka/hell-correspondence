package main

import (
//	"log"
	"net/http"
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
	//	log.Fatal("$PORT must be set")
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")
	router.Static("/media", "media")

	router.GET("/en", func(c *gin.Context) {
		c.HTML(http.StatusOK, "en.index.html", nil)
	})

	router.GET("/jp", func(c *gin.Context) {
		c.HTML(http.StatusOK, "jp.index.html", nil)
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "redirect.index.html", nil)
	})

	router.GET("/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new.index.html", nil)
		name := c.PostFormArray("name")
	})

	router.GET("/site-down", func(c *gin.Context) {
		c.HTML(http.StatusOK, "maintain.html", nil)
	})
	router.Run(":" + port)
}
