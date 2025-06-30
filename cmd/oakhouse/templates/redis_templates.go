// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

// RedisAdapterTemplate generates the Redis adapter file
const RedisAdapterTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"%s/config"
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
		return nil, fmt.Errorf("failed to connect to Redis: %%v", err)
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
`

// RedisUtilTemplate generates the Redis utility file
const RedisUtilTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package util

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"%s/adapter"
)

type CacheManager struct {
	redisAdapter *adapter.RedisAdapter
}

// NewCacheManager creates a new cache manager
func NewCacheManager(redisAdapter *adapter.RedisAdapter) *CacheManager {
	return &CacheManager{
		redisAdapter: redisAdapter,
	}
}

// SetCacheWithTags stores data with tags for easy invalidation
func (cm *CacheManager) SetCacheWithTags(ctx context.Context, key string, value interface{}, expiration time.Duration, tags []string) error {
	// Store the main data
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal cache data: %%v", err)
	}

	if err := cm.redisAdapter.Set(ctx, key, data, expiration); err != nil {
		return fmt.Errorf("failed to set cache: %%v", err)
	}

	// Store tag associations
	for _, tag := range tags {
		tagKey := fmt.Sprintf("tag:%%s", tag)
		if err := cm.redisAdapter.GetClient().SAdd(ctx, tagKey, key).Err(); err != nil {
			return fmt.Errorf("failed to add tag association: %%v", err)
		}
		// Set expiration for tag keys (slightly longer than cache expiration)
		cm.redisAdapter.GetClient().Expire(ctx, tagKey, expiration+time.Hour)
	}

	return nil
}

// GetCache retrieves cached data
func (cm *CacheManager) GetCache(ctx context.Context, key string, dest interface{}) error {
	data, err := cm.redisAdapter.Get(ctx, key)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), dest)
}

// InvalidateByTag removes all cache entries associated with a tag
func (cm *CacheManager) InvalidateByTag(ctx context.Context, tag string) error {
	tagKey := fmt.Sprintf("tag:%%s", tag)
	
	// Get all keys associated with this tag
	keys, err := cm.redisAdapter.GetClient().SMembers(ctx, tagKey).Result()
	if err != nil {
		return fmt.Errorf("failed to get tag members: %%v", err)
	}

	// Delete all associated keys
	if len(keys) > 0 {
		if err := cm.redisAdapter.GetClient().Del(ctx, keys...).Err(); err != nil {
			return fmt.Errorf("failed to delete cache keys: %%v", err)
		}
	}

	// Delete the tag key itself
	if err := cm.redisAdapter.Delete(ctx, tagKey); err != nil {
		return fmt.Errorf("failed to delete tag key: %%v", err)
	}

	return nil
}

// InvalidatePattern removes all cache entries matching a pattern
func (cm *CacheManager) InvalidatePattern(ctx context.Context, pattern string) error {
	keys, err := cm.redisAdapter.GetClient().Keys(ctx, pattern).Result()
	if err != nil {
		return fmt.Errorf("failed to get keys by pattern: %%v", err)
	}

	if len(keys) > 0 {
		if err := cm.redisAdapter.GetClient().Del(ctx, keys...).Err(); err != nil {
			return fmt.Errorf("failed to delete cache keys: %%v", err)
		}
	}

	return nil
}

// ClearAll removes all cache entries (use with caution)
func (cm *CacheManager) ClearAll(ctx context.Context) error {
	return cm.redisAdapter.GetClient().FlushDB(ctx).Err()
}
`
