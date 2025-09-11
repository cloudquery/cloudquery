package cloudapi

import (
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/rs/zerolog"
)

// CacheItem represents a cached item with expiration
type CacheItem struct {
	Data      any
	ExpiresAt time.Time
}

// Cache provides a thread-safe LRU cache with TTL support
type Cache struct {
	cache *lru.Cache[string, CacheItem]
	mutex sync.RWMutex
}

// NewCache creates a new cache with the specified capacity
func NewCache(capacity int) (*Cache, error) {
	lruCache, err := lru.New[string, CacheItem](capacity)
	if err != nil {
		return nil, err
	}
	return &Cache{cache: lruCache}, nil
}

// Get retrieves an item from the cache if it exists and hasn't expired
func (c *Cache) Get(key string) (any, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, exists := c.cache.Get(key)
	if !exists {
		return nil, false
	}

	if time.Now().After(item.ExpiresAt) {
		c.mutex.RUnlock()
		c.mutex.Lock()
		c.cache.Remove(key)
		c.mutex.Unlock()
		c.mutex.RLock()
		return nil, false
	}

	return item.Data, true
}

// Set stores an item in the cache with the specified TTL
func (c *Cache) Set(key string, data any, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache.Add(key, CacheItem{
		Data:      data,
		ExpiresAt: time.Now().Add(ttl),
	})
}

// Remove removes an item from the cache
func (c *Cache) Remove(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache.Remove(key)
}

// Clear removes all items from the cache
func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache.Purge()
}

// Len returns the number of items in the cache
func (c *Cache) Len() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.cache.Len()
}

// Keys returns all keys in the cache
func (c *Cache) Keys() []string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.cache.Keys()
}

// MemoryStorage provides thread-safe storage for conversations and token usage
type MemoryStorage struct {
	conversations map[string]string
	tokenUsage    map[string]int
	mutex         sync.RWMutex
}

// NewMemoryStorage creates a new memory storage instance
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		conversations: make(map[string]string),
		tokenUsage:    make(map[string]int),
	}
}

// GetConversationID retrieves a conversation ID for a user
func (m *MemoryStorage) GetConversationID(userID string) (string, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	conversationID, exists := m.conversations[userID]
	return conversationID, exists
}

// SetConversationID stores a conversation ID for a user
func (m *MemoryStorage) SetConversationID(userID, conversationID string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.conversations[userID] = conversationID
}

// GetTokenUsage retrieves token usage for a team
func (m *MemoryStorage) GetTokenUsage(teamName string) (int, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	tokenUsage, exists := m.tokenUsage[teamName]
	return tokenUsage, exists
}

// SetTokenUsage stores token usage for a team
func (m *MemoryStorage) SetTokenUsage(teamName string, tokenUsage int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.tokenUsage[teamName] = tokenUsage
}

// CacheManager manages both API response cache and memory storage
type CacheManager struct {
	cache         *Cache
	memoryStorage *MemoryStorage
	logger        zerolog.Logger
}

// NewCacheManager creates a new cache manager
func NewCacheManager(logger zerolog.Logger) (*CacheManager, error) {
	cache, err := NewCache(1000)
	if err != nil {
		return nil, err
	}

	return &CacheManager{
		cache:         cache,
		memoryStorage: NewMemoryStorage(),
		logger:        logger,
	}, nil
}

// Get retrieves an item from the API cache
func (cm *CacheManager) Get(key string) (any, bool) {
	return cm.cache.Get(key)
}

// Set stores an item in the API cache with 1-hour TTL
func (cm *CacheManager) Set(key string, data any) {
	cm.cache.Set(key, data, time.Hour)
}

// GetConversationID retrieves a conversation ID for a user
func (cm *CacheManager) GetConversationID(userID string) (string, bool) {
	return cm.memoryStorage.GetConversationID(userID)
}

// SetConversationID stores a conversation ID for a user
func (cm *CacheManager) SetConversationID(userID, conversationID string) {
	cm.memoryStorage.SetConversationID(userID, conversationID)
}

// GetTokenUsage retrieves token usage for a team
func (cm *CacheManager) GetTokenUsage(teamName string) (int, bool) {
	return cm.memoryStorage.GetTokenUsage(teamName)
}

// SetTokenUsage stores token usage for a team
func (cm *CacheManager) SetTokenUsage(teamName string, tokenUsage int) {
	cm.memoryStorage.SetTokenUsage(teamName, tokenUsage)
}
