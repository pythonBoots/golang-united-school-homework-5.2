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
			return kv.value, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	for i := 0; i < len(c.kvPs); i++ {
		if c.kvPs[i].key == key {
			c.kvPs[i].value = value
			c.kvPs[i].shouldexpire = false
			return
		}
	}
	c.kvPs = append(c.kvPs, kvP{key: key, value: value, shouldexpire: false})

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
	for i := 0; i < len(c.kvPs); i++ {
		if c.kvPs[i].key == key {
			c.kvPs[i].value = value
			c.kvPs[i].deadline = deadline
			c.kvPs[i].shouldexpire = true
			return
		}
	}
	c.kvPs = append(c.kvPs, kvP{
		key:          key,
		value:        value,
		deadline:     deadline,
		shouldexpire: true,
	})

}
