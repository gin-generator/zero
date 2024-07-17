package middlewares

import (
	http "github.com/gin-generator/zero/package/respond"
	"github.com/gin-gonic/gin"
)

// ForceUA Force the request header to include a User-Agent.
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.UserAgent() == "" {
			http.Alert400WithoutMessage(c, http.MissUserAgent)
		}
		c.Next()
	}
}
