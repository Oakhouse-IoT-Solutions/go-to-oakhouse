// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"oakhouse-release-latest/config"
	"github.com/redis/go-redis/v9"
)

type RedisAdapter struct {
	client *redis.Client
}

// NewRedisAdapter creates a new Redis adapter
func NewRedisAdapter(cfg *config.Config) (*RedisAdapter, error) {
	// Parse Redis DB number
	db, err := strconv.Atoi(cfg.RedisDB)
	if err != nil {
		db = 0 // default to DB 0
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL,
		Password: cfg.RedisPassword,
		DB:       db,
	})

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return &RedisAdapter{client: rdb}, nil
}

// GetClient returns the Redis client
func (r *RedisAdapter) GetClient() *redis.Client {
	return r.client
}

// Set stores a key-value pair with expiration
func (r *RedisAdapter) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value by key
func (r *RedisAdapter) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Delete removes a key
func (r *RedisAdapter) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// Exists checks if a key exists
func (r *RedisAdapter) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.client.Exists(ctx, key).Result()
	return result > 0, err
}

// SetJSON stores a JSON object
func (r *RedisAdapter) SetJSON(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.JSONSet(ctx, key, "$", value).Err()
}

// GetJSON retrieves a JSON object
func (r *RedisAdapter) GetJSON(ctx context.Context, key string, dest interface{}) error {
	cmd := r.client.Do(ctx, "JSON.GET", key, "$")
	if cmd.Err() != nil {
		return cmd.Err()
	}
	val, err := cmd.Text()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}
	
// Close closes the Redis connection
func (r *RedisAdapter) Close() error {
	return r.client.Close()
}
