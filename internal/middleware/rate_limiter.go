package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter(c *gin.Context, limiter *rate.Limiter) {

	if !limiter.Allow() {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests (DDoS detected)"})
		c.Abort()
		return
	}

	c.Next()
}

type IpLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.Mutex
}

func NewIPLimiter() *IpLimiter {
	return &IpLimiter{
		limiters: make(map[string]*rate.Limiter),
	}
}

func (i *IpLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.limiters[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 5)
		i.limiters[ip] = limiter
	}
	return limiter
}

func RateLimitMiddleware(i *IpLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := i.GetLimiter(ip)

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
