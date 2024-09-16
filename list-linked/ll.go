package ll

import "fmt"

type LinkedList struct {
	Head *ListNode
	Size int
}

type ListNode struct {
	Data any
	Next *ListNode
}

func (ll *LinkedList) Display() error {
	current := ll.Head

	if ll.Head == nil {
		return fmt.Errorf("nothing to display")
	}
	for current != nil {
		fmt.Printf("%v -->  ", current.Data)
		current = current.Next
	}
	fmt.Printf("nil\n")
	return nil

}

func (ll *LinkedList) Length() int {
	current := ll.Head
	ll.Size = 0
	
	for current != nil {
	
		current = current.Next
		ll.Size++

	}
	return ll.Size
}

func (ll *LinkedList) InsertAtBeginning(data any) error {
	current := ll.Head

	newHead := &ListNode{Data: data, Next: current} 
	ll.Head = newHead
	ll.Size++
	return nil

}


func (ll *LinkedList) InsertEnd(data any) error {
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	newEnd := ListNode{Data: data, Next: nil}
	current.Next = &newEnd
	return nil
}

func (ll *LinkedList) InsertAt (position int, data int) error {
	current, pos := ll.Head, 1
	if position < 1 || position > ll.Length() {
		return fmt.Errorf("invalid insertion position")
	}

	for pos < position-1 {
		current = current.Next
		pos++
	}

	fmt.Println("correct", current.Data.(int))
	
	newNode := ListNode{Data: data, Next: current.Next}
	current.Next = &newNode
	ll.Size++
	return nil
}


func (ll *LinkedList) DeleteFirstNode() error {
	if ll.Head == nil {
		return fmt.Errorf("linked list is empty")
	}

	current := ll.Head
	current = current.Next
	ll.Head = current
	ll.Size--

	return nil
}


func (ll *LinkedList) DeleteLastNode() error {
	current := ll.Head
	prev := current
	if ll.Head == nil {
		return fmt.Errorf("list is empty already")
	}

	for current.Next != nil {
		prev = current
		current = current.Next
	}
	prev.Next = nil
	ll.Size--
	return nil
}

func (ll *LinkedList) DeleteAt(position int) error {
	if position<1 || position > ll.Length() {
		return fmt.Errorf("invalid position passed")
	}

	current, pos := ll.Head, 1
	prev := current

	for pos<position {
		pos++
		prev = current
		current = current.Next
	}
	prev.Next = current.Next
	return nil
}
