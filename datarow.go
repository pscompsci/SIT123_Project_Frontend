package main

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

// DataRow repesents a row of data from a single measurement
type DataRow struct {
	rowID       int
	stamp       int
	timeStamp   time.Time
	temperature float64
	humidity    float64
	moisture1   float64
	moisture2   float64
	light       float64
}

func convertTimeStamp(t string) time.Time {
	format := "2006/02/01 15:04:05"
	// confident in the format of the string, so discarding the error
	date, _ := time.Parse(format, t)
	return date
}

func convertFloat64(d string) float64 {
	// confident in the format of the string, so discarding the error
	data, _ := strconv.ParseFloat(d, 64)
	return data
}

func convertInt(d string) int {
	// confident in the format of the string, so discarding the error
	data, _ := strconv.Atoi(d)
	return data
}

func extractData(data []string) (*DataRow, error) {
	row := DataRow{}
	if len(data) != reflect.TypeOf(DataRow{}).NumField() {
		return &row, errors.New("data has incorrect structure")
	}
	row.rowID = convertInt(data[0])
	row.stamp = convertInt(data[1])
	row.timeStamp = convertTimeStamp(data[2])
	row.temperature = convertFloat64(data[3])
	row.humidity = convertFloat64(data[4])
	row.moisture1 = convertFloat64(data[5])
	row.moisture2 = convertFloat64(data[6])
	row.light = convertFloat64(data[7])
	return &row, nil
}
