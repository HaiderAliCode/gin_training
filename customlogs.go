package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	//this disables console colors
	// gin.DisableConsoleColor()
	//this enforce to colorise output
	// gin.ForceConsoleColor()

	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp,
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "ping")
	})
	router.Run(":8080")

}
