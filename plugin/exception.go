package plugin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huhx/common-go/exception"
	"net/http"
)

func ExceptionInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(exception.Exception); ok {
					fmt.Println("Recovered from panic:", err.Message())
					c.JSON(err.Code(), gin.H{"code": err.Code(), "message": err.Message()})
				} else {
					fmt.Println("Recovered from panic with unknown error", err.Message())
					serverError := http.StatusInternalServerError
					c.JSON(serverError, gin.H{"code": serverError, "message": "Unknown error"})
				}
			}
		}()
		c.Next()
	}
}
