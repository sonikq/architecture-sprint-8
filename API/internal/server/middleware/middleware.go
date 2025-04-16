package middleware

import (
	"backend-api/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func HasAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const source = "middleware.HasAuth"
		rawToken := c.GetHeader("Authorization")
		if rawToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{
					models.ErrMsgKey: models.ErrNotAuthenticated.Error(),
				})
			log.Printf(models.ErrTraceLayout, source, "Missing Authorization header")
			return
		}

		s := strings.Split(rawToken, " ")
		if len(s) != 2 || s[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{
					models.ErrMsgKey: models.ErrNotAuthenticated.Error(),
				})
			log.Printf(models.ErrTraceLayout, source, models.ErrInvalidJWTToken.Error())
			return
		}
		c.Next()
	}
}
