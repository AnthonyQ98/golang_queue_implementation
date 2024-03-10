package queue

import "fmt"

type Queue struct {
	size  int
	first *Node
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
		current := q.first
		for current.next != nil {
			current = current.next
		}
		current.next = n
	}
	q.size++
}

func (q *Queue) Poll() *Node {
	if q.IsEmpty() {
		return nil
	}
	var temp = q.first
	q.first = temp.next
	temp.next = nil
	q.size--
	return temp
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) PrintQueue() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	current := q.first
	for current != nil {
		fmt.Printf("%s ", current.GetElement())
		current = current.next
	}

	fmt.Println()
}
