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

func getSortedTestData() []Comparable {
	tests := []Comparable{}
	for i := 0; i <= 5; i++ {
		var node Comparable = TestNode{i}
		tests = append(tests, node)
	}
	return tests
}

func getUnsortedTestData() []Comparable {
	tests := []Comparable{}
	for _, value := range []int{34, 1, 64, 12, 97} {
		var node Comparable = TestNode{value}
		tests = append(tests, node)
	}
	return tests
}

func TestLinearSearch(t *testing.T) {

	tests := getSortedTestData()
	for i, test := range tests {
		assert.Equal(t, i, LinearSearch(tests, test))
	}
	assert.Equal(t, -1, LinearSearch(tests, TestNode{99}))

	tests = getUnsortedTestData()
	for i, test := range tests {
		assert.Equal(t, i, LinearSearch(tests, test))
	}
	assert.Equal(t, -1, LinearSearch(tests, TestNode{99}))
}
