package linkedlist

import (
	"fmt"
	"go.uber.org/zap"
)

var Logger *zap.Logger

type Node struct {
	val int
	next *Node
}

type Linkedlist struct {
	head *Node
	tail *Node
	length int
}

func (l *Linkedlist) Insert(value int) {
	newNode := &Node{value, nil}
	if l.head == nil {
		l.head = newNode
		l.tail = l.head
		l.length += 1
		Logger.Info("Created new head node", zap.Int("Value", value), zap.Int("length", l.length))
		return
	} else {
		l.tail.next = newNode
		l.tail = newNode
		l.length += 1
		Logger.Info("Inserted new node", zap.Int("Value", value), zap.Int("length", l.length))
		return
	}
}

func (l *Linkedlist) Print()  {
	node := l.head
	if node == nil {
		Logger.Error("Linkedlist is empty")
	}
	for node != nil {
		fmt.Println(node.val)
		node = node.next
	}
	return
}