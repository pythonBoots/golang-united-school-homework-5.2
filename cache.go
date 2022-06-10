package cache

import "time"

type kvPair struct {
	key          string
	value        string
	deadline     time.Time
	shouldexpire bool
}

type Cache struct {
	kvPAirs []kvPair
}

func NewCache() Cache {
	return Cache{}
}

func (c Cache) Get(key string) (string, bool) {
	for _, kv := range c.kvPAirs {
		if kv.key == key {
			return kv.key, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	for _, kv := range c.kvPAirs {
		if kv.key == key {
			kv.value = value
		}
	}
	c.kvPAirs = append(c.kvPAirs, kvPair{key: key, value: value})

}

func (c Cache) Keys() []string {
	slice := []string{}
	for _, kv := range c.kvPAirs {
		if !kv.shouldexpire {
			slice = append(slice, kv.key)
		}
	}
	return slice
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	for _, kv := range c.kvPAirs {
		if kv.key == key {
			kv.value = value
			kv.deadline = deadline
			kv.shouldexpire = true
		}
	}
	c.kvPAirs = append(c.kvPAirs, kvPair{
		key:          key,
		value:        value,
		deadline:     deadline,
		shouldexpire: true,
	})

}
