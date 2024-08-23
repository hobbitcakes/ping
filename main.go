package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type message struct {
	Message string `json:"message"`
}

var msg = message{Message : "pong"}

func ping(c *gin.Context) {

	c.IndentedJSON(http.StatusOK,msg) 
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.Run(":8080")
}
