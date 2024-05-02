package pokecache

import (
    "sync"
    "time"
)

//Cached Entry response value and time
type cacheEntry struct {
    createdAt time.Time
    val []byte
}

//All cached values from responses
type Cache struct {
    cachedValues map[string]cacheEntry
    mu  *sync.RWMutex
}

//Interface for interacting with cache
//Allows adding responses, retrieving responses
//Will reap cached values older than 5 seconds
type storage interface {
    Add(key string, val []byte)
    Get(key string) (val []byte, found bool)
    reapLoop(interval time.Duration)
}

//Generates a new cache map called _cache_storage
//Makes a new async function to reap the cache of old reponses
//Returns the cache map and error of creating reap function
func NewCache(interval time.Duration) (Cache, chan bool) {
    _cache_storage := Cache{
        cachedValues: make(map[string]cacheEntry),
        mu: &sync.RWMutex{},
    }
    quitChan := make(chan bool)
    _cache_storage.reapLoop(5, quitChan)
    return _cache_storage, quitChan
}

//Adds a response to the cache
//Requires the html reponse as a array of bytes, key as the API request
//Returns nothing
func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.cachedValues[key] = cacheEntry{time.Now(), val}
}

//Returns a response from the cache map if it exists
//Returns an error if it is not contained
//Requires the string API request as the key
//Returns the reponse if found and true
//Returns no string and false if not found
func (c *Cache) Get(key string) (val []byte, found bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()

    ce, found := c.cachedValues[key]
    return ce.val, found
}

//Asyncronous loop responsible for pruning the cache map for old values
//On every ticker tick, compares the time to the maximum interval
//If the interval is greater, than the value is removed
func (c *Cache) reapLoop(interval time.Duration, quit chan bool) {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    go func() {
        for {
            select {
            case <- quit:
                return
            case <-ticker.C:
                c.mu.Lock()
                for val := range(c.cachedValues) {
                    now := time.Now()
                    if  now.Sub(c.cachedValues[val].createdAt) > interval {
                        delete(c.cachedValues, val)
                    }
                }
                c.mu.Unlock()
            }
        }
    }()
}

