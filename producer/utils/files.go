package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)

type salesCsv struct {
	Product       string  `json:"product"`
	Brand         string  `json:"brand"`
	OriginalPrice float32 `json:"originalPrice"`
	DiscountRate  float32 `json:"discountRate"`
}

func ReadFromCSV(path string) salesCsv {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all the CSV data into a slice of Data struct
	var dataSlice []salesCsv
	for {
		record, err := reader.Read()
		if err != nil {
			break // Break the loop when we reach the end of the file
		}

		data := salesCsv{
			Product:       record[0],
			Brand:         record[1],
			OriginalPrice: StringToFloat(record[2]),
			DiscountRate:  StringToFloat(record[3]),
		}
		dataSlice = append(dataSlice, data)
	}

	// Get a random row from the dataSlice
	randomIndex := rand.Intn(len(dataSlice))
	return dataSlice[randomIndex]
}

func ReadRandomLineFromTxt(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// Create a new scanner to read from the file
	scanner := bufio.NewScanner(file)

	// Read the number of lines in the file
	var totalLines int
	for scanner.Scan() {
		totalLines++
	}

	// Rewind the file pointer back to the beginning
	file.Seek(0, 0)

	// Generate a random number to select a line
	randomLineNumber := rand.Intn(totalLines) + 1

	// Read the lines until reaching the selected random line
	lineNumber := 1
	var randomValue string
	for scanner.Scan() {
		if lineNumber == randomLineNumber {
			// Print the random line
			randomValue = scanner.Text()
		}
		lineNumber++
	}
	return randomValue
}
