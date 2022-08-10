package tree

// valueSearch - searching for a value inside fields
func (n *node) valueSearch(value interface{}) []string {
	var result = make([]string, 0)
	for key, field := range n.Fields() {
		if field == value {
			result = append(result, key)
		}
	}
	if len(result) != 0 {
		return result
	}
	return nil
}

// loadParents
func (n *node) loadParents(childs []*node) {
	for _, child := range childs {
		child.NTree = n.NTree
		child.NParent = n
		child.loadParents(child.NChilds)
	}
}
