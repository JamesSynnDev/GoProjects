package linkedList

import "fmt"

type SNode struct {
	Value int
	SNext *SNode
}

func addNode(t *SNode, v int) int {
	if root == nil {
		t = &SNode{v, nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists: ", v)
		return -1
	}

	if t.SNext == nil {
		t.SNext = &SNode{v, nil}
		return -2

	}
	return addNode(t.SNext, v)
}

func traverse(t *SNode) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.SNext

	}
	fmt.Println()
}

var root = new(SNode)

func Llist_Test() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, 1)

	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 0)
	addNode(root, 0)

	traverse(root)
	addNode(root, 100)
	traverse(root)
}
