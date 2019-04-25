package linkedlist


func (ll *LinkedList) Push(item Item) {
	temp := new(Node)
	temp.val = item
	temp.next = nil
	if ll.head == nil && ll.tail == nil {
		ll.head = temp
		ll.tail = temp
		ll.size = 1
	} else {
		ll.tail.next = temp
		ll.tail = temp
		ll.size++
	}
}