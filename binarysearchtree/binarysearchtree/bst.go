package binarysearchtree

import (
	//"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var encoderCfg = zap.NewProductionEncoderConfig()
var atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
var Logger = zap.New(zapcore.NewCore(
	zapcore.NewJSONEncoder(encoderCfg),
	zapcore.Lock(os.Stdout),
	atomicLevel,
))


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
		Logger.Debug("Created Root node", zap.Int("Root value", val))
	} else {
		Logger.Debug("Calling the helper func to insert the new node")
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
			Logger.Debug("Created new node", zap.Int("Parent value", n.val), zap.Int("Child value", val))
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{val: val}
			Logger.Debug("Created new node", zap.Int("Parent value", n.val), zap.Int("Child value", val))
		} else {
			n.right.insert(val)
		}
	}
	return
}

func (bst *BST) Inorder()  {
	if bst.root == nil {
		Logger.Warn("Inorder traversal called for an empty tree")
	} else {
		Logger.Debug("Helper inorder traversal called")
		bst.root.inorder()
	}
	return
}

func (n *Node) inorder() {
	if n == nil {
		return
	}
	n.left.inorder()
	Logger.Info("Node", zap.Int("Value", n.val))
	n.right.inorder()
	return
}

func Changeloglevel(newlevel string) {
	switch newlevel {
	case "Debug":
		atomicLevel.SetLevel(zapcore.DebugLevel)
	case "Warn":
		atomicLevel.SetLevel(zapcore.WarnLevel)
	default:
		atomicLevel.SetLevel(zapcore.InfoLevel)
	}
	return
}