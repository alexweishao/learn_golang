package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("working/GolangTraing/27_code-in-process/32_package-encoding-csv/state_table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	fmt.Println(csvReader)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(record)
	}
}
