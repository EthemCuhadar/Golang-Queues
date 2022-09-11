package main

import (
	"fmt"
	"sync"
)

// Any - The type of item in queue
type Any interface{}

// Queue - Create queue item
type Queue struct {

	// data - Stores items in queue.
	// Item type is Any(interface).
	data []Any

	// mux - Locks and unlocks for
	// concurrent applications in queue.
	mux sync.RWMutex
}

// Len - Returns size of items in queue.
func (q *Queue) Len() int {

	// Lock and unlock the mutex.
	q.mux.Lock()
	defer q.mux.Unlock()
	return len(q.data)
}

// IsEmpty - Checks is the queue empty or not.
func (q *Queue) IsEmpty() bool {

	// Lock and unlock the mutex.
	q.mux.Lock()
	defer q.mux.Unlock()
	return len(q.data) == 0
}

// Enqueue - Adds an element at the end of queue.
func (q *Queue) Enqueue(a Any) {

	// Lock the mutex.
	q.mux.Lock()
	// Apply append operation.
	q.data = append(q.data, a)
	// Unlock the mutex.
	q.mux.Unlock()
}

// Dequeue - Discards the last element of queue.
func (q *Queue) Dequeue() Any {

	// Lock and Unlock the mutex.
	q.mux.Lock()
	defer q.mux.Unlock()
	// Check if queue is empty.
	if len(q.data) == 0 {
		return nil
	}
	// Apply discard operation.
	l := len(q.data)
	tmp := q.data[0]
	q.data = q.data[1:l]
	return tmp
}

// First - Returns the last element in queue.
func (q *Queue) First() Any {
	// Lock the mutex.
	q.mux.Lock()
	// Unlock the mutex
	defer q.mux.Unlock()

	// Check if queue is empty.
	if len(q.data) == 0 {
		return nil
	}
	return q.data[0]
}

// Show - Displays the elements that are stored.
func (q *Queue) Show() Any {
	// Lock the mutex.
	q.mux.Lock()
	// Unlock the mutex
	defer q.mux.Unlock()

	// Check if queue is queue.
	if len(q.data) == 0 {
		return nil
	}
	return q.data
}

func main() {

	// Program starting...
	fmt.Println("Program starting...")

	// Create a new Queue.
	Q := new(Queue)
	fmt.Println("empty? ", Q.IsEmpty())

	// Enqueue items
	Q.Enqueue("cat")
	Q.Enqueue(1999)
	Q.Enqueue("car")
	Q.Enqueue(11.111)
	Q.Enqueue(true)

	// Print
	fmt.Println(Q.Show())
	fmt.Println("len(Q) ", Q.Len())
	fmt.Println("dequeue ", Q.Dequeue())
	fmt.Println("dequeue ", Q.Dequeue())
	fmt.Println(Q.Show())
	fmt.Println("len(Q) ", Q.Len())
	fmt.Println("first ", Q.First())
	fmt.Println("empty? ", Q.IsEmpty())
}
