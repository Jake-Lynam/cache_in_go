package main

import "testing"

func TestCache(t *testing.T) {
	cache := NewLRU(1)
	cache.Set("UK", "London")
	u := cache.Get("UK")
	if u != "London" {
		t.Fatalf("expected output: `London`, returned: `%s`", u)
	}

	cache.Set("UK", "Manchester")
	u = cache.Get("UK")
	if u != "Manchester" {
		t.Fatalf("Expected output `manchester`, actual output: `%s", u)
	}
}

func TestCache2(t *testing.T) {
	cache := NewLRU(3)
	cache.Set("UK", "London")
	cache.Set("Germany", "Berlin")
	cache.Set("France", "Paris")

	//order will be France, Germany, UK

	cache.Set("Ireland", "Cork")

	u := cache.Get("UK")
	if u != "" {
		t.Fatalf("UK/London should have been removed")

	}

	//order now Ireland, France, Germany

	r := cache.Get("Germany")
	if r != "Berlin" {
		t.Fatalf("expected output Berlin, got `%s`", r)
	}

	// order now Germany, Ireland France

	c := cache.Get("UK")
	if c != "" {
		t.Fatalf("Failed: UK shouldn't be present, `%s", c)
	}

	n := cache.Get("Germany")
	if n != "Berlin" {
		t.Fatalf("Failed: germany isn't present, `%s", n)
	}

	m := cache.Get("Ireland")
	if m != "Cork" {
		t.Fatalf("Failed: Ireland isn't present, `%s", m)
	}

	x := cache.Get("France")
	if x != "Paris" {
		t.Fatalf("Failed: France isn't present, `%s", x)
	}

	println("cache result: ", n, m, x)
}
