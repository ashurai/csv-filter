package pkg

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Installation struct represents an installation record
type Installation struct {
	ComputerID    int
	UserID        int
	ApplicationID int
	ComputerType  string
	Comment       string
}

// Reader function to read CSV file line by line and ready data to evaluation
func Reader(file *os.File) []Installation {
	reader := csv.NewReader(file)
	var rows []Installation

	// Parse CSV
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			os.Exit(1)
		}
		fmt.Println("1:", record[0], "2:", record[1], "3:", record[2])
		computerID, _ := strconv.Atoi(record[0])

		userID, _ := strconv.Atoi(record[1])
		fmt.Println("computerID", computerID, "userID", userID)
		applicationID, _ := strconv.Atoi(record[2])

		row := Installation{
			ComputerID:    computerID,
			UserID:        userID,
			ApplicationID: applicationID,
			ComputerType:  record[3],
			Comment:       record[4],
		}

		rows = append(rows, row)
	}
	//fmt.Println("Rowa:", rows)
	return rows
}
