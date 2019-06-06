package hashmap

import "github.com/zhikiri/cs-basics/LinkedList"

// HashMap represents a corresponding abstract data structure
type HashMap struct {
	size    int
	buckets []*linkedlist.LinkedList
}

// NewHashMap is a function that creates new "instance" for structure
func NewHashMap(size int) *HashMap {
	hashMap := &HashMap{size: size}
	for i := 0; i < size; i++ {
		list := linkedlist.NewList(linkedlist.DefaultCompare)
		hashMap.buckets = append(hashMap.buckets, &list)
	}
	return hashMap
}

// Set is a function that insert new key value pair in to the hash map
func (hashMap *HashMap) Set(key string, value int) *linkedlist.Node {
	bucket := hashMap.getBucket(key)
	return bucket.Append(newHashMapNode(key, value))
}

// Get is a function that retrieve element from hash map by key
func (hashMap *HashMap) Get(key string) int {
	bucket := hashMap.getBucket(key)
	node := bucket.Map(func(node *linkedlist.Node) bool {
		return convertNode(node).Key == key
	})
	return convertNode(node).Value
}

func (hashMap *HashMap) getBucket(key string) *linkedlist.LinkedList {
	hash := getHash(key)
	return hashMap.buckets[hash%hashMap.size]
}

func convertNode(node *linkedlist.Node) *HashMapNode {
	return node.Value.(*HashMapNode)
}

func getHash(key string) int {
	var hash int
	for i, ch := range key {
		hash += i + int(ch)
	}
	return hash
}
