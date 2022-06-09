package cache

import "time"

type Cache struct {
	key          []string
	value        []string
	deadline     []time.Time
	shouldexpire []bool
}

func NewCache(key []string, value []string, deadline []time.Time) Cache {
	return Cache{key: key, value: value, deadline: deadline}
}

func (c Cache) Get(key string) (string, bool) {
	for n := 0; n < len(c.key); n++ {
		if c.key[n] == key {
			return c.value[n], true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	for n := 0; n < len(c.key); n++ {
		if c.key[n] == key {
			c.value[n] = value
		}
	}
	c.key = append(c.key, key)
	c.value = append(c.value, value)

}

func (c Cache) Keys() []string {
	slice := []string{}
	for n := 0; n < len(c.key); n++ {
		if !c.shouldexpire[n] {
			slice = append(slice, c.key[n])
		}
	}
	return slice
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	for n := 0; n < len(c.key); n++ {
		if c.key[n] == key {
			c.value[n] = value
			c.deadline[n] = deadline
			c.shouldexpire[n] = true
		}
	}
	c.key = append(c.key, key)
	c.value = append(c.value, value)
	c.shouldexpire = append(c.shouldexpire, true)
	c.deadline = append(c.deadline, deadline)
}
