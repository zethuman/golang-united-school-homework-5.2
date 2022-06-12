package cache

import (
	"time"
)

type Cache struct {
	db map[string]struct {
		value    string
		deadline time.Time
	}
}

func NewCache() Cache {
	return Cache{db: make(map[string]struct {
		value    string
		deadline time.Time
	})}
}

func (receiver Cache) Get(key string) (string, bool) {
	entry, ok := receiver.db[key]
	if !entry.deadline.IsZero() && entry.deadline.Before(time.Now()) {
		ok = false
		entry.value = ""
	}
	return entry.value, ok
}

func (receiver *Cache) Put(key, value string) {
	receiver.db[key] = struct {
		value    string
		deadline time.Time
	}{value: value}
}

func (receiver Cache) Keys() []string {
	keys := make([]string, 0)

	for k := range receiver.db {
		if !receiver.db[k].deadline.IsZero() && receiver.db[k].deadline.Before(time.Now()) {
			continue
		}
		keys = append(keys, k)
	}

	return keys
}

func (receiver *Cache) PutTill(key, value string, deadline time.Time) {
	receiver.db[key] = struct {
		value    string
		deadline time.Time
	}{value: value, deadline: deadline}
}
