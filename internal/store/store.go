package store

import (
	"log"
	"sync"
	"time"
)

// dataItem holds the value and its expiration timestamp.
type dataItem struct {
	value      string
	expiration int64
}

// Store is the main data structure for our key-value store.
// It is safe for concurrent use.
type Store struct {
	sync.RWMutex
	data map[string]dataItem
}

// NewStore creates a new Store, initializes it, and starts the reaper.
func NewStore() *Store {
	s := &Store{
		data: make(map[string]dataItem),
	}
	// Start the background reaper process.
	go s.reaper()
	return s
}

// Set adds a key-value pair to the store with a specific TTL.
func (s *Store) Set(key, value string, ttl int) {
	var expiration int64
	if ttl > 0 {
		expiration = time.Now().Unix() + int64(ttl)
	}

	s.Lock()
	s.data[key] = dataItem{value: value, expiration: expiration}
	s.Unlock()
}

// Get retrieves a value by its key. It returns the value and a boolean
// indicating whether the key was found (and not expired).
func (s *Store) Get(key string) (string, bool) {
	s.RLock()
	item, ok := s.data[key]
	s.RUnlock()

	if !ok {
		return "", false
	}

	// Check if the key has expired.
	if item.expiration != 0 && time.Now().Unix() > item.expiration {
		return "", false // Treat expired keys as not found.
	}

	return item.value, true
}

// reaper is the background goroutine that cleans up expired keys.
func (s *Store) reaper() {
	for {
		time.Sleep(1 * time.Second)
		now := time.Now().Unix()

		s.Lock()
		for key, item := range s.data {
			if item.expiration != 0 && now > item.expiration {
				log.Printf("Reaping key: %s (expired at %d)\n", key, item.expiration)
				delete(s.data, key)
			}
		}
		s.Unlock()
	}
}
