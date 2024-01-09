package main

import (
	"net/http"

	"github.com/weizengyipp/swf/engine"
)

func main() {
	engine := engine.New()
	engine.Get("/", indexHandler)
	engine.Get("/hello", helloHandler)
	engine.Post("/login", loginHandler)
	engine.Run(":8080")
}

func indexHandler(c *engine.Context) {
	c.HTML(http.StatusOK, "<h1>Hello World!</h1>")
}

// handler echoes r.URL.Header
func helloHandler(c *engine.Context) {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}

func loginHandler(c *engine.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}
