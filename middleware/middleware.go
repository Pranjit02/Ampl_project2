package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("Authorization")
		if apiKey != "Bearer mysecrettoken" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func RateLimiterMiddleware() gin.HandlerFunc {
	store := memory.NewStore()
	rate, _ := limiter.NewRateFromFormatted("5-M")
	limiterInstance := limiter.New(store, rate)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		res, _ := limiterInstance.Get(c, ip)
		if res.Reached {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
