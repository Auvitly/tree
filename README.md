<h1> <img src="https://img.icons8.com/external-flatart-icons-lineal-color-flatarticons/344/external-tree-autumn-flatart-icons-lineal-color-flatarticons-3.png" width="50px" align="center"> Tree</h1>

This package allows you to create information **trees** based on **nodes**. An example of a tree is shown in the figure below.

<img src="./img/tree_01.png" width="400px">

<h2> 1. Tree </h2>

A tree is a structure based on nodes and has a root node. For a tree, there are a number of functions provided in Tree interface.

````go
type Tree interface {
	// Self - returns a pointer to a structure, for deeper data manipulation
	Self() *tree
	// Root - returns the root node
	Root() Node
	// SetRoot - Allows you to set the node as the root node. If the node is part of a tree, 
	//then the nodes are restructured to form a valid tree. If the node is not part of the tree, 
	// then the root node will be added to the childs, and the added node will become the root.
	SetRoot(node Node) error
	// FindByKey - searches for a key within a tree
	FindByKey(key string) Node
	// FindByValue - searches by field value within a tree 
	FindByValue(value interface{}) []Node
	// Separate - allows you to split a node from a tree into a new tree. All links are removed.
	Separate(node *node) (Tree, error)
	// SaveAsJSON - save to json file
	SaveAsJSON(name, path string) error
}
````

<h2> 2. Nodes </h2>

**A node** is the main structural element of **a tree**. 

Each **node** has: 
* key - for a tree is a **unique identifier** that allows you to quickly find a node;
* fields - is a **map[string]interface{}**. Stores data by keys;
* pointer to tree - allows you to quickly access the tree;
* pointer to parent;
* pointers to children.

To create a node, use the function *NewNode*. The first argument is the key, the second is the fields. It is allowed to create a node without specifying a second argument, however, if an empty key is specified, it will return **nil**.
````go
  ...
  node := tree.NewNode("key",nil)
  ...
````

