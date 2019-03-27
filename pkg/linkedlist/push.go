package linkedlist


func (ll *LinkedList) Push(item Item) {
	temp := new(Node)
	temp.val = item
	temp.next = nil
	ll.head = temp
	ll.tail = temp
	ll.size = 1
}