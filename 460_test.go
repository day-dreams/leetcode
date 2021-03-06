package leetcode

import "testing"

func TestLFU1(t *testing.T) {
	cache := LFUConstructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	t.Log(cache.Get(1))
	cache.Put(3, 3)
	t.Log(cache.Get(2))
	t.Log(cache.Get(3))
	cache.Put(4, 4)
	t.Log(cache.Get(1))
	t.Log(cache.Get(3))
	t.Log(cache.Get(4))
}

func TestLFU2(t *testing.T) {
	cache := LFUConstructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	t.Log(cache.Get(2))
	cache.Put(1, 1)
	cache.Put(4, 1)
	t.Log(cache.Get(2))
}

func TestLFU3(t *testing.T) {
	cache := LFUConstructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	t.Log(cache.Get(1))
	t.Log(cache.Get(2))
}
