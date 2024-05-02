package pokecache

import (
	"reflect"
	"sync"
	"testing"
    "fmt"
    "time"
)

func TestCreateCache(t *testing.T) {
	_test_cache, err := NewCache(7)
	want := Cache{
		cachedValues: make(map[string]cacheEntry),
		mu:           &sync.RWMutex{},
	}
	if !(reflect.DeepEqual(want, _test_cache) || err != nil) {
		t.Fatalf(`NewCache(7) = %v, %v, want match for %v, nil`, _test_cache, err, want)
	}
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "htps://example.com/path",
			val: []byte("moretestdata"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(T *testing.T) {
			cache,_ := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}
