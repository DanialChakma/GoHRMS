package initializers

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

var RateLimiterMiddleware gin.HandlerFunc

func InitRateLimiter() {
	reqStr := os.Getenv("RATE_LIMIT_REQUESTS")
	durStr := os.Getenv("RATE_LIMIT_DURATION")

	if reqStr == "" {
		reqStr = "100"
	}
	if durStr == "" {
		durStr = "M" // default to minute
	}

	reqs, err := strconv.Atoi(reqStr)
	if err != nil {
		log.Fatalf("Invalid RATE_LIMIT_REQUESTS: %v", err)
	}

	rateStr := strconv.Itoa(reqs) + "-" + durStr // e.g., "100-M"
	log.Printf("Rate Limiter Str: %v\n", rateStr)
	rate, err := limiter.NewRateFromFormatted(rateStr)
	if err != nil {
		log.Fatalf("failed to create rate limit: %v", err)
	}

	store := memory.NewStore()
	instance := limiter.New(store, rate)
	//Use LimitReachedHandler for custom message
	RateLimiterMiddleware = mgin.NewMiddleware(
		instance,
		mgin.WithLimitReachedHandler(func(c *gin.Context) {
			c.JSON(429, gin.H{"error": "Too many requests, slow down!"})
			c.Abort()
		}),
	)

}
