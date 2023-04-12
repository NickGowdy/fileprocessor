package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NickGowdy/fileprocessor/parser"
	"github.com/NickGowdy/fileprocessor/processor"
)

/*
Go CMD program for processing of files.
This program reads files from ./files directory, parses the data and prints
the data in JSON format.

Currently this program only supports parsing of CSV (comma seperated value)
data.

Usage:

go run main.go <datatype> <filename.ext>

e.g. go run main.go csv orders.txt

. . .
*/
func main() {
	var output []byte
	var err error

	const (
		csv string = "csv"
		xml string = "xml"
		dir string = "./files"
	)

	if len(os.Args[0:]) != 3 {
		log.Fatalf(`program must be supplied with correct args,
			e.g go run main.go csv orders.txt. But was run with: %v`, os.Args[0:])
		return
	}

	fileType := os.Args[0:][1]
	filename := os.Args[0:][2]

	switch fileType {
	case csv:
		orderProcessor := processor.NewOrderProcessor(filename, dir, parser.NewCsvParser())
		output, err = processor.Processor.DoProcessing(orderProcessor)
	case xml:
		log.Printf("%s is not supported in this version.", fileType)
	default:
		log.Printf("%s is not supported in this version.", fileType)
	}

	if err != nil {
		log.Print(err)
		return
	}

	jsonStr := string(output)

	fmt.Print(jsonStr)
}
