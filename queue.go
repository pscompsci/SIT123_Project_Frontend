package main

import (
	"container/list"
)

// FloatDataQueue provides a first-in first-out data structure
// to hold a required number of data readings to display
// in the browser. Data points are float64 values
type FloatDataQueue struct {
	Queue   *list.List
	MaxSize int
}

// NewFloatDataQueue returns a reference to a new FloatDataQueue to hold
// data readings
func NewFloatDataQueue(size int) *FloatDataQueue {
	queue := FloatDataQueue{}
	queue.MaxSize = size
	queue.Queue = list.New()
	return &queue
}

// Enqueue adds a datapoint to the FloatDataQueue
func (q *FloatDataQueue) Enqueue(data float64) float64 {
	if q.Queue.Len() == q.MaxSize {
		q.DeQueue()
	}
	q.Queue.PushBack(data)
	return data
}

// DeQueue removes a datapoint from the front of a FloatDataQueue
func (q *FloatDataQueue) DeQueue() float64 {
	e := q.Queue.Front()
	if q.Queue.Len() > 0 {
		q.Queue.Remove(q.Queue.Front())
	}
	return e.Value.(float64)
}

// AsSlice returns the queue elements as a slice
func (q FloatDataQueue) AsSlice() []float64 {
	slice := []float64{}
	for e := q.Queue.Front(); e != nil; e = e.Next() {
		slice = append(slice, e.Value.(float64))
	}
	return slice
}
