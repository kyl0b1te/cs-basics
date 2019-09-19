package linearsearch

type Comparable interface {
	compare(value interface{}) bool
}

func LinearSearch(nodes []Comparable, search Comparable) int {

	for i, node := range nodes {
		if ok := node.compare(search); ok == true {
			return i
		}
	}
	return -1
}
