package parser

import (
	"encoding/csv"
	"os"
)

type CsvParser struct {
}

func NewCsvParser() CsvParser {
	return CsvParser{}
}

func (p CsvParser) Parse(fullDir string) ([][]string, error) {

	file, err := os.Open(fullDir)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, err
}
