package codeLink

func main() {
	node1 := new(Node)
	node2 := new(Node)
	node3 := new(Node)
	node4 := new(Node)
	node5 := new(Node)

	node1.data = 1
	node1.pNext = node2
	node2.data = 1
	node2.pNext = node2
	node3.data = 1
	node4.pNext = node2
	node5.data = 1

}
