package tree

import "testing"

func TestInheritedFields(t *testing.T) {

	n1, _ := NewNode("root", nil)

	tree := NewTree(n1)

	tree.Root().Fields()["key2"] = "value1"

	node, _ := NewNode("node1", nil)

	tree.Root().AddChilds(node)

	node.Fields()["key1"] = "value"

	nt, _ := LoadTree("root", "")

	t.Log(nt)

}
