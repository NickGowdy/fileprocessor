package processor

import (
	"fmt"

	"github.com/NickGowdy/fileprocessor/order"
	"github.com/NickGowdy/fileprocessor/parser"
)

type OrderProcessor struct {
	fileName string
	fileDir  string
	Parser   parser.Parser
}

// Create new order processor using file name, file directory and required data format parser.
func NewOrderProcessor(fileName string, fileDir string, parser parser.Parser) OrderProcessor {
	return OrderProcessor{
		fileName: fileName,
		fileDir:  fileDir,
		Parser:   parser,
	}
}

// Processes records and returns []bytes to be serialized.
func (op OrderProcessor) DoProcessing() (bytes []byte, err error) {
	return op.process()
}

func (op OrderProcessor) process() (bytes []byte, err error) {
	fullDir := fmt.Sprintf("%s/%s", op.fileDir, op.fileName)
	records, err := op.Parser.Parse(fullDir)

	if err != nil {
		return nil, err
	}

	order := order.NewOrder(records)
	order.Fill(records)
	j, err := order.MarshalJSON()

	if err != nil {
		return nil, err
	}

	return j, err
}
