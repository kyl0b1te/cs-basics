package hashmap

import (
	"testing"

	"github.com/zhikiri/data-structures/LinkedList"
)

func TestNewHashMapWithZeroLen(t *testing.T) {
	hm := NewHashMap(0)
	if length := len(hm.buckets); length != 0 {
		t.Errorf("Expected '%+v' actual '%+v'", 0, length)
	}
}

func TestNewHashMapWithNonZeroLen(t *testing.T) {
	hm := NewHashMap(5)
	if length := len(hm.buckets); length != 5 {
		t.Errorf("Expected '%+v' actual '%+v'", 5, length)
	}
}

func TestGetHashOnEmptyString(t *testing.T) {
	if hash := getHash(""); hash != 0 {
		t.Errorf("Expected '%+v' actual '%+v'", 0, hash)
	}
}

func TestGetHash(t *testing.T) {
	tests := map[string]int{
		"table": 530,
		"pool":  448,
		"genre": 539,
	}
	for word, expected := range tests {
		if actual := getHash(word); actual != expected {
			t.Errorf("Expected '%+v' actual '%+v'", expected, actual)
		}
	}
}

func TestSet(t *testing.T) {
	hm := NewHashMap(5)
	tests := []string{"table", "pool", "genre"}

	nodes := make(map[int]*linkedlist.Node)
	for i, word := range tests {
		index := getHash(word) % hm.size
		nodes[index] = hm.Set(word, i)
	}

	for _, word := range tests {
		index := getHash(word) % hm.size
		if hm.buckets[index].Head != nodes[index] {
			t.Errorf("Expected '%+v' actual '%+v'", nodes[index], hm.buckets[index].Head)
		}
	}
}

func TestGet(t *testing.T) {
	hm := NewHashMap(5)
	tests := map[string]int{"k01": 10, "k02": 20, "k03": 30}

	for key, val := range tests {
		hm.Set(key, val)
	}

	for key, expected := range tests {
		actual := hm.Get(key)
		if expected != actual {
			t.Errorf("Expected '%+v' actual '%+v'", expected, actual)
		}
	}
}
