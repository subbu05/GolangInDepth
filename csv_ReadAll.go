package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	//String in CSV format
	csvString := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"`
	// Read CSV context stored in csvString
	csv := csv.NewReader(strings.NewReader(csvString))
	// Read all rows into records
	records, err := csv.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// Print the records
		for k, v := range records {
			fmt.Println(k, v)
		}
	}
}
