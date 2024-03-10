package queue

type Node struct {
	next    *Node
	element string
}

func NewNode(e string) *Node {
	return &Node{nil, e}
}

func (n *Node) GetElement() any {
	return n.element
}

func (n *Node) GetNext() *Node {
	return n.next
}
