package seeds

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func getParamSet(colCount int, rowCount int) string {
	var tmpParamSet []string
	for i := 0; i < colCount; i++ {
		tmpParamSet = append(tmpParamSet, "?")
	}
	paramUnit := fmt.Sprintf("(%s)", strings.Join(tmpParamSet, ", "))

	var listParamSet []string
	for i := 0; i < rowCount; i++ {
		listParamSet = append(listParamSet, paramUnit)
	}
	paramSet := strings.Join(listParamSet, ", ")

	return paramSet
}

func readSeedCSV(table string) [][]string {
	filename := fmt.Sprintf("seeds/%[1]s/%[1]s.csv", table)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records[1:]
}
