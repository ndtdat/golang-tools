package main

import (
	"fmt"
	"github.com/tiendat3550/golang-tools/csv_tool"
	"github.com/tiendat3550/golang-tools/eip712"
)

func main() {
	eip712.CheckVerification()
	records := csv_tool.ReadCsvFile("./csv_tool/csv_files/dssv_1.csv")
	for i, _ := range records {
		for k, l := range records[i] {
			fmt.Println(k, l)
		}
	}
}
