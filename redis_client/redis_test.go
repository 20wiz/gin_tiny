package redis_client

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func TestRedisVersion(t *testing.T) {
	client := NewRedisClient("localhost:6379")
	ctx := context.Background()

	// Get Redis INFO
	info, err := client.client.Info(ctx).Result()
	assert.NoError(t, err)

	// Parse version from INFO command output
	for _, line := range strings.Split(info, "\n") {
		if strings.HasPrefix(line, "redis_version:") {
			version := strings.TrimPrefix(line, "redis_version:")
			t.Logf("Redis Version: %s", version)
			assert.NotEmpty(t, version)
			return
		}
	}
	t.Fatal("Redis version not found in INFO output")
}

func TestRedisClient(t *testing.T) {
	// Setup
	client := NewRedisClient("localhost:6379")
	ctx := context.Background()

	t.Run("Test Set and Get JSON", func(t *testing.T) {
		// Test data
		testData := TestData{
			Message:   "test message",
			Timestamp: time.Now(),
		}

		// Test SetJSON
		err := client.SetJSON(ctx, "test:key", testData, 1*time.Minute)
		assert.NoError(t, err)

		// Test GetJSON
		var retrieved TestData
		err = client.GetJSON(ctx, "test:key", &retrieved)
		assert.NoError(t, err)
		assert.Equal(t, testData.Message, retrieved.Message)
	})

	t.Run("Test Non-existent Key", func(t *testing.T) {
		var empty TestData
		err := client.GetJSON(ctx, "non:existent", &empty)
		assert.Error(t, err)
	})

	// Cleanup
	client.client.FlushDB(ctx)
}
