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
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	searchField := os.Args[2]
	searchValue := os.Args[3]

	csvReader := csv.NewReader(f)
	rows, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	people := make([]map[string]string, 0)
	headers := make([]string, 0)
	for idx, row := range rows {
		if idx == 0 {
			headers = row
		}
		person := make(map[string]string)
		for jdx, cell := range row {
			person[headers[jdx]] = cell
		}
		people = append(people, person)
	}

	for _, person := range people {
		if person[searchField] == searchValue {
			fmt.Printf("%s,%s,%s,%s\n", person["first"], person["last"], person["city"], person["state"])
		}
	}
}
