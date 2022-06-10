package cache

import "time"

type kvP struct {
	key          string
	value        string
	deadline     time.Time
	shouldexpire bool
}

type Cache struct {
	kvPs []kvP
}

func NewCache() Cache {
	return Cache{}
}

func (c Cache) Get(key string) (string, bool) {
	for _, kv := range c.kvPs {
		if kv.key == key {
			return kv.key, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	for _, kv := range c.kvPs {
		if kv.key == key {
			kv.value = value
		}
	}
	c.kvPs = append(c.kvPs, kvP{key: key, value: value})

}

func (c Cache) Keys() []string {
	slice := []string{}
	for _, kv := range c.kvPs {
		if !kv.shouldexpire {
			slice = append(slice, kv.key)
		}
	}
	return slice
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	for _, kv := range c.kvPs {
		if kv.key == key {
			kv.value = value
			kv.deadline = deadline
			kv.shouldexpire = true
		}
	}
	c.kvPs = append(c.kvPs, kvP{
		key:          key,
		value:        value,
		deadline:     deadline,
		shouldexpire: true,
	})

}
