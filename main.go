package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s, file_name.csv first kevin", os.Args[0])
		os.Exit(2)
	}

	filePath := os.Args[1]

	searchField := os.Args[2]
	searchValue := os.Args[3]

	records := loadRecords(filePath)

	for _, record := range records {
		if searchValue[0] == "-"[0] {
			if "-"+record[searchField] != searchValue {
				fmt.Printf("%s\n", record["first"])
			}
		} else if record[searchField] == searchValue {
			fmt.Printf("%s\n", record["first"])
		}
	}
}

func loadRecords(filePath string) []map[string]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	rows, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	records := make([]map[string]string, 0)
	headers := make([]string, 0)
	for idx, row := range rows {
		if idx == 0 {
			headers = row
			continue
		}
		record := make(map[string]string)
		for jdx, cell := range row {
			record[headers[jdx]] = cell
		}
		records = append(records, record)
	}

	return records
}
