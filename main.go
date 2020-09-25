package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: 5 * time.Second,
	}
)

func callClient(url string) (io.Reader, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func getData(url string) (data [][]string, err error) {
	r, err := callClient(url)
	if err != nil {
		return data, err
	}
	reader := csv.NewReader(r)
	reader.Comma = ','
	data, err = reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return data, nil
}

func saveData([][]string) {

}

func queueData(data [][]string, queue *DataQueue) {
	for _, d := range data {
		queue.Enqueue(d)
	}
}

func pollForData(url string, queue *DataQueue) {
	for {
		start := time.Now()
		data, err := getData(url)
		if err != nil {
			log.Printf("Error reading data from response: %v\n", err)
		}
		fmt.Println(data)
		elapsed := time.Since(start)
		time.Sleep(10*time.Second - elapsed)
	}
}

func main() {
	queue := NewDataQueue(100)
	queue.Enqueue([]string{"10001", "213456789", "2020/09/17 14:15:10", "46", "17.6", "760", "767", "1018"})
	queue.Enqueue([]string{"20001", "213466789", "2020/09/17 14:15:20", "46", "17.6", "760", "767", "1018"})

	_, err := queue.Enqueue([]string{"20001", "213466789", "2020/09/17 14:15:20"})
	if err != nil {
		fmt.Println("Error correctly captured")
	}

	for e := queue.Queue.Front(); e != nil; e = e.Next() {
		d := e.Value.(*DataRow)
		fmt.Println(d.rowID)
	}
}
