// redis/redis.go
package redis_client

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(addr string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	return &RedisClient{client: client}
}

// JSON 문서로 저장하기 위한 메서드
func (rc *RedisClient) SetJSON(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	// 구조체를 JSON으로 마샬링
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// Redis JSON 명령어를 사용하여 저장
	// JSON.SET key $ <json-string>
	cmd := rc.client.Do(ctx, "JSON.SET", key, "$", string(json))
	if cmd.Err() != nil {
		return cmd.Err()
	}

	// 만료 시간 설정
	if expiration > 0 {
		rc.client.Expire(ctx, key, expiration)
	}

	return nil
}

// JSON 문서 가져오기
func (rc *RedisClient) GetJSON(ctx context.Context, key string, dest interface{}) error {
	cmd := rc.client.Do(ctx, "JSON.GET", key, "$")
	if cmd.Err() != nil {
		return cmd.Err()
	}

	jsonStr, err := cmd.Text()
	if err != nil {
		return err
	}

	// Unmarshal JSON string into a slice of interface{}
	var result []interface{}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return err
	}

	// Check if the result is not empty and assign the first element to dest
	if len(result) > 0 {
		jsonBytes, err := json.Marshal(result[0])
		if err != nil {
			return err
		}
		return json.Unmarshal(jsonBytes, dest)
	}

	return nil
}

// BTCPrice represents Bitcoin price data
type BTCPrice struct {
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

// SetBTCPrice stores BTC price data
func (rc *RedisClient) SetBTCPrice(ctx context.Context, price *BTCPrice) error {
	return rc.SetJSON(ctx, "btc:price", price, 5*time.Minute)
}

// GetBTCPrice retrieves latest BTC price
func (rc *RedisClient) GetBTCPrice(ctx context.Context) (*BTCPrice, error) {
	price := &BTCPrice{}
	err := rc.GetJSON(ctx, "btc:price", price)
	return price, err
}
