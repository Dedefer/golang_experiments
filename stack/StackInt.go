package stack

//Stack of int in dynamic memory
type Stack struct {
	head *node
}

type node struct {
	value int
	next  *node
}

//Push int value to stack
func (obj *Stack) Push(val int) {
	tempHead := obj.head
	obj.head = new(node)
	obj.head.next = tempHead
	obj.head.value = val
}

//Pop value from stack. Returns value and true if it exist, else returns false.
func (obj *Stack) Pop() (int, bool) {
	var value int
	if obj.head != nil {
		value = obj.head.value
		obj.head = obj.head.next
		return value, true
	}
	return value, false
}

//Peek value from stack. Returns value and true if it exist, else returns false.
func (obj Stack) Peek() (int, bool) {
	if obj.head != nil {
		return obj.head.value, true
	}
	var value int
	return value, false
}

//Clear stack (obvivious)
func (obj *Stack) Clear() {
	obj.head = nil
}

//Copy aplies deep copy of stack
func (obj Stack) Copy() (copy Stack) {
	if obj.head != nil {
		copy.head = new(node)
		copy.head.value = obj.head.value
		copy.head.next = obj.head.next
		for tempNode := copy.head; tempNode != nil; tempNode = tempNode.next {
			temp := tempNode.next
			if temp != nil {
				tempNode.next = new(node)
				tempNode.next.value = temp.value
				tempNode.next.next = temp.next
			}
		}
	}
	return
}

//Empty chek that stack is empty
func (obj Stack) Empty() bool {
	if obj.head != nil {
		return false
	}
	return true
}
