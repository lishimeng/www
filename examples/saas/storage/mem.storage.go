package storage

import (
	"sync"
	"time"
)

const cleanupInterval = time.Minute * 10

// Item 封装存储的值和过期时间
type Item struct {
	Value      any
	ExpireTime time.Time
}

// MemoryStorage 实现 Storage 接口
type MemoryStorage struct {
	items map[string]*Item
	mu    sync.RWMutex
	stop  chan struct{} // 用于优雅关闭后台清理协程
}

// NewMemoryStorage 创建一个新的内存存储实例，并启动后台清理协程
func NewMemoryStorage() Storage {
	ms := &MemoryStorage{
		items: make(map[string]*Item),
		stop:  make(chan struct{}),
	}

	// 启动后台清理协程
	go ms.cleanupExpired(cleanupInterval)

	return ms
}

// Save 保存键值对，并设置 TTL
func (ms *MemoryStorage) Save(key string, value any, ttl time.Duration) error {
	expireTime := time.Now().Add(ttl)
	item := &Item{
		Value:      value,
		ExpireTime: expireTime,
	}

	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.items[key] = item

	return nil
}

// Get 获取键对应的值（若已过期则返回 nil）
func (ms *MemoryStorage) Get(key string) (any, error) {
	ms.mu.RLock()
	item, exists := ms.items[key]
	ms.mu.RUnlock()

	if !exists {
		return nil, nil
	}

	// 惰性检查：若已过期，返回 nil（但不立即删除，由后台协程清理）
	if time.Now().After(item.ExpireTime) {
		return nil, nil
	}

	return item.Value, nil
}

// Del 删除键对应的值
func (ms *MemoryStorage) Del(key string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	delete(ms.items, key)
	return nil
}

// cleanupExpired 后台协程：定期扫描并清理过期项
func (ms *MemoryStorage) cleanupExpired(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ms.mu.Lock()
			now := time.Now()
			for k, item := range ms.items {
				if now.After(item.ExpireTime) {
					delete(ms.items, k)
				}
			}
			ms.mu.Unlock()
		case <-ms.stop:
			return // 收到停止信号，退出协程
		}
	}
}

// Close 停止后台协程（可选，用于优雅关闭）
func (ms *MemoryStorage) Close() {
	close(ms.stop)
}

// 实现 Storage 接口（隐式满足）
var _ Storage = (*MemoryStorage)(nil)
