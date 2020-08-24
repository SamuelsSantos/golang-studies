package main

// Queue type
type Queue []int

// Empty check queue is empty
func (q Queue) Empty() bool    { return len(q) == 0 }

// Enqueue add queue element
func (q *Queue) Enqueue(v int) { (*q) = append((*q), v) }

// Dequeue remove queue element
func (q *Queue) Dequeue() int {
	v := (*q)[0]
	(*q) = (*q)[1:len(*q)]
	return v
}
