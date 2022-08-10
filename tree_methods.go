package tree

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/pkg/errors"
)

// NewTree - creating a tree structure with an empty root node (key and value are empty)
func NewTree(root *node) Tree {
	tree := &tree{}
	if root == nil {
		tree.TRoot = &node{
			NKey:    "root",
			NFields: make(Fields),
			NParent: nil,
			NChilds: make([]*node, 0),
		}
	} else {
		tree.TRoot = root
	}
	return tree
}

// Self
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
		node.AddChildNodes(t.Root().Childs()...)
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
			n1.AddChildNodes(n2)
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

// SaveAsJSON - saves the tree to a JSON file
func (t *tree) SaveAsJSON(name, path string) error {
	// Filepath definition
	fpath := getFilepath(name, path)
	// Create file
	file, err := os.OpenFile(fpath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
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

// LoadTree - Loading a tree from a file at the specified path.
// Returns a tree instance and an error. If the file was not found or was corrupted.
func LoadTree(name, path string) (Tree, error) {
	// Filename definition
	var filename string
	switch {
	case len(name) == 0:
		filename = fmt.Sprintf("tree_%s.json", time.Now().String())
	default:
		filename = fmt.Sprintf("%s.json", name)
	}
	// Path definition
	var filepath string
	switch {
	case len(path) == 0:
		filepath = filename
	default:
		filepath = path + filename
	}
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
	t.Root().Self().fixParent(t.Root().Self().NChilds)

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
	t.Root().Self().fixParent(t.Root().Self().NChilds)
	return t, nil
}