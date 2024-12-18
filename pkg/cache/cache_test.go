package cache

import (
	"testing"
)

func TestNewInMemoryCache(t *testing.T) {
	cache := NewInMemoryCache()
	if cache == nil {
		t.Error("Expected a new InMemoryCache instance, got nil")
	}
	if len(cache.data) != 0 {
		t.Errorf("Expected empty data map, got %v", cache.data)
	}
}

func TestInMemoryCache_Get(t *testing.T) {
	cache := NewInMemoryCache()

	// Test retrieving non-existent key
	value, exists := cache.Get("missing")
	if exists {
		t.Error("Expected key to not exist")
	}
	if value != "" {
		t.Errorf("Expected value to be an empty string, got %v", value)
	}

	// Test retrieving existing key
	cache.Set("key1", "value1")
	value, exists = cache.Get("key1")
	if !exists {
		t.Error("Expected key to exist")
	}
	if value != "value1" {
		t.Errorf("Expected value 'value1', got %v", value)
	}
}

func TestInMemoryCache_Set(t *testing.T) {
	cache := NewInMemoryCache()

	// Test setting a new key-value pair
	cache.Set("key1", "value1")
	value, exists := cache.Get("key1")
	if !exists {
		t.Error("Expected key to exist after setting")
	}
	if value != "value1" {
		t.Errorf("Expected value 'value1', got %v", value)
	}

	// Test updating an existing key
	cache.Set("key1", "newValue")
	value, exists = cache.Get("key1")
	if !exists {
		t.Error("Expected key to exist after updating")
	}
	if value != "newValue" {
		t.Errorf("Expected value 'newValue', got %v", value)
	}
}

func TestInMemoryCache_Flush(t *testing.T) {
	cache := NewInMemoryCache()

	// Test flushing an empty cache
	cache.Flush()
	if len(cache.data) != 0 {
		t.Errorf("Expected empty cache, got %v", cache.data)
	}

	// Test flushing a populated cache
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Flush()
	if len(cache.data) != 0 {
		t.Errorf("Expected empty cache after flush, got %v", cache.data)
	}
	if _, exists := cache.Get("key1"); exists {
		t.Error("Expected key1 to be removed after flush")
	}
	if _, exists := cache.Get("key2"); exists {
		t.Error("Expected key2 to be removed after flush")
	}
}
