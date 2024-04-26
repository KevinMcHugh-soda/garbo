package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Criterion struct {
	searchField string
	searchValue string
}

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s, file_name.csv first kevin last mchugh", os.Args[0])
		os.Exit(2)
	}

	filePath := os.Args[1]

	criteria := make([]Criterion, 0)
	for idx := 2; idx < len(os.Args); idx += 2 {
		criteria = append(criteria, Criterion{
			searchField: os.Args[idx],
			searchValue: os.Args[idx+1],
		})
	}

	headers, records := loadRecords(filePath)

	for _, record := range records {
		if matchesAllCriteria(record, criteria) {
			str := ""
			// Since go randomizes map access order, iterating over the headers preserves the correct order.
			for _, key := range headers {
				str += record[key] + ","
			}
			fmt.Printf("%s\n", str)
		}
	}
}

func matchesAllCriteria(record map[string]string, criteria []Criterion) bool {
	for _, criterion := range criteria {
		if criterion.searchValue[0] == "-"[0] {
			if "-"+record[criterion.searchField] != criterion.searchValue {
				return false
			}
		} else if record[criterion.searchField] == criterion.searchValue {
			return false
		}
	}
	return true
}

func loadRecords(filePath string) ([]string, []map[string]string) {
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

	return headers, records
}
