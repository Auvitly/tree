package tree

// Tree - structured data tree with the root node
type tree struct {
	TRoot *node `json:"root"`
}

type Tree interface {
	Self() *tree
	Root() Node
	SetRoot(node Node) error
	FindByKey(key string) Node
	Separate(node *node) (Tree, error)
	SaveAsJSON(name, path string) error
}

// Node - the smallest unit of a data tree structure.
// Has a key, a set of fields, and parent and child nodes.
type node struct {
	NKey    string  `json:"key"`
	NFields Fields  `json:"fields"`
	NParent *node   `json:"-"`
	NChilds []*node `json:"childs"`
}

type Node interface {
	Self() *node
	Root() Node
	Key() string
	SetKey(key string) error
	Fields() Fields
	InheritedFields() Fields
	Parent() Node
	SetParent(Node)
	Childs() []Node
	AddChildNodes(nodes ...Node) error
	RemoveChild(node Node)
	FindingNodeByKey(key string) Node
}

// Fields - set of key-value parameters. Any data type is accepted as a value.
type Fields map[string]interface{}
