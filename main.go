package main

import "container/list"

type Cache interface {
	Get(key string) string
	Set(key, value string)
}

type LRU_Cache struct {
	m   map[string]*list.Element
	cap int
	L   list.List
}

func NewLRU(cap int) LRU_Cache {
	return LRU_Cache{
		m:   map[string]*list.Element{},
		cap: cap,
		L:   list.List{},
	}
}

func (L *LRU_Cache) Get(key string) string {
	return L.m[key]
}

func (L *LRU_Cache) Set(key, value string) {
	L.m[key] = value
}
