package dstructs

import "fmt"

var searchString = ""

//Node : Creates a Node for a binary tree
type Node struct {
	Data  int
	Left  *BST
	Right *BST
}

//BNode : Creates a Node for a balanced binary tree
type BNode struct {
	Data   int
	Left   *BBST
	Right  *BBST
	Parent *BNode
}

//BBST : Creates a balanced binary tree
type BBST struct {
	Root *BNode
}

//BST creates a Binary tree
type BST struct {
	Root *Node
}

/*
MakeTree : Takes Multiple nodes as args and creates a ***Binary Search Tree*** out of them.

* Takes the first Node passed as the *Root Node*.
*/
func MakeTree(nodes ...Node) BST {
	treeObj := BST{Root: &nodes[0]}
	for i := 1; i < len(nodes); i++ {
		treeObj.Insert(&nodes[i])
	}
	return treeObj
}

/*
MakeBTree : Takes Multiple nodes as args and creates a ***Balanced Binary Search Tree*** out of them.

* Takes the first Node passed as the *Root Node*.
*/
func MakeBTree(nodes ...BNode) BBST {
	treeObj := BBST{Root: &nodes[0]}
	for i := 1; i < len(nodes); i++ {
		treeObj.Insert(&nodes[i])
	}
	return treeObj
}

//Search finds node with matching data in a Binary Search Tree
func (bt *BST) Search(data int) *Node {
	if bt.Root == nil {
		return bt.Root
	}
	if data < bt.Root.Data {
		if bt.Root.Left != nil {
			searchString += "L"
			return bt.Root.Left.Search(data)
		}
		searchString += "R"
		return bt.Root.Right.Search(data)

	} else if data > bt.Root.Data {
		if bt.Root.Right != nil {
			searchString += "R"
			return bt.Root.Right.Search(data)
		}
		searchString += "L"
		return bt.Root.Left.Search(data)
	}
	fmt.Printf("Path: %s\n", searchString)
	searchString = ""
	return bt.Root
}

// ----------------------------------------------------------------------------------------------------------------

//! needs work
func (bbt *BBST) Search(data int) *BNode {
	if bbt.Root == nil {
		return bbt.Root
	}

	if data < bbt.Root.Data {
		if bbt.Root.Left != nil {
			searchString += "L"
			return bbt.Root.Left.Search(data)
		}
		searchString += "R"
		return bbt.Root.Right.Search(data)

	} else if data > bbt.Root.Data {
		if bbt.Root.Right != nil {
			searchString += "R"
			return bbt.Root.Right.Search(data)
		}
		searchString += "L"
		return bbt.Root.Left.Search(data)
	}

	fmt.Printf("Path: %s\n", searchString)
	searchString = ""
	return bbt.Root
}

// ----------------------------------------------------------------------------------------------------------------

//Min reports the node with smallest data-value in a binary tree
func (bt *BST) Min() int {
	if bt.Root.Left == nil {
		return bt.Root.Data
	}
	return bt.Root.Left.Min()
}

func (bbt *BBST) Min() int {
	if bbt.Root.Left == nil {
		return bbt.Root.Data
	}
	return bbt.Root.Left.Min()
}

//Max reports the node with largest data-value in a binary tree
func (bt *BST) Max() int {
	if bt.Root.Right == nil {
		return bt.Root.Data
	}
	return bt.Root.Right.Max()
}

func (bbt *BBST) Max() int {
	if bbt.Root.Right == nil {
		return bbt.Root.Data
	}
	return bbt.Root.Right.Max()
}

//Insert adds [node] to a binary tree
func (bt *BST) Insert(node *Node) {
	if bt.Root == nil {
		bt.Root = node
	} else {
		if node.Data < bt.Root.Data {
			if bt.Root.Left != nil {
				bt.Root.Left.Insert(node)
			} else {
				bt.Root.Left = &BST{Root: node}
			}
		} else if node.Data > bt.Root.Data {
			if bt.Root.Right != nil {
				bt.Root.Right.Insert(node)
			} else {
				bt.Root.Right = &BST{Root: node}
			}
		}
	}
}

// ----------------------------------------------------------------------------------------------------------------

//! needs work
func (bbt *BBST) Insert(node *BNode) {
	if bbt.Root == nil {
		bbt.Root = node
	} else {
		if node.Data < bbt.Root.Data {
			if bbt.Root.Left != nil {
				bbt.Root.Left.Insert(node)
				if bbt.Root.Left.Root.Data > node.Data {
					_swap(node, bbt.Root.Left.Root)
				}
			} else {
				bbt.Root.Left = &BBST{Root: node}
			}
		} else if node.Data > bbt.Root.Data {
			if bbt.Root.Right != nil {
				bbt.Root.Right.Insert(node)
				if bbt.Root.Right.Root.Data < node.Data {
					_swap(node, bbt.Root.Right.Root)
				}
			} else {
				bbt.Root.Right = &BBST{Root: node}
			}
		}
	}
}

func _swap(node1 *BNode, node2 *BNode) {
	node1.Data, node2.Data = node2.Data, node1.Data
}

// ----------------------------------------------------------------------------------------------------------------

//PreOrder traverses the tree pre-order, and prints its path
func (bt *BST) PreOrder() {

	fmt.Printf("%p -> ", bt.Root.Data)

	if bt.Root.Left != nil {
		bt.Root.Left.PreOrder()
	}
	if bt.Root.Right != nil {
		bt.Root.Right.PreOrder()
	}

}

//InOrder traverses the tree in-order, and reports its path in form of an array
func (bt *BST) InOrder() (res []Node) {
	if bt.Root.Left != nil {
		bt.Root.Left.InOrder()
	}

	fmt.Printf("%p -> ", bt.Root.Data)

	if bt.Root.Right != nil {
		bt.Root.Right.InOrder()
	}

	return
}

//PostOrder traverses the tree post-order, and prints its path
func (bt *BST) PostOrder() (res []Node) {

	if bt.Root.Left != nil {
		bt.Root.Left.PostOrder()
	}
	if bt.Root.Right != nil {
		bt.Root.Right.PostOrder()
	}

	fmt.Printf("%p -> ", bt.Root.Data)
}

//PrintTree prints out the binary tree
func (bt *BST) PrintTree() {

}

func (bt *BST) SortTree() {

}
