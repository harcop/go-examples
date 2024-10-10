package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the CSV file
	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Create a CSV reader
	reader := csv.NewReader(csvFile)

	// Read the header (first line)
	header, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV header:", err)
		return
	}

	// Prepare a slice to hold the rows as maps
	var data []map[string]string

	// Read each record (row)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break // End of file
		}
		if err != nil {
			fmt.Println("Error reading CSV row:", err)
			return
		}

		// Create a map for each row, using the header as keys
		row := make(map[string]string)
		for i, value := range record {
			row[header[i]] = value
		}

		// Append the row map to the data slice
		data = append(data, row)
	}

	// Convert the data to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Print or save the JSON data
	fmt.Println(string(jsonData))

	// Optionally write the JSON data to a file
	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Println("CSV data successfully converted to JSON and saved to output.json")
}
