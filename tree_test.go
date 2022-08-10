package tree

import "testing"

func TestInheritedFields(t *testing.T) {

	tree := NewTree(NewNode("root", nil).Self())

	tree.Root().Fields()["key2"] = "value1"

	node := NewNode("node1", nil)

	tree.Root().AddChilds(node)

	node.Fields()["key1"] = "value"

	tree.SaveAsJSON("", "")

	nt, _ := LoadTree("", "")

	t.Log(nt)

}
