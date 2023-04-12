package parser

import (
	"bytes"
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"
)

func TestParsingCsv(t *testing.T) {
	parser := NewCsvParser()

	records := [][]string{
		{"A", "Red", "80", "20"},
	}

	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	w.WriteAll(records)

	dir := os.TempDir()

	fullDir := filepath.Join(dir, "orders.txt")
	os.WriteFile(fullDir, b.Bytes(), 0644)

	data, err := parser.Parse(fullDir)

	if err != nil {
		t.Errorf("error should not be null but is: %v", err)
	}

	if data[0][0] != "A" {
		t.Errorf("expected A, but got: %s", data[0][0])
	}

	if data[0][1] != "Red" {
		t.Errorf("expected Red, but got: %s", data[0][1])
	}

	if data[0][2] != "80" {
		t.Errorf("expected 80, but got: %s", data[0][2])
	}

	if data[0][3] != "20" {
		t.Errorf("expected 20, but got: %s", data[0][3])
	}
}

func TestParsingCsvCantFindDir(t *testing.T) {
	parser := NewCsvParser()

	records := [][]string{
		{"A", "Red", "80", "20"},
	}

	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	w.WriteAll(records)

	_, err := parser.Parse("fullDir")

	if err == nil {
		t.Errorf("error should be null but is: %v", err)
	}
}
