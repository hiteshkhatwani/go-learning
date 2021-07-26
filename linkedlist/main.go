package main

import (
	"go.uber.org/zap"
	ll "linkedlist/linkedlist"
)

func main()  {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	ll.Logger = logger
	logger.Info("Initializing first linkedlist")
	l1 := &ll.Linkedlist{}
	l1.Insert(4)
	l1.Insert(43)
	l1.Insert(54)
	logger.Info("Initializing second linkedlist")
	l2 :=&ll.Linkedlist{}
	l2.Insert(3)
	l2.Insert(78)
	logger.Info("Merging both the linkedlist")
	l3 := ll.Merge(l1, l2)
	l3.Print()
	return
}