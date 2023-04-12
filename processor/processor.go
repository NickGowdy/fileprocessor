package processor

// Processor is a common interface for reading data from a file and constructing
// data to be returned as an array of bytes.
type Processor interface {
	DoProcessing() (bytes []byte, err error)
}
