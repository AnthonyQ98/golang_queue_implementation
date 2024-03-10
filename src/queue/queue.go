package queue

type Queue struct {
	size  int
	first *Node
	queue []*Node
}

func (q *Queue) Peek() *Node {
	if q.IsEmpty() {
		return nil
	}
	return q.first
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Offer(n *Node) {
	if q.IsEmpty() {
		q.first = n
	} else {
		q.queue = append(q.queue, n)
	}
	q.size++
}

func (q *Queue) Poll() *Node {
	if q.IsEmpty() {
		return nil
	}
	var temp = *q.first
	q.first = temp.next
	q.size--
	return &temp
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

type Node struct {
	next    *Node
	element string
}

func (n *Node) GetElement() any {
	return n.element
}

func NewNode(e string) *Node {
	return &Node{nil, e}
}
