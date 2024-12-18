package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewIPRateLimiter(t *testing.T) {
	rate := 1.0
	capacity := 5

	limiter := NewIPRateLimiter(rate, capacity)

	assert.NotNil(t, limiter)
	assert.Equal(t, rate, limiter.rate)
	assert.Equal(t, capacity, limiter.capacity)
	assert.NotNil(t, limiter.limiters)
}

func TestGetLimiter(t *testing.T) {
	limiter := NewIPRateLimiter(1.0, 5)

	ip1 := "192.168.1.1"
	ip2 := "192.168.1.2"


	rateLimiter1 := limiter.GetLimiter(ip1)
	assert.NotNil(t, rateLimiter1)

	rateLimiter1Again := limiter.GetLimiter(ip1)
	assert.Equal(t, rateLimiter1, rateLimiter1Again)


	rateLimiter2 := limiter.GetLimiter(ip2)
	assert.NotNil(t, rateLimiter2)
	assert.NotEqual(t, rateLimiter1, rateLimiter2)
}

func TestRateLimiterAllow(t *testing.T) {
	limiter := &RateLimiter{
		rate:     1.0,
		capacity: 3,
		tokens:   3.0,
		lastTime: time.Now(),
	}

	for i := 0; i < 3; i++ {
		assert.True(t, limiter.Allow())
	}

	assert.False(t, limiter.Allow())


	time.Sleep(2 * time.Second)

	assert.True(t, limiter.Allow())
}

func TestRateLimitMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(RateLimitMiddleware(1.0, 2))

	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "192.168.1.1:12345"

	
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "2", w.Header().Get("X-RateLimit-Limit"))
	assert.Contains(t, w.Header().Get("X-RateLimit-Remaining"), "1")


	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "2", w.Header().Get("X-RateLimit-Limit"))
	assert.Contains(t, w.Header().Get("X-RateLimit-Remaining"), "0")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusTooManyRequests, w.Code)
}

func TestRateLimiterConcurrency(t *testing.T) {
	limiter := &RateLimiter{
		rate:     1.0,
		capacity: 5,
		tokens:   5.0,
		lastTime: time.Now(),
	}

	done := make(chan bool)
	allowCount := 0

	for i := 0; i < 10; i++ {
		go func() {
			if limiter.Allow() {
				allowCount++
			}
			done <- true
		}()
	}

	
	for i := 0; i < 10; i++ {
		<-done
	}

	
	assert.Equal(t, 5, allowCount)
}

func TestRateLimiterTokenReplenishment(t *testing.T) {
	limiter := &RateLimiter{
		rate:     2.0, 
		capacity: 4,
		tokens:   0.0,
		lastTime: time.Now(),
	}

	
	assert.False(t, limiter.Allow())

	
	time.Sleep(2 * time.Second)

	
	for i := 0; i < 4; i++ {
		assert.True(t, limiter.Allow())
	}


	assert.False(t, limiter.Allow())
}
