package linkedList

import "fmt"

func LLlist_test() {
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
