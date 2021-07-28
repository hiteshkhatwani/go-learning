package main

import (
	"fmt"
	ll "linkedlist/linkedlist"
)

func main()  {
	ll.Changeloglevel("Info")
	l1 := &ll.Linkedlist{}
	l1.Insert(2)
	l1.Insert(3)
	l1.Insert(54)
	fmt.Println("Initializing second linkedlist")
	l2 :=&ll.Linkedlist{}
	l1.Merge(l2)
	l2.Insert(3)
	l2.Insert(78)
	fmt.Println("Merging both the linkedlist")
	l1.Merge(l2)
	l1.Print()
	return
}

