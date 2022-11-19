package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteFile(data [][]string) {
	fileName := "data.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("There is error %v", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.WriteAll(data)
}
