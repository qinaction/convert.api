package cors

import (
	"convert.api/libs/common/error_wrapper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() error_wrapper.HandlerFunc {
	return func(c *gin.Context) error {
		method := c.Request.Method
		//origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Accept, Authorization,Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Set("content-type", "application/json")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()

		return nil
	}
}
