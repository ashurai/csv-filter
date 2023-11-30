package main

import (
	"fmt"
	"os"

	"github.com/ashurai/csv-filter/pkg/filter"
	"github.com/ashurai/csv-filter/pkg/reader"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <csv_file_path>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	rows := reader.Reader(file)
	copiesNeeded := filter.Filter(rows)

	fmt.Printf("Total number of copies of application id 374 needed by the company: %d\n", copiesNeeded)
}
