package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// Check if there is an error and exit the program
func check_error(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

type car_per_manufacturers struct {
	manufacturer string
	amount       int
}

func (cpm car_per_manufacturers) String() string {
	return fmt.Sprintf("%s: %d\n", cpm.manufacturer, cpm.amount)
}

func (cpm *car_per_manufacturers) dump_to_file(file *os.File) {
	_, err := io.Copy(file, strings.NewReader(cpm.String()))
	check_error(err)
}

var reader_chan = make(chan []string)

// Read the csv file and send each line to the reader_chan channel
func csv_reader() {
	file, err := os.Open("Electric_Vehicle_Population_Data.csv")
	check_error(err)
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 0
	reader.TrimLeadingSpace = true
	reader.LazyQuotes = true
	for {
		records, err := reader.Read()
		// on end of file close the channel
		if err == io.EOF {
			close(reader_chan)
			return
		}
		check_error(err)
		reader_chan <- records
	}
}

func main() {
	// Start the csv reader in a goroutine
	go csv_reader()
	// Store the amount of cars per manufacturer in a map
	cars := make(map[string]int)
	for records := range reader_chan {
		cars[records[6]]++
	}
	// Transform the map into a slice of car_per_manufacturer
	cpm := make([]car_per_manufacturers, 0)
	for k, v := range cars {
		cpm = append(cpm, car_per_manufacturers{k, v})
	}
	// Sort the slice by amount of cars
	sort.Slice(cpm, func(i, j int) bool {
		return cpm[i].amount > cpm[j].amount
	})
	// Save the result to a file and print it to the console
	file, err := os.Create("manufacturer.txt")
	check_error(err)
	for _, v := range cpm {
		v.dump_to_file(file)
		fmt.Println(v.manufacturer, ": ", v.amount)
	}
}
