package limitbucket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(80, 200)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			message := message{
				Status: "Request Failed",
				Body:   "The API is at capacity, try again later.",
			}
			c.JSON(http.StatusTooManyRequests, message)
			c.Abort()
			return
		}
		c.Next()
	}
}
