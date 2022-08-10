package tree

import "testing"

func TestInheritedFields(t *testing.T) {

	tree := NewTree(NewNode("root", nil).Self())

	tree.Root().Fields()["key2"] = "value1"

	node := NewNode("node1", nil)

	tree.Root().AddChildNodes(node)

	node.Fields()["key1"] = "value"

	t.Log(node.InheritedFields())

	tree.SaveAsJSON("", "")

}
