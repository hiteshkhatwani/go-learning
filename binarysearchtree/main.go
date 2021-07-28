package main

import (
	bst "binarysearchtree/binarysearchtree"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	bst.Logger = logger
	fmt.Println("Creating BST")
	bstree := &bst.BST{}
	// adding nodes into the tree
	bstree.Insert(3)
	bstree.Insert(1)
	bstree.Insert(5)
	bstree.Insert(7)
	// inorder traversal of bst
	bstree.Inorder()
	return
}
