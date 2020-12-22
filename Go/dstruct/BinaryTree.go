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

func (bbt *BBST) Insert(node *BNode) {
	if bbt.Root == nil {
		bbt.Root = node
	} else {
		if node.Data < bbt.Root.Data {
			if bbt.Root.Left != nil {
				bbt.Root.Left.Insert(node)
			} else {
				bbt.Root.Left = &BBST{Root: node}
			}
		} else if node.Data > bbt.Root.Data {
			if bbt.Root.Right != nil {
				bbt.Root.Right.Insert(node)
			} else {
				bbt.Root.Right = &BBST{Root: node}
			}
		}

	}
}

//PreOrder traverses the tree pre-order, and prints its path
func (bt *BST) PreOrder() {
	bt._preOrder()
	fmt.Println("")
}

func (bt *BST) _preOrder() {

	fmt.Printf("%d -> ", bt.Root.Data)

	if bt.Root.Left != nil {
		bt.Root.Left._preOrder()
	}
	if bt.Root.Right != nil {
		bt.Root.Right._preOrder()
	}

}

func (bbt *BBST) PreOrder() {
	bbt._preOrder()
	fmt.Println("")
}

func (bbt *BBST) _preOrder() {
	fmt.Printf("%d -> ", bbt.Root.Data)

	if bbt.Root.Left != nil {
		bbt.Root.Left._preOrder()
	}
	if bbt.Root.Right != nil {
		bbt.Root.Right._preOrder()
	}
}

//InOrder traverses the tree in-order, and prints its path
func (bt *BST) InOrder() {
	bt._inOrder()
	fmt.Println("")
}

func (bt *BST) _inOrder() {
	if bt.Root.Left != nil {
		bt.Root.Left._inOrder()
	}

	fmt.Printf("%d -> ", bt.Root.Data)

	if bt.Root.Right != nil {
		bt.Root.Right._inOrder()
	}

}

func (bbt *BBST) InOrder() {
	bbt._inOrder()
	fmt.Println("")
}

func (bbt *BBST) _inOrder() {
	if bbt.Root.Left != nil {
		bbt.Root.Left._inOrder()
	}
	fmt.Printf("%d -> ", bbt.Root.Data)

	if bbt.Root.Right != nil {
		bbt.Root.Right._inOrder()
	}
}

//PostOrder traverses the tree post-order, and prints its path
func (bt *BST) PostOrder() {
	bt._postOrder()
	fmt.Println("")
}

func (bt *BST) _postOrder() {
	if bt.Root.Left != nil {
		bt.Root.Left._postOrder()
	}
	if bt.Root.Right != nil {
		bt.Root.Right._postOrder()
	}

	fmt.Printf("%d -> ", bt.Root.Data)
}

func (bbt *BBST) PostOrder() {
	bbt._postOrder()
	fmt.Println("")
}

func (bbt *BBST) _postOrder() {
	if bbt.Root.Left != nil {
		bbt.Root.Left._postOrder()
	}
	if bbt.Root.Right != nil {
		bbt.Root.Right._postOrder()
	}

	fmt.Printf("%d -> ", bbt.Root.Data)
}

//PrintTree prints out the binary tree
func (bt *BST) PrintTree() {

}

func (bt *BST) SortTree() {

}
