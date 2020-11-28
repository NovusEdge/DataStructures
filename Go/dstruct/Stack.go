package dstructs

type sItem struct {
	value interface{}
	next  *sItem
}

//Stack :To initialize a "Stack" datastructure
type Stack struct {
	top  *sItem
	size int
}

//Len : Reports the size of a stack
func (stack *Stack) Len() int {
	return stack.size
}

//Push : Pushes a value into the Stack
func (stack *Stack) Push(value interface{}) {
	stack.top = &sItem{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

//Pop : Pops the top value of the stack
func (stack *Stack) Pop() (value interface{}) {
	if stack.IsEmpty() {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}
	return nil
}

//IsEmpty reports if the stack is empty
func (stack *Stack) IsEmpty() bool {
	return stack.Len() == 0
}
