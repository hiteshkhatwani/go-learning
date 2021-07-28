package linkedlist

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var encoderCfg = zap.NewProductionEncoderConfig()
var atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
var logger = zap.New(zapcore.NewCore(
	zapcore.NewJSONEncoder(encoderCfg),
	zapcore.Lock(os.Stdout),
	atomicLevel,
))
var sugar = logger.Sugar()

type Node struct {
	val interface{}
	next *Node
}

type Linkedlist struct {
	head *Node
	tail *Node
	length int
}

func (l *Linkedlist) Insert(value interface{}) {
	newNode := &Node{value, nil}
	if l.head == nil {
		l.head = newNode
		l.tail = l.head
		l.length += 1
		sugar.Debugf("Created new head node with value: %v", value)
		return
	} else {
		l.tail.next = newNode
		l.tail = newNode
		l.length += 1
		sugar.Debugf("Inserted new node with value: %v", value)
		return
	}
}

func (l *Linkedlist) Print()  {
	node := l.head
	if node == nil {
		sugar.Warnf("Trying to print an empty linkedlist.")
	}
	for node != nil {
		sugar.Infof("%v", node.val)
		node = node.next
	}
	return
}

func (l *Linkedlist) Merge(l2 * Linkedlist) {
	switch {
	case l.head == nil && l2.head == nil:
		sugar.Warnf("Trying to merge to two empty linkedlist.")
		return
	case l.head == nil:
		sugar.Warnf("Trying to merge into an empty linkedlist.")
		return
	case l2.head == nil:
		sugar.Warnf("Trying to merge an empty linkedlist.")
		return
	}
	l.tail.next = l2.head
	l.tail = l2.tail
	l.length += l2.length
	return
}


func Changeloglevel(newlevel string){
	switch newlevel {
	case "Debug":
		atomicLevel.SetLevel(zap.DebugLevel)
	case "Warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	default:
		atomicLevel.SetLevel(zap.InfoLevel)
	}
}
