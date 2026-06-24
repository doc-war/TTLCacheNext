// cache.go
package cache

import (
	"time"
	"github.com/doc-war/lru-next/internal" // 引入内部包
)

// Cache 是面向最终用户的强类型安全缓存客户端
type Cache[T any] struct {
	sharded *internal.ShardedCache[T]
}

// New 创建一个支持高并发分片、带 TTL 和 LRU 淘汰的缓存实例
func New[T any](maxKeys int, ttl time.Duration) (*Cache[T], error) {
	sc, err := internal.NewShardedCache[T](maxKeys, ttl)
	if err != nil {
		return nil, err
	}
	return &Cache[T]{sharded: sc}, nil
}

// GetOrLoad 从缓存中获取数据。如果过期或不存在，则通过 loader 函数自动回源加载。
func (c *Cache[T]) GetOrLoad(id string, loader func(string) (T, error)) (T, error) {
	return c.sharded.GetOrLoad(id, loader)
}

// Get 只读缓存。不存在或已过期时返回 zero value + false
func (c *Cache[T]) Get(id string) (T, bool) {
	return c.sharded.Get(id)
}

// Set 直接写入缓存，覆盖已有值，重置 TTL
func (c *Cache[T]) Set(id string, val T) {
	c.sharded.Set(id, val)
}

// Delete 主动删除指定 key 的缓存
func (c *Cache[T]) Delete(id string) {
	c.sharded.Delete(id)
}

// Clear 清空所有缓存分片
func (c *Cache[T]) Clear() {
	c.sharded.Clear()
}

// Len 返回当前缓存的近似总数
func (c *Cache[T]) Len() int {
	return c.sharded.Len()
}