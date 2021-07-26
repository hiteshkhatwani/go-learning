package linkedlist

import "go.uber.org/zap"

func Merge(l1, l2 *Linkedlist) ( *Linkedlist){
	l3 := &Linkedlist{}
	ptr1 := l1.head
	ptr2 := l2.head

	if ptr1 == nil && ptr2 == nil {
		Logger.Error("Both linkedlist are empty")
		return nil
	}

	for ptr1 != nil {
		l3.Insert(ptr1.val)
		Logger.Info("Adding node from first list", zap.Int("Value", ptr1.val))
		ptr1 = ptr1.next
	}
	for ptr2 != nil {
		l3.Insert(ptr2.val)
		Logger.Info("Adding node from second list", zap.Int("Value", ptr2.val))
		ptr2 = ptr2.next
	}
	return l3
}
