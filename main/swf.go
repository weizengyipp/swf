package main

import (
	"net/http"

	"github.com/weizengyipp/swf/engine"
)

func main() {
	eng := engine.Default()
	eng.GET("/", indexHandler)
	v1 := eng.Group("/v1")
	v1.GET("/hello", helloHandlerByQuery)

	v2 := eng.Group("/v2")
	v2.GET("/hello/:name", helloHandlerByParam)
	v2.POST("/login", loginHandler)
	eng.Run(":8080")
}

func indexHandler(c *engine.Context) {
	c.HTML(http.StatusOK, "<h1>Hello World!</h1>")
}

func helloHandlerByQuery(c *engine.Context) {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}

func helloHandlerByParam(c *engine.Context) {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
}

func loginHandler(c *engine.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}
