package hashmap

type HashMapNode struct {
	Key   string
	Value int
}

func newHashMapNode(key string, value int) *HashMapNode {
	return &HashMapNode{Key: key, Value: value}
}
