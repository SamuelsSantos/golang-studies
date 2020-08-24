package main

import (
	"testing"
)

func TestQueue_Empty(t *testing.T) {
	tests := []struct {
		name string
		q    Queue
		want bool
	}{
		{name: "Check queue is empty", q: Queue{}, want: true},
		{name: "Check queue is empty", q: Queue{1, 2}, want: false},
	}
	for _, tt := range tests {

		var x Queue
		x.Empty()

		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Empty(); got != tt.want {
				t.Errorf("Queue.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	
	queue := Queue{1,2,3}
	expected := 2
	firstElement := 2

	queue.Dequeue();

	if len(queue) != expected || queue[0] != firstElement {
		t.Errorf("Queue.Dequeue() = %v, want %v or %v, want %v", len(queue), expected, queue[0], firstElement)
	}
}


func TestQueue_Enqueue(t *testing.T) {
	
	queue := Queue{1,2,3}
	expected := 4
	lastElement := 4

	queue.Enqueue(4);

	if len(queue) != expected || queue[len(queue)-1] != lastElement {
		t.Errorf("Queue.Dequeue() = %v, want %v or %v, want %v", len(queue), expected, queue[0], lastElement)
	}
}