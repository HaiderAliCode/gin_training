package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func loginEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "I am in home")
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("request", "client_request")
		//passing the request to next request
		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		fmt.Printf("%s::%s \t %s \t %s ", time.Now().Format("2006-01-02 15:04:05"), host, url, method)
		context.Next()
		fmt.Println(context.Writer.Status())
	}
}

func main() {
	router := gin.New()

	//global middleware
	// router.Use(Logger())
	//default logger
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//local middleware
	router.GET("/home", AuthRequired(), loginEndpoint)
	router.Run(":8080")
}
