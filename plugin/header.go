package plugin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HeaderInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		headers := c.Request.Header
		method := c.Request.Method
		fmt.Printf("Request Path: %s %s\n", method, path)
		fmt.Println("Request Headers:")
		for key, values := range headers {
			for _, value := range values {
				fmt.Printf("%s: %s\n", key, value)
			}
		}
		c.Next()
	}
}
