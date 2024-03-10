package main

import (
	"fmt"
	"queue/src/queue"
)

func main() {
	var myQueue queue.Queue
	fmt.Println("Queue size: ", myQueue.Size())
	fmt.Println("Queue empty? ", myQueue.IsEmpty())

	n1 := queue.NewNode("Anthony")
	myQueue.Offer(n1)
	myQueue.Offer(queue.NewNode("Tim"))
	myQueue.Offer(queue.NewNode("Pauline"))
	fmt.Println("Queue peek:", myQueue.Peek().GetElement())
	fmt.Println("Queue poll: ", myQueue.Poll().GetElement())
	fmt.Println(myQueue)
	fmt.Println("Queue peek:", myQueue.Peek().GetElement())

}
