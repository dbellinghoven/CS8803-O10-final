package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func main() {
	// Open the raw data file
	original, err := os.Open("/go/src/app/fhlmc_sf2018a_loans.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer original.Close()

	// Create a new CSV file to write to
	modified, err := os.OpenFile("/go/src/app/fhlmc_sf2018a_loans.csv", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer modified.Close()

	// Create a CSV writer
	w := csv.NewWriter(modified)

	w.Write([]string{
		"msa_code",
		"percent_minority",
		"tract_income_ratio",
		"borrower_income_ratio",
		"ltv",
		"purpose_of_loan",
		"federal_guarantee",
		"borrower_race",
		"coborrower_race",
		"borrower_gender",
		"coborrower_gender",
		"number_of_units",
		"affordability_category",
	})

	// Read the original file line by line
	scanner := bufio.NewScanner(original)
	for scanner.Scan() {
		row := strings.Fields(scanner.Text()[2:])
		if err = w.Write(row); err != nil {
			log.Fatal(err)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	w.Flush()

	if err = w.Error(); err != nil {
		log.Fatal(err)
	}
}
