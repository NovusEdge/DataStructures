package dstructs

import "fmt"

//LNode : Creates a *Node* Object for a ***Singly*** Linked List
type LNode struct {
	Data interface{}
	Next *LNode
}

//DNode : Creates a *Node* Object for a ***Doubly*** Linked List
type DNode struct {
	Data interface{}
	Next *DNode
	Prev *DNode
}

//LinkedList : Creates a ***Singly*** linked list
type LinkedList struct {
	Head *LNode
}

//DLinkedList : Creates a ***Doubly*** linked list
type DLinkedList struct {
	Head *DNode
}

//CLinkedList : Creates a ***Circular*** linked list
type CLinkedList struct {
	Head *LNode
}

//DCLinkedList : Creates a ***Circular*** ***Doubly*** linked list
type DCLinkedList struct {
	Head *DNode
}

//Len : Reports the length of a Linked-List
func (l *LinkedList) Len() int {
	tempNode := l.Head
	counter := 0
	for tempNode.Next != nil {
		counter++
		tempNode = tempNode.Next
	}
	return counter
}

func (dl *DLinkedList) Len() int {
	tempNode := dl.Head
	counter := 0
	for tempNode.Next != nil {
		counter++
		tempNode = tempNode.Next
	}
	return counter
}

func (cl *CLinkedList) Len() int {
	tempNode := cl.Head.Next
	start := cl.Head
	counter := 1
	for tempNode != start {
		counter++
		tempNode = tempNode.Next
	}
	return counter
}

func (dcl *DCLinkedList) Len() int {
	tempNode := dcl.Head.Next
	counter := 1
	for tempNode != dcl.Head {
		tempNode = tempNode.Next
		counter++
	}
	return counter
}

/*
MakeList : Takes Multiple nodes as args and creates a ***Singly*** Linked list out of them.

* Takes the first Node passed as the *Head Node*.
*/
func MakeList(nodes ...LNode) LinkedList {
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = &nodes[i+1]
	}
	return LinkedList{Head: &nodes[0]}
}

/*
MakeDList : Takes **Multiple** nodes as args and creates a ***Doubly*** Linked list out of them.

* Takes the first *Node* passed as the *Head Node*.
*/
func MakeDList(nodes ...DNode) DLinkedList {
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = &nodes[i+1]
		nodes[i+1].Prev = &nodes[i]
	}
	return DLinkedList{Head: &nodes[0]}
}

/*
MakeCList : Takes **Multiple** nodes as args and creates a ***Circular*** Linked list out of them.

* Takes the first *Node* passed as the *Head Node*.

NOTE: If only one node is passed then returns a circular list that has only one node, adviced to just use struct instead
*/
func MakeCList(nodes ...LNode) CLinkedList {
	if len(nodes) == 1 {
		nodes[0].Next = &nodes[0]
		nodes[0].Next.Next = &nodes[0]
		return CLinkedList{Head: &nodes[0]}
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = &nodes[i+1]
	}

	nodes[len(nodes)-1].Next = &nodes[0]
	return CLinkedList{Head: &nodes[0]}
}

/*
MakeDCList : Takes **Multiple** nodes as args and creates a ***Circular*** ***Doubly*** Linked list out of them.

* Takes the first *Node* passed as the *Head Node*.

NOTE: If only one node is passed then returns a circular list that has only one node, adviced to just use struct instead
*/
func MakeDCList(nodes ...DNode) DCLinkedList {

	if len(nodes) == 1 {
		nodes[0].Next = &nodes[0]
		nodes[0].Prev = &nodes[0]
		return DCLinkedList{Head: &nodes[0]}
	}

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = &nodes[i+1]
		nodes[i+1].Prev = &nodes[i]
	}
	nodes[len(nodes)-1].Next = &nodes[0]
	nodes[0].Prev = &nodes[len(nodes)-1]
	return DCLinkedList{Head: &nodes[0]}
}

//PrintList  : **Prints** the List
func (l *LinkedList) PrintList() {
	nodeList := []LNode{}
	tempNode := l.Head
	for tempNode != nil {
		nodeList = append(nodeList, *tempNode)
		tempNode = tempNode.Next
	}
	fmt.Println(nodeList)
}

func (dl *DLinkedList) PrintList() {
	nodeList := []DNode{}
	tempNode := dl.Head
	for tempNode != nil {
		nodeList = append(nodeList, *tempNode)
		tempNode = tempNode.Next
	}
	fmt.Println(nodeList)
}

func (cl *CLinkedList) PrintList() {
	tempNode := cl.Head.Next
	nodeList := []LNode{*cl.Head}
	for tempNode != cl.Head {
		nodeList = append(nodeList, *tempNode)
		tempNode = tempNode.Next
	}
	nodeList = append(nodeList, *tempNode)
	fmt.Println(nodeList)
}

func (dcl *DCLinkedList) PrintList() {
	nodeList := []DNode{*dcl.Head}
	tempNode := dcl.Head.Next
	for tempNode != dcl.Head {
		nodeList = append(nodeList, *tempNode)
		tempNode = tempNode.Next
	}
	nodeList = append(nodeList, *tempNode)
	fmt.Println(nodeList)
}

//Append : appends to the linked list
func (l *LinkedList) Append(node LNode) {
	tempNode := l.Head.Next
	prevNode := l.Head
	for tempNode != nil {
		tempNode = tempNode.Next
		prevNode = prevNode.Next
	}
	prevNode.Next = &node
}

func (dl *DLinkedList) Append(node DNode) {
	tempNode := dl.Head
	for tempNode.Next != nil {
		tempNode = tempNode.Next
	}
	tempNode.Next = &node
	tempNode.Next.Prev = tempNode
}

func (cl *CLinkedList) Append(node LNode) {
	tempNode := cl.Head.Next
	start := cl.Head
	for tempNode.Next != start {
		tempNode = tempNode.Next
	}
	node.Next = start
	tempNode.Next = &node
}

func (dcl *DCLinkedList) Append(node DNode) {
	node.Prev = dcl.Head.Prev // fix tis method
	node.Prev.Next = &node
	dcl.Head.Prev = &node
	node.Next = dcl.Head
}

//Search : Runs a *linear search* on the linked list to find node with Data == [Data]
func (l *LinkedList) Search(Data interface{}) *LNode {
	tempNode := l.Head
	for tempNode != nil {
		if tempNode.Data == Data {
			return tempNode
		}
		tempNode = tempNode.Next
	}
	return nil
}

func (dl *DLinkedList) Search(Data interface{}) *DNode {
	tempNode := dl.Head
	for tempNode != nil {
		if tempNode.Data == Data {
			return tempNode
		}
		tempNode = tempNode.Next
	}
	return nil
}

func (cl *CLinkedList) Search(Data interface{}) *LNode {
	tempNode := cl.Head.Next
	start := cl.Head
	for tempNode.Next != start {
		if tempNode.Data == Data {
			return tempNode
		}
		tempNode = tempNode.Next
	}
	return nil
}

func (dcl *DCLinkedList) Search(Data interface{}) *DNode {
	tempNode := dcl.Head.Next
	for tempNode != dcl.Head {
		if tempNode.Data == Data {
			return tempNode
		}
		tempNode = tempNode.Next
	}
	return nil
}

//Delete : Deletes the first node from the linked list which has [data]
func (l *LinkedList) Delete(Data interface{}) {
	if l.Head.Data == Data {
		l.Head = l.Head.Next
	} else {

		tempNode := l.Head.Next
		prevNode := l.Head
		for tempNode != nil {
			if tempNode.Data == Data {
				prevNode.Next = tempNode.Next
				break
			}
			tempNode = tempNode.Next
			prevNode = prevNode.Next
		}
	}
}

func (dl *DLinkedList) Delete(Data interface{}) {
	node := dl.Search(Data)
	tempNode := node
	node.Next.Prev = tempNode.Prev
	node.Prev.Next = tempNode.Next
	node.Next = nil
	node.Prev = nil
}

func (cl *CLinkedList) Delete(data interface{}) {
	if cl.Head.Data != data {
		cl._delNode(data)
	} else {
		cl._deleteHead()
	}
}

func (cl *CLinkedList) _deleteHead() {
	if cl.Len() > 1 {
		tempNode := cl.Head
		ref := cl.Head.Next

		for tempNode.Next != cl.Head {
			tempNode = tempNode.Next
		}
		fmt.Println(tempNode)
		tempNode.Next = ref
		cl.Head = tempNode.Next
	}
}

func (cl *CLinkedList) _delNode(data interface{}) {
	tempNode := cl.Head.Next
	prevNode := cl.Head

	for tempNode.Next != cl.Head {
		if tempNode.Data == data {
			prevNode.Next = tempNode.Next
			tempNode.Next = nil
			break
		}
		tempNode = tempNode.Next
		prevNode = prevNode.Next
	}
}

func (dcl *DCLinkedList) Delete(data interface{}) {
	if dcl.Search(data) != nil {
		if data == dcl.Head.Data {
			dcl.Head.Next.Prev = dcl.Head.Prev
			dcl.Head.Prev.Next = dcl.Head.Next
			dcl.Head = dcl.Head.Next
		} else {
			tempNode := dcl.Head
			for tempNode.Data != data {
				tempNode = tempNode.Next
			}
			tempNode.Prev.Next = tempNode.Next
			tempNode.Next.Prev = tempNode.Prev
			tempNode.Next = nil
			tempNode.Prev = nil
		}
	}
}
