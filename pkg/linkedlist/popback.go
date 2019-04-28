package linkedlist


func (ll *LinkedList) PopBack(item Item) {
	if ll.size == 0 {
		ll.head = nil
		ll.tail =  nil
		ll.size = 0
	} else if ll.size == 1 {
		ll.head = nil
		ll.tail =  nil
		ll.size = 0
	} else {
		for iter := ll.head; iter.next != ll.tail; iter = iter.next {
			continue
		}
		ll.tail = iter
		ll.size--
	}
}