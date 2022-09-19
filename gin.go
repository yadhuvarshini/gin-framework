package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.New()
	s.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	s.GET("/market/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "You asked to get item %s", c.Param("id"))
	})
	s.PUT("/market/:id", func(c *gin.Context) {
		c.String(http.StatusAccepted, "You asked to edit item %s", c.Param("id"))
	})
	log.Fatal(http.ListenAndServe(":8080", s))
}
