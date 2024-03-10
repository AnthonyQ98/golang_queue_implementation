package main

import (
	"fmt"
	"queue/src/queue"
)

func main() {
	// Create a new queue
	myQueue := queue.Queue{}

	// Return queue size as integer
	fmt.Println("Queue size: ", myQueue.Size())

	// Check if queue is empty, returns boolean value.
	fmt.Println("Queue empty? ", myQueue.IsEmpty())

	// Fill up the queue with our nodes.
	myQueue.Offer(queue.NewNode("Anthony"))
	myQueue.Offer(queue.NewNode("Tim"))
	myQueue.Offer(queue.NewNode("Pauline"))

	// View element at the front of the queue, but don't remove it.
	fmt.Println("Queue peek:", myQueue.Peek().GetElement())

	// View the element at the front of the queue, and remove it.
	fmt.Println("Queue poll: ", myQueue.Poll().GetElement())

	// Print all the elements in the queue.
	myQueue.PrintQueue()

	// View element at front of queue, but don't remove it.
	fmt.Println("Queue peek:", myQueue.Peek().GetElement())

	// Remove the remaining two elements from the queue.
	fmt.Println("Queue poll: ", myQueue.Poll().GetElement())
	fmt.Println("Queue poll: ", myQueue.Poll().GetElement())

	// Verifying nil is returned when the queue is empty.
	node := myQueue.Poll()
	if node != nil {
		fmt.Println("Queue poll: ", node.GetElement())
	}

}
