package binarysearchtree

import (
	"fmt"
	"go.uber.org/zap"
)

var Logger *zap.Logger

type Node struct {
	left, right *Node
	val int
}

type BST struct {
	root *Node
}

func (bst *BST) Insert(val int) {
	if bst.root == nil {
		bst.root = &Node{val: val}
		Logger.Info("Created Root node", zap.Int("Root value", val))
	} else {
		Logger.Info("Calling the helper func to insert the new node")
		bst.root.insert(val)
	}
	return
}

func (n *Node) insert(val int) {
	if n == nil {
		Logger.Error("Should not have reached here.")
		return
	}
	if val <= n.val {
		if n.left == nil {
			n.left = &Node{val: val}
			Logger.Info("Created new node", zap.Int("Parent value", n.val), zap.Int("Child value", val))
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{val: val}
			Logger.Info("Created new node", zap.Int("Parent value", n.val), zap.Int("Child value", val))
		} else {
			n.right.insert(val)
		}
	}
	return
}

func (bst *BST) Inorder()  {
	if bst.root == nil {
		Logger.Error("Inorder traversal called for an empty tree")
	} else {
		Logger.Info("Helper inorder traversal called")
		bst.root.inorder()
	}
	return
}

func (n *Node) inorder() {
	if n == nil {
		return
	}
	n.left.inorder()
	fmt.Println(n.val)
	n.right.inorder()
	return
}
