package main

import (
	"fmt"
	"sync"
)

type Platform struct {
	Name string
}

type AccountPreferences struct {
	NotificationsEnabled bool
}

// type Preferences interface {
// 	NotificationsEnabled() bool
// }

// type AccountPreferences struct {}

// func (a AccountPreferences) NotificationsEnabled() bool {
// 	return true
// }

type MyCacheable interface {
	Platform | AccountPreferences // cannot use an interface type in union for the constraints
}

type MyCache[T MyCacheable] struct {
	data map[string]T
	mu   *sync.RWMutex
}

func (c *MyCache[T]) Set(key string, val T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = val
}

func (c *MyCache[T]) Get(key string) (v T) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if v, ok := c.data[key]; ok {
		return v
	}

	return
}

func New[T MyCacheable]() *MyCache[T] {
	c := MyCache[T]{
		data: make(map[string]T),
		mu:   &sync.RWMutex{},
	}

	return &c
}

// constaints with structs example
func main() {
	// platforms
	p := Platform{
		Name: "Music Studio",
	}
	pCache := New[Platform]()
	pCache.Set("music", p)
	fmt.Println(pCache.Get("music"))

	// preferences
	pref := AccountPreferences{
		NotificationsEnabled: true,
	}
	prCache := New[AccountPreferences]()
	prCache.Set("settings", pref)
	fmt.Println(prCache.Get("settings"))
}
