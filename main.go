package main

import "container/list"

type Cache interface {
	Get(key string) string
	Set(key, value string)
}

type LRU_Cache struct {
	m   map[string]*cacheMapElement
	cap int
	L   *list.List
}

type cacheMapElement struct { // map structure created to store a pointer and the value
	el    *list.Element
	value string
}

func NewLRU(cap int) LRU_Cache {
	return LRU_Cache{
		m:   map[string]*cacheMapElement{}, // map that stores the structures blueprint
		cap: cap,                           // capacity of cache
		L:   list.New(),
	}
}

func (c *LRU_Cache) Get(key string) string {
	v, ok := c.m[key] // retrieving the value from cache map

	if !ok { // fail-safe if the value doesn't exist
		return ""
	}

	c.L.MoveToFront(v.el) //moves the key node to the front of the list to follow caching logic
	return v.value
}

func (c *LRU_Cache) Set(key, value string) {
	v, ok := c.m[key] // attempts to retrieve value from map, checks if already exists

	if !ok { // if doesn't exist
		el := c.L.PushFront(key)     // var el becomes the new key that is pushed to the front of the list
		c.m[key] = &cacheMapElement{ // assigning the new value and pointer to the cache map
			el:    el,
			value: value,
		}

		if c.L.Len() > c.cap { // checks if cache has reached capacity and removes accordingly
			backEL := c.L.Back()                    // grabs list value that is at back of list
			backElementKey := backEL.Value.(string) // grabs the key from the node
			c.L.Remove(backEL)                      //removes the node from list
			delete(c.m, backElementKey)             // removes corresponding map item using key from node

		}

	} else { // if the key already exists, change its value and move its node to the front of the list
		v.value = value
		c.L.MoveToFront(v.el)
	}

}
