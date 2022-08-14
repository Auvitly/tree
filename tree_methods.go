package tree

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

// NewTree - creating a tree structure with an empty root node (key and value are empty)
func NewTree(root Node) Tree {
	tree := &tree{}
	if root.Self() == nil {
		tree.TRoot = &node{
			NKey:    "root",
			NFields: make(Fields),
			NParent: nil,
			NChilds: make([]*node, 0),
			NTree:   tree,
		}
	} else {
		tree.TRoot = root.Self()
		root.Self().NTree = tree
	}
	return tree
}

// Self - returns a pointer to the structure that is under the interface
func (t *tree) Self() *tree {
	return t
}

// Root - getting the root node
func (t *tree) Root() Node {
	return t.TRoot
}

// SetRoot - function sets a node from the tree as the root of the tree
func (t *tree) SetRoot(node Node) error {
	// Finding a node in a tree
	root := t.TRoot.FindingNodeByKey(node.Key())
	if root.Self() == nil {
		// If not founded
		node.AddChilds(t.Root().Childs()...)
		t.TRoot = node.Self()
		t.Root().SetParent(nil)
		return nil
	}
	// The operation to change the root of the tree
	switch {
	case root.Parent().Self() == nil:
		return nil
	case root.Parent().Self() != nil:
		var n1, n2, n3 Node
		n1 = root
		n2 = root.Parent()
		n3 = root.Parent().Parent()
		for n2.Self() != nil || n3.Self() != nil {
			n1.AddChilds(n2)
			n2.RemoveChild(n1)
			n1 = n2
			n2 = n3
			if n3 != nil {
				n3 = n3.Parent()
			}
		}
		root.SetParent(nil)
		t.TRoot = root.Self()
	}
	return nil
}

// FindByKey - search by node key in tree (depth-first search)
func (t *tree) FindByKey(key string) Node {
	return t.TRoot.FindingNodeByKey(key)
}

// FindByValue - search by node value in tree (depth-first search)
func (t *tree) FindByValue(value interface{}) []Node {
	return t.TRoot.FindingNodesByValue(value)
}

// Separate - branching a new tree from a node in the original tree
func (t *tree) Separate(node *node) (Tree, error) {
	// Finding a node in a tree
	root := t.TRoot.FindingNodeByKey(node.NKey)
	if root.Self() == nil {
		// If not founded
		return nil, errors.Errorf("node not found in current tree")
	}
	// Tree separation stage
	parent := root.Parent()
	parent.RemoveChild(root)
	tree := NewTree(nil)
	tree.SetRoot(root)
	return tree, nil
}

// JSON - transform sucture of the tree to JSON-format
func (t *tree) JSON() ([]byte, error) {
	// Marshaling
	data, err := json.Marshal(t)
	return data, err
}

// SaveAsJSON - saves the tree to a JSON file
func (t *tree) SaveAsJSON(name, path string) error {
	// Filepath definition
	fpath := getFilepath(name, path)
	// Create file
	file, err := os.OpenFile(fpath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	// Marshaling
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	// Write in file
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	// Close file
	err = file.Close()
	if err != nil {
		return err
	}
	return nil

}

// LoadTree - loading a tree from a file at the specified path.
// Returns a tree instance and an error. If the file was not found or was corrupted.
func LoadTree(name, path string) (Tree, error) {
	// Filepath definition
	filepath := getFilepath(name, path)
	// Create file
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	// Get data from file
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	// Unmarshal
	t := new(tree)
	err = json.Unmarshal(data, t)
	if err != nil {
		return nil, err
	}
	// Fix parents
	t.Root().Self().NTree = t
	t.Root().Self().loadParents(t.Root().Self().NChilds)

	return t, nil
}

func CreateTreeFromJSON(data []byte) (Tree, error) {
	// Unmarshal
	t := new(tree)
	err := json.Unmarshal(data, t)
	if err != nil {
		return nil, err
	}
	// Fix parents
	t.Root().Self().loadParents(t.Root().Self().NChilds)
	return t, nil
}
