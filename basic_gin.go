package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func benchEndpoint(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("benchEndpoint"))
}

func MyBenchLogger() {

}

//go get -u github.com/gin-gonic/gin
func main() {
	// r := gin.Default()

	//simple routing examples
	// r.GET("/ping", func(c *gin.Context) {
	// c.JSON(200, gin.H{
	// 	"mode": "gin",
	// })
	// c.String(http.StatusOK, "hello world")
	// })

	//passing variable
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "hello %s", name)
	// })

	//additional routes
	// r.GET("/user/:name/*action", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	action := c.Param("action")

	// 	c.String(http.StatusOK, "hello %s with action %s", name, action, c.FullPath())
	// })

	// r.GET("/users/groups", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "The available groups are [...]")
	// })

	////querying parameters
	///welcome?firstname=Jane&lastname=Doe
	// r.GET("/welcome", func(c *gin.Context) {
	// 	//if not available then ali will be in it
	// 	firstName := c.DefaultQuery("firstname", "Ali")
	// 	//if not available then it will be empty
	// 	lastName := c.Query("lastname")

	// 	c.String(http.StatusOK, "hello %s %s", firstName, lastName)
	// })

	//multipart/urlencode
	//getting form data from api
	// r.POST("/formpost", func(c *gin.Context) {
	// 	message := c.PostForm("message")
	// 	name := c.DefaultPostForm("name", "Ali")

	// 	c.JSON(200, gin.H{
	// 		"status":  "posted",
	// 		"message": message,
	// 		"name":    name,
	// 	})
	// })

	//using query and postform together with post
	// r.POST("/post", func(c *gin.Context) {

	// 	id := c.Query("id")
	// 	page := c.DefaultQuery("page", "0")
	// 	name := c.PostForm("name")
	// 	message := c.PostForm("message")

	// 	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	// })

	//getting  /post?ids[a]=1234&ids[b]=hello
	//getting names[first]=thinkerou&names[second]=tianou
	// r.POST("/post", func(c *gin.Context) {
	// 	ids := c.QueryMap("ids")
	// 	names := c.PostFormMap("names")

	// 	fmt.Printf("%s, %s", ids, names)
	// })

	//uploading a file on server
	// r.MaxMultipartMemory = 8 << 20
	// r.POST("/upload", func(c *gin.Context) {

	// 	file, _ := c.FormFile("filename")

	// 	c.SaveUploadedFile(file, fmt.Sprintf("./%s", file.Filename))

	// 	c.String(http.StatusOK, "processing done")
	// })

	//uploading multiple files on the server
	// r.MaxMultipartMemory = 8 << 20
	// r.POST("/upload", func(c *gin.Context) {
	// 	form, _ := c.MultipartForm()
	// 	files := form.File["filename"]
	// 	for _, file := range files {
	// 		log.Printf("%s", file.Filename)
	// 	}
	// })

	// //router grouping
	// v1 := r.Group("/v1")
	// {
	// 	v1.GET("/hello", func(c *gin.Context) {
	// 		c.String(http.StatusOK, "hello fron v1")
	// 	})
	// }

	// v2 := r.Group("/v2")
	// {
	// 	v2.GET("/hello", func(c *gin.Context) {
	// 		c.String(http.StatusOK, "hello from v2")
	// 	})
	// }

	//middle ware for logging errors
	// r := gin.New()
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	// //above lines same as gin.Default()

	// r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// r.Run()
}
