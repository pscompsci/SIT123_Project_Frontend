package main

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

// DataRow repesents a row of data from a single measurement
type DataRow struct {
	RowID       int
	Stamp       int
	TimeStamp   time.Time
	Temperature float64
	Humidity    float64
	Moisture1   float64
	Moisture2   float64
	Light       float64
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
	row.RowID = convertInt(data[0])
	row.Stamp = convertInt(data[1])
	row.TimeStamp = convertTimeStamp(data[2])
	row.Temperature = convertFloat64(data[3])
	row.Humidity = convertFloat64(data[4])
	row.Moisture1 = convertFloat64(data[5])
	row.Moisture2 = convertFloat64(data[6])
	row.Light = convertFloat64(data[7])
	return &row, nil
}
