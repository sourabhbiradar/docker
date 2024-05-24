package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		fmt.Println("hello!!!") // on console
		c.String(200, "hello")  // on browser
	})

	r.Run(":7011") // localhost:7011
}
