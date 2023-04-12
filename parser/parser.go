package parser

// Parser is a common interface for parsing different types of files.
//
// e.g csv, xml, xls
type Parser interface {
	Parse(fileName string) (records [][]string, err error)
}
