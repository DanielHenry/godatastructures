package linkedlist


func New() *LinkedList {
	ll := new(LinkedList)
	ll.head = nil
	ll.tail = nil
	ll.size = 0
	return ll
}