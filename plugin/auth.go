package plugin

import (
	"github.com/gin-gonic/gin"
	"github.com/huhx-headhunter/headhunter-common/store"
	"github.com/huhx/common-go/exception"
	"github.com/samber/lo"
	"strings"
)

var whiteList = []string{
	"users/login",
	"users/register",
	"/versions/",
	"/swagger/",

	// for testing
	"/send",
	"/test",
	"/resend",
	"/trigger",
}

func AuthInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		if shouldSkipAuth(c.Request.URL.Path) {
			c.Next()
			return
		}

		username := c.Request.Header.Get("X-Request-Username")
		if username == "" {
			c.Abort()
			panic(exception.BadRequest{Content: "Username header is required"})
		}

		//originalHost := c.Request.Header.Get("X-Forwarded-Host")
		//if originalHost != "47.105.152.148:9019" {
		//	c.Abort()
		//	panic(exception.BadRequest{Content: "Host header is not correct"})
		//}

		store.Save("username", username)
		c.Next()
	}
}

func shouldSkipAuth(path string) bool {
	return lo.ContainsBy(whiteList, func(item string) bool { return strings.Contains(path, item) })
}
