<h1> <img src="https://img.icons8.com/external-flatart-icons-lineal-color-flatarticons/344/external-tree-autumn-flatart-icons-lineal-color-flatarticons-3.png" width="50px" align="center"> Tree</h1>

<hr>

This package allows you to create information **trees** based on **nodes**. An example of a tree is shown in the figure below.

<img src="./img/tree_01.png" width="400px">

<hr>

<h2> 1. Nodes </h2>

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
