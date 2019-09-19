package linearsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestNode struct {
	value int
}

func (n TestNode) compare(with interface{}) bool {
	node := with.(TestNode)
	return n.value == node.value
}

func getTestData() []Comparable {
	tests := []Comparable{}
	for i := 0; i <= 5; i++ {
		var node Comparable = TestNode{i}
		tests = append(tests, node)
	}
	return tests
}

func TestLinearSearch(t *testing.T) {

	tests := getTestData()
	for i, test := range tests {
		assert.Equal(t, i, LinearSearch(tests, test))
	}

	assert.Equal(t, -1, LinearSearch(tests, TestNode{99}))
}
