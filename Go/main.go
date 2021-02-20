package main

import (
	ds "dstructs/dstruct"
	"fmt"
)

func main() {

	bt := ds.MakeTree(ds.Node{45, nil, nil},
		ds.Node{20, nil, nil},
		ds.Node{60, nil, nil},
		ds.Node{15, nil, nil},
		ds.Node{40, nil, nil},
		ds.Node{50, nil, nil},
		ds.Node{70, nil, nil},
	)

	fmt.Print("\nInOrder Transversal: ")
	bt.InOrder()
	fmt.Print("\nPostOrder Teansversal: ")
	bt.PostOrder()
	fmt.Print("\nPreOrder Transversal: ")
	bt.PreOrder()

	fmt.Printf("\nMin: %d\n", bt.Min())
	fmt.Printf("Max: %d\n\n", bt.Max())

	fmt.Printf("\nInserting values ... \n")

	bt.Insert(&ds.Node{100, nil, nil})
	bt.Insert(&ds.Node{1, nil, nil})
	bt.Insert(&ds.Node{19, nil, nil})
	bt.Insert(&ds.Node{24, nil, nil})
	bt.Insert(&ds.Node{123, nil, nil})

	fmt.Print("\nInOrder Transversal: ")
	bt.InOrder()
	fmt.Print("\nPostOrder Teansversal: ")
	bt.PostOrder()
	fmt.Print("\nPreOrder Transversal: ")
	bt.PreOrder()

	bt.Search(123)

	fmt.Printf("\nMin: %d\n", bt.Min())
	fmt.Printf("Max: %d\n\n", bt.Max())

	bbt := ds.MakeBTree(ds.BNode{45, nil, nil, nil},
		ds.BNode{20, nil, nil, nil},
		ds.BNode{60, nil, nil, nil},
		ds.BNode{15, nil, nil, nil},
		ds.BNode{40, nil, nil, nil},
		ds.BNode{50, nil, nil, nil},
		ds.BNode{70, nil, nil, nil},
	)

	fmt.Print("\nInOrder Transversal: ")
	bbt.InOrder()
	fmt.Print("\nPostOrder Teansversal: ")
	bbt.PostOrder()
	fmt.Print("\nPreOrder Transversal: ")
	bbt.PreOrder()

	fmt.Printf("\nMin: %d\n", bbt.Min())
	fmt.Printf("Max: %d\n\n", bbt.Max())

	fmt.Printf("\nInserting values ... \n")

	bbt.Insert(&ds.BNode{100, nil, nil, nil})
	bbt.Insert(&ds.BNode{1, nil, nil, nil})
	bbt.Insert(&ds.BNode{19, nil, nil, nil})
	bbt.Insert(&ds.BNode{24, nil, nil, nil})
	bbt.Insert(&ds.BNode{123, nil, nil, nil})

	fmt.Print("\nInOrder Transversal: ")
	bbt.InOrder()
	fmt.Print("\nPostOrder Teansversal: ")
	bbt.PostOrder()
	fmt.Print("\nPreOrder Transversal: ")
	bbt.PreOrder()

	bbt.Search(123)

	fmt.Printf("\nMin: %d\n", bbt.Min())
	fmt.Printf("Max: %d\n\n", bbt.Max())

}
