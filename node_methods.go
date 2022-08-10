package tree

import (
	"github.com/pkg/errors"
)

// NewNode - returns a pointer to the node structure
func NewNode(key string, fields Fields) Node {
	if len(key) == 0 {
		return nil
	}
	node := &node{
		NKey:    key,
		NChilds: make([]*node, 0),
	}
	if fields == nil {
		node.NFields = make(Fields)
	} else {
		node.NFields = fields
	}
	return node
}

// Key - getting the key of the current node
func (n *node) Key() string {
	return n.NKey
}

// SetKey - setting the key for the current node. Returns an error if the key already exists.
func (n *node) SetKey(key string) error {
	if node := n.Root().FindingNodeByKey(key); node == nil {
		n.NKey = key
		return nil
	} else {
		return ErrorAlreadyExist
	}
}

// Fields - getting fields.
func (n *node) Fields() Fields {
	return n.NFields
}

// InheritedFields - allows you to get fields that are inherited from parent nodes.
// If the field of the child node has the same key as the field of the parent node,
// then the priority is given to the child node.
func (n *node) InheritedFields() Fields {
	// make fields
	fields := make(Fields)
	// work with nodes
	parent := n.Parent()
	nodes := make([]Node, 0)
	nodes = append(nodes, n)
	for parent.Self() != nil {
		nodes = append(nodes, parent)
		parent = parent.Parent()
	}
	for i := len(nodes) - 1; i >= 0; i-- {
		for key, value := range nodes[i].Fields() {
			fields[key] = value
		}
	}
	return fields
}

// Parent -
func (n *node) Parent() Node {
	return n.NParent
}

// SetParent -
func (n *node) SetParent(node Node) {
	if node == nil {
		n.NParent = nil
	} else {
		// TODO
	}
}

// Self -
func (n *node) Self() *node {
	return n
}

// Childs -
func (n *node) Childs() []Node {
	var childs []Node
	for _, child := range n.NChilds {
		childs = append(childs, child)
	}
	return childs
}

// AddChildNodes - adding child nodes for the current node.
// Child nodes are stripped of their current parent nodes.
// The current node is set as the parent node.
func (n *node) AddChildNodes(nodes ...Node) error {
	for index, node := range nodes {
		if result := n.Root().FindingNodeByKey(node.Key()); result != nil {
			return errors.Errorf("Node key - arg[%d] already exists", index)
		}
		if node.Parent().Self() != nil {
			node.Parent().RemoveChild(node)
		}
		node.SetParent(n)
		n.NChilds = append(n.NChilds, node.Self())
	}
	return nil
}

// Root - finding the root of a tree relative to the current node
func (n *node) Root() (root Node) {
	var p Node = n
	for p != nil {
		if p.Parent().Self() == nil {
			return p
		}
		p = p.Parent()
	}
	return n
}

// FindingNodeByKey - search for a node in the tree relative to the current node
func (n *node) FindingNodeByKey(key string) (node Node) {
	switch {
	case n.NKey == key:
		return n
	case len(n.NChilds) != 0:
		for _, child := range n.NChilds {
			node = child.FindingNodeByKey(key)
			if node != nil {
				return
			}
		}
		return nil
	default:
		return nil
	}
}

// FindingNodesByValue - search for a node relative to the current node that
// has a value passed as a function argument
func (n *node) FindingNodesByValue(value interface{}) (result []Node) {
	// make a slice
	result = make([]Node, 0)
	// search in the node fields
	key := n.valueSearch(value)
	if key != nil {
		result = append(result, n)
	}
	// search in childs
	for _, child := range n.NChilds {
		nodes := child.FindingNodesByValue(value)
		result = append(result, nodes...)
	}
	return
}

// RemoveChild - deleting a child node for the current node
func (n *node) RemoveChild(node Node) {
	switch {
	case len(n.NChilds) > 0:
		for index, child := range n.NChilds {
			if node.Self() == child {
				child.NParent = nil
				n.NChilds = append(n.NChilds[:index], n.NChilds[index+1:]...)
			}
		}
	}
}
