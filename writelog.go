package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	//this disable color in output since we are writing this in a log file we dont need colors
	gin.DisableConsoleColor()

	//this writes all logs to a log file
	f, _ := os.Create("log.txt")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "i am home")
	})
	router.Run(":8080")
}
