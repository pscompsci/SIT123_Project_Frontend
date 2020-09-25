package main

import (
	"container/list"
)

// DataQueue provides a first-in first-out data structure
// to hold a required number of data readings to display
// in the browser. Data points are float64 values
type DataQueue struct {
	Queue   *list.List
	MaxSize int
}

// NewDataQueue returns a reference to a new DataQueue to hold
// data readings
func NewDataQueue(size int) DataQueue {
	queue := DataQueue{}
	queue.MaxSize = size
	queue.Queue = list.New()
	return queue
}

// Enqueue adds an element to the end of the DataQueue
func (q *DataQueue) Enqueue(data []string) (*DataRow, error) {
	if q.Queue.Len() == q.MaxSize {
		q.DeQueue()
	}
	element, err := extractData(data)
	if err != nil {
		return &DataRow{}, err
	}
	q.Queue.PushBack(element)
	return element, nil
}

// DeQueue removes the first element from the front of a DataQueue
func (q *DataQueue) DeQueue() DataRow {
	e := q.Queue.Front()
	if q.Queue.Len() > 0 {
		q.Queue.Remove(q.Queue.Front())
	}
	return e.Value.(DataRow)
}
