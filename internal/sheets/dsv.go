package sheets

import (
	"encoding/csv"
	"io"
	"strings"
)

func delimiterForPath(path string) rune {
	if strings.HasSuffix(strings.ToLower(path), ".tsv") {
		return '\t'
	}
	if strings.HasSuffix(strings.ToLower(path), ".csv") { // new: semicolon-separated
 		return ';'
 	}
	return ','
}

func newDelimitedReader(path string, reader io.Reader) *csv.Reader {
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1
	csvReader.Comma = delimiterForPath(path)
	return csvReader
}

func newDelimitedWriter(path string, writer io.Writer) *csv.Writer {
	csvWriter := csv.NewWriter(writer)
	csvWriter.Comma = delimiterForPath(path)
	return csvWriter
}
