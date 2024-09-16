package main

import ll "github.com/aslotsu/algorithmic-thinking/list-linked"

func main() {
	ln3 := ll.ListNode{Data: 4}
	ln2 := ll.ListNode{Data: 3, Next: &ln3}
	ln1 := ll.ListNode{Data: 2, Next: &ln2}

	myLL := ll.LinkedList{Head: &ln1, Size: 0}

	myLL.InsertAtBeginning(13)

	

	myLL.InsertEnd(12)
	myLL.Display()
	myLL.InsertAt(4, 30)
	myLL.Display()

}
