package main

import (
	bst "binarysearchtree/binarysearchtree"
	"fmt"
	//"go.uber.org/zap"
)

func main() {
	bst.Changeloglevel("Debug")
	fmt.Println("Creating BST")
	bstree := &bst.BST{}
	bstree.Inorder()
	// adding nodes into the tree
	bstree.Insert(3)
	bstree.Insert(1)
	bstree.Insert(5)
	bstree.Insert(7)
	// inorder traversal of bst
	bstree.Inorder()
	return
}
