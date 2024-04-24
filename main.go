package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Person struct {
	first    string
	last     string
	city     string
	state    string
	coolness int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s, file_name.csv", os.Args[0])
		os.Exit(2)
	}

	filePath := os.Args[1]
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	people := make([]Person, 0)
	for idx, record := range records {
		if idx == 0 {
			continue
		}
		a, err := strconv.Atoi(record[4])
		if err != nil {
			fmt.Println(err)
			return
		}
		people = append(people, Person{
			first:    record[0],
			last:     record[1],
			city:     record[2],
			state:    record[3],
			coolness: a,
		})
	}

	for _, person := range people {
		if person.first == "kevin" {
			fmt.Println(person)
		}
	}
}
