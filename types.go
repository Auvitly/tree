package tree

// Tree - structured data tree with the root node
type tree struct {
	TRoot *node `json:"root"`
}

type Tree interface {
	// Self - returns a pointer to the structure that is under the interface
	Self() *tree
	// Root - getting the root node
	Root() Node
	// SetRoot - function sets a node from the tree as the root of the tree
	SetRoot(node Node) error
	// FindByKey - search by node key in tree (depth-first search)
	FindByKey(key string) Node
	// FindByValue - search by node value in tree (depth-first search)
	FindByValue(value interface{}) []Node
	// Separate - branching a new tree from a node in the original tree
	Separate(node *node) (Tree, error)
	// JSON - marshaling to JSON
	JSON() ([]byte, error)
}

// Node - the smallest unit of a data tree structure.
// Has a key, a set of fields, and parent and child nodes.
type node struct {
	NKey    string  `json:"key"`
	NFields Fields  `json:"fields"`
	NParent *node   `json:"-"`
	NChilds []*node `json:"childs"`
	NTree   *tree   `json:"-"`
}

type Node interface {
	// Self - returns a pointer to a structure, for deeper data manipulation
	Self() *node
	// Tree - returns the tree.
	Tree() Tree
	// Root - returns root of the tree.
	Root() Node
	// Key - returns key of current node.
	Key() string
	// SetKey - sets the passed value as the key. Returns an error if the key already exists in the tree.
	SetKey(key string) error
	// Fields - return fields of current node. Fields is map[string]interface{}
	Fields() Fields
	// InheritedFields - allows you to get fields that are inherited from parent nodes.
	// If the field of the child node has the same key as the field of the parent node,
	// then the priority is given to the child node.
	InheritedFields() Fields
	// Parent - returns parent current node.
	Parent() Node
	// SetParent - sets the passed node as the parent node.
	// The correctness of the tree is maintained by the appropriate reassignment of parent and child nodes
	SetParent(Node)
	// Childs - returns child nodes for current node.
	Childs() []Node
	// AddChilds - adding child nodes for the current node.
	// Child nodes are stripped of their current parent nodes.
	// The current node is set as the parent node.
	AddChilds(nodes ...Node) error
	// RemoveChild - deleting a child node for the current node
	RemoveChild(node Node)
	// FindingNodeByKey - search for a node in the tree relative to the current node
	FindingNodeByKey(key string) Node
}

// Fields - set of key-value parameters. Any data type is accepted as a value.
type Fields map[string]interface{}
