package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

// Initialize the in-memory cache
var c = cache.New(5*time.Minute, 10*time.Minute)

// Model struct containing a Korean string and a timestamp
type Model struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// SetupRouter sets up the gin router with the endpoints
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Create endpoint
	r.POST("/model", func(ctx *gin.Context) {
		var model Model
		if err := ctx.ShouldBindJSON(&model); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		model.Timestamp = time.Now()
		c.Set("model", model, cache.DefaultExpiration)
		ctx.JSON(http.StatusOK, gin.H{"message": "Model created", "data": model})
	})

	// Read endpoint
	r.GET("/model", func(ctx *gin.Context) {
		if value, found := c.Get("model"); found {
			ctx.JSON(http.StatusOK, gin.H{"data": value})
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Model not found"})
	})

	return r
}

func StartApi() {
	r := SetupRouter()
	// Start the server
	r.Run(":8080")
}
