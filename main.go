package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Person struct {
	first    string
	last     string
	city     string
	state    string
	coolness int
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s, file_name.csv first_name", os.Args[0])
		os.Exit(2)
	}

	filePath := os.Args[1]
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	firstName := os.Args[2]

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
		if person["first"] == firstName {
			fmt.Println(person)
		}
	}
}
