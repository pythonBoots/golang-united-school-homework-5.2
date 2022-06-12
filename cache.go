package cache

import "time"

type kvP struct {
	key             string
	value           string
	deadline        time.Time
	canNotBeExpired bool
}

type Cache struct {
	kvPs []kvP
}

func NewCache() Cache {
	return Cache{}
}

func (c Cache) Get(key string) (string, bool) {
	for _, kv := range c.kvPs {
		if kv.key == key && (!time.Now().After(kv.deadline) || kv.canNotBeExpired) {
			return kv.value, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	for i := 0; i < len(c.kvPs); i++ {
		if c.kvPs[i].key == key {
			c.kvPs[i].value = value
			c.kvPs[i].canNotBeExpired = true
			return
		}
	}
	c.kvPs = append(c.kvPs, kvP{key: key, value: value, canNotBeExpired: true})

}

func (c Cache) Keys() []string {
	slice := []string{}
	for _, kv := range c.kvPs {
		if !time.Now().After(kv.deadline) || kv.deadline.IsZero() {
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
			return
		}
	}
	c.kvPs = append(c.kvPs, kvP{
		key:      key,
		value:    value,
		deadline: deadline,
	})

}
