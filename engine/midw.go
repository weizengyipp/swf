package engine

import (
	"log"
	"net/http"
	"time"
)

func Logger() HandlerFunc {
	log.Printf("Log handler called")
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func InteralServerError() HandlerFunc {
	log.Printf("InteralServerError handler called")
	return func(c *Context) {
		c.Fail(http.StatusInternalServerError, "Internal Server Error")
	}
}
