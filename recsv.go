package main

import (
	"cmp"
	"encoding/csv"
	"log"
	"os"
	"slices"
)

func main() {
	reader := csv.NewReader(os.Stdin)

	// ReadAll reads all the records from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	records, err := reader.ReadAll()

	// Checks for the error
	if err != nil {
		log.Println("Error reading records")
	}

	// Sort input CSV lines
	slices.SortFunc(records, func(a, b []string) int {
		eq := 0
		for i := 0; i < len(a) && eq == 0; i++ {
			eq = cmp.Compare(a[i], b[i])
		}
		return eq
	})

	w := csv.NewWriter(os.Stdout)
	// Loop to iterate through
	// and write out each of the string slice
	for _, eachrecord := range records {
		err = w.Write(eachrecord)
		if err != nil {
			log.Println("Error writing records")
		}
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		log.Fatalln(err, "Writing report output to CSV")
	}
}
