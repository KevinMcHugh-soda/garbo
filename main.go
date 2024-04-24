package main

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

//go:embed cool_people.csv
var peopleBytes []byte

type Person struct {
	first    string
	last     string
	city     string
	state    string
	coolness int
}

func main() {
	r := bytes.NewReader(peopleBytes)
	reader := csv.NewReader(r)
	// throw the header row away
	_, err := reader.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Hello, World!")

	people := make([]Person, 0)
	for i := 0; true; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
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
