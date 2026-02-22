package pokecache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_AddCache(t *testing.T) {
	// Arrange
	cache := NewCache(5 * time.Second)

	// Act
	cache.Add("key1", []byte("value1"))
	cache.Add("key2", []byte("value2"))

	// Assert
	assert.Len(t, cache.Data, 2)
}

func Test_GetCache(t *testing.T) {
	// Arrange
	cache := NewCache(5 * time.Second)
	cache.Add("key1", []byte("value1"))

	// Act
	val, ok := cache.Get("key1")
	_, missing := cache.Get("missing")

	// Assert
	assert.True(t, ok)
	assert.Equal(t, []byte("value1"), val)
	assert.False(t, missing)
}

func Test_ReapLoop(t *testing.T) {
	// Arrange
	interval := 50 * time.Millisecond
	cache := NewCache(interval)
	cache.Add("old", []byte("data"))

	// Act
	time.Sleep(3 * interval)

	// Assert
	_, ok := cache.Get("old")
	assert.False(t, ok, "expected old entry to be reaped")

	cache.Add("new", []byte("data"))
	_, ok = cache.Get("new")
	assert.True(t, ok, "expected new entry to exist")
}
