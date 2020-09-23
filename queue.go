package main

import (
	"container/list"
)

// DataQueue provides a first-in first-out data structure
// to hold a required number of data readings to display
// in the browser
type DataQueue struct {
	Queue   *list.List
	MaxSize int
}

// NewDataQueue returns a reference to a new DataQueue to hold
// data readings
func NewDataQueue(size int) *DataQueue {
	queue := DataQueue{}
	queue.MaxSize = size
	queue.Queue = list.New()
	return &queue
}

// Enqueue adds a datapoint to the DataQueue
func (q *DataQueue) Enqueue(data float64) float64 {
	if q.Queue.Len() == q.MaxSize {
		q.DeQueue()
	}
	q.Queue.PushBack(data)
	return data
}

// DeQueue removes a datapoint from the front of a DataQueue
func (q *DataQueue) DeQueue() float64 {
	e := q.Queue.Front()
	if q.Queue.Len() > 0 {
		q.Queue.Remove(q.Queue.Front())
	}
	return e.Value.(float64)
}
