package processor

import (
	"bytes"
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"

	"github.com/NickGowdy/fileprocessor/parser"
)

func TestParsingCsv(t *testing.T) {
	expected := "{\"120\":[{\"client\":\"B\",\"quantity\":20},{\"client\":\"D\",\"quantity\":10}],\"100\":[{\"client\":\"C\",\"quantity\":30}],\"80\":[{\"client\":\"A\",\"quantity\":20}]}"

	records := [][]string{
		{"A", "Red", "80", "20"},
		{"B", "Red", "120", "20"},
		{"C", "Red", "100", "30"},
		{"D", "Red", "120", "10"},
	}

	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	w.WriteAll(records)

	dir := os.TempDir()

	fullDir := filepath.Join(dir, "orders.txt")
	os.WriteFile(fullDir, b.Bytes(), 0644)

	orderProcessor := NewOrderProcessor("orders.txt", dir, parser.NewCsvParser())
	bytes, err := orderProcessor.process()

	if err != nil {
		t.Error(err)
	}

	actual := string(bytes)
	if expected != actual {
		t.Errorf("JSON returned should be: %s, but was: %s", expected, actual)
	}
}

func TestParsingCsvDirError(t *testing.T) {
	dir := os.TempDir()
	fullDir := filepath.Join(dir, "this does not exit")
	os.WriteFile(fullDir, []byte{}, 0644)

	orderProcessor := NewOrderProcessor("this does not exit.txt", dir, parser.NewCsvParser())
	_, err := orderProcessor.process()

	if err == nil {
		t.Error(err)
	}
}
