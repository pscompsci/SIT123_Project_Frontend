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

func pollForData(url string) {
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
	queue.Enqueue(17.6)

	fmt.Println("Hello World")
}
