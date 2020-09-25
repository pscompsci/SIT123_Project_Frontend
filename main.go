package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func saveData(data [][]string, filename string) error {
	file := filename + ".csv"
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("failed to create or open %s\n", file)
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	for _, d := range data {
		w.Write(d)
	}
	w.Flush()
	return nil
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
		fmt.Println(d.RowID)
	}

	http.HandleFunc("/", handler(queue))

	log.Println("Listening on :3000...")

	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
