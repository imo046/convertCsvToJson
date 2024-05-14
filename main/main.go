package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Item struct {
	Duration string `json:"Duration, sec"`
	ID       string `json:"ID"`
	Notes    string `json:"Notes"`
	Text1    string `json:"Text_Session1"`
	Text2    string `json:"Text_Session2"`
	Text3    string `json:"Text_Session3"`
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func check(jsonData []byte) {
	var myArr []Item
	err := json.Unmarshal(jsonData, &myArr)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range myArr {
		println(item.ID, item.Text1, item.Text2, item.Text3)
	}
}

func convertToJsonArr(jsonData []byte, output string) {
	//Write JSON data to file
	jsonFile, err := os.Create(output)
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing JSON data to file:", err)
		return
	}

	fmt.Println("Conversion completed. JSON data written to output.json")

}

func run(pathToFile string, output string) {
	// Open the CSV file
	csvFile, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Read the CSV file
	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Convert records to JSON
	var jsonData []map[string]string
	for _, row := range records[1:] {
		data := make(map[string]string)
		for i, col := range row {
			// Assuming the first row contains the header
			header := records[0][i]
			data[header] = standardizeSpaces(col)
		}
		jsonData = append(jsonData, data)
	}

	// Convert JSON data to string
	jsonDataString, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	convertToJsonArr(jsonDataString, output)

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go input output")
		os.Exit(1)
	}
	filePath := os.Args[1]
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {

		}
	}
	outPath := os.Args[2]

	if len(outPath) <= 0 {
		log.Fatal("Output cannot be empty")
	}

	run(filePath, outPath)

}
