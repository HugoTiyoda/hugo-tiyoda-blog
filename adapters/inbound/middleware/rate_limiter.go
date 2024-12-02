package middleware

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type IPRateLimiter struct {
	sync.RWMutex
	limiters map[string]*RateLimiter
	rate     float64
	capacity int
}

type RateLimiter struct {
	rate     float64
	capacity int
	tokens   float64
	lastTime time.Time
	mu       sync.Mutex
}

func NewIPRateLimiter(rate float64, capacity int) *IPRateLimiter {
	return &IPRateLimiter{
		limiters: make(map[string]*RateLimiter),
		rate:     rate,
		capacity: capacity,
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastTime).Seconds()
	rl.tokens += elapsed * rl.rate

	if rl.tokens > float64(rl.capacity) {
		rl.tokens = float64(rl.capacity)
	}

	if rl.tokens < 1 {
		return false
	}

	rl.tokens--
	rl.lastTime = now
	return true
}

func (i *IPRateLimiter) GetLimiter(ip string) *RateLimiter {
	i.Lock()
	defer i.Unlock()

	limiter, exists := i.limiters[ip]
	if !exists {
		limiter = &RateLimiter{
			rate:     i.rate,
			capacity: i.capacity,
			tokens:   float64(i.capacity),
			lastTime: time.Now(),
		}
		i.limiters[ip] = limiter
	}

	return limiter
}

func RateLimitMiddleware(rate float64, capacity int) gin.HandlerFunc {
	ipLimiter := NewIPRateLimiter(rate, capacity)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := ipLimiter.GetLimiter(ip)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(429, gin.H{
				"error": "rate limit exceeded",
			})
			return
		}

		c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", capacity))
		remaining := int(limiter.tokens)
		c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))

		c.Next()
	}
}
