// dataLoader.go

package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"os"
)

func LoadData(fileName string) dataframe.DataFrame {
	// Open the CSV file
	file, err := os.Open(fileName)
	// if the error isnt nil, then print it
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Read the file
	df := dataframe.ReadCSV(file)

	// Close the file
	file.Close()

	// return the dataframe
	return df
}
