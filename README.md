# LRU - Least Recently Used Cache (Go)

Caching system that uses a map and linked-list to record data. This design was bulit to be later implemented into a larger project (edge-caching proxy)

# Status

Completed as a standalone project, to be implemented within larger project later

# Why

I wanted to gain the knowledge and skill of how to design a caching system within Go without using a pre-bulit script. The caching script will also be used within a
larger project that invloves caching request types and responses within a proxy. 

# How it works

Two data structures work together:

A map is used to withhold the value (city) and pointer to the node held within the linked-list.
A linked-list is used to withhold the most recently used keys (countries) which are nodes, allowing for the logic of caching. 

The storage of pointers within the map allows for accessed keys to be moved to the front of the list in 0(1) time. Each list node contains the key so once a node is deleted from the
list, the corresponding map item can be deleted along side it.

Get(key):
  - look up key in the map
  - if found: move its list node to the front, return the value
  - if not found: return empty

Set(key, value):
  - if key exists: update its value, move its list node to the front
  - if key is new: create a list node and map item
  - if now over capacity: remove the node at the back of the
    list and delete its key from the map


# Design notes

For the implementation within the proxy, the use of mutex will be needed to prevent error. For this specifc project the focus was just on creating and understanding how to bulid
a LRU cache in Go.

The keys stored within the nodes instead of the map allow for proper cache logic, when a node is removed from the list it's key can identify the exact map item to delete. 
If we had included the values instead of the keys within the list, once we delete the map item there would be no suitable way of knowing what it's value was. 

# Built while learning from:

Go's container/list docs: https://pkg.go.dev/container/list
LLD: Implementing LRU Cache with Go (Medium): https://medium.com/code-simplified/lld-lru-cache-implementation-in-go-f8b20268fa3b
youtube guide: https://www.youtube.com/watch?v=DvmMYD8oaZw&t=367s 