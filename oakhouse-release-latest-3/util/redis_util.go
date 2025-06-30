// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package util

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"oakhouse-release-latest-3/adapter"
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
		return fmt.Errorf("failed to marshal cache data: %v", err)
	}

	if err := cm.redisAdapter.Set(ctx, key, data, expiration); err != nil {
		return fmt.Errorf("failed to set cache: %v", err)
	}

	// Store tag associations
	for _, tag := range tags {
		tagKey := fmt.Sprintf("tag:%s", tag)
		if err := cm.redisAdapter.GetClient().SAdd(ctx, tagKey, key).Err(); err != nil {
			return fmt.Errorf("failed to add tag association: %v", err)
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
	tagKey := fmt.Sprintf("tag:%s", tag)
	
	// Get all keys associated with this tag
	keys, err := cm.redisAdapter.GetClient().SMembers(ctx, tagKey).Result()
	if err != nil {
		return fmt.Errorf("failed to get tag members: %v", err)
	}

	// Delete all associated keys
	if len(keys) > 0 {
		if err := cm.redisAdapter.GetClient().Del(ctx, keys...).Err(); err != nil {
			return fmt.Errorf("failed to delete cache keys: %v", err)
		}
	}

	// Delete the tag key itself
	if err := cm.redisAdapter.Delete(ctx, tagKey); err != nil {
		return fmt.Errorf("failed to delete tag key: %v", err)
	}

	return nil
}

// InvalidatePattern removes all cache entries matching a pattern
func (cm *CacheManager) InvalidatePattern(ctx context.Context, pattern string) error {
	keys, err := cm.redisAdapter.GetClient().Keys(ctx, pattern).Result()
	if err != nil {
		return fmt.Errorf("failed to get keys by pattern: %v", err)
	}

	if len(keys) > 0 {
		if err := cm.redisAdapter.GetClient().Del(ctx, keys...).Err(); err != nil {
			return fmt.Errorf("failed to delete cache keys: %v", err)
		}
	}

	return nil
}

// ClearAll removes all cache entries (use with caution)
func (cm *CacheManager) ClearAll(ctx context.Context) error {
	return cm.redisAdapter.GetClient().FlushDB(ctx).Err()
}
