package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/apache/arrow/go/v10/parquet"
	"github.com/apache/arrow/go/v10/parquet/file"
)

import "C"

func OpenParquetFile(filename string, opts ...file.ReadOption) (*file.Reader, error) {
	var source parquet.ReaderAtSeeker

	var err error

	source, err = os.Open(filename)
	if err != nil {
		return nil, err
	}

	return file.NewParquetReader(source, opts...)
}

func WriteFileInfoToBuf(filename string, buf *bytes.Buffer) {
	r, err := OpenParquetFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// log.Printf("Rows: %d", r.NumRows())

	buf.WriteString(fmt.Sprintf("<strong>Number of Rows:</strong> %d<br>\n", r.NumRows()))

	s := r.MetaData().Schema
	ncols := s.NumColumns()

	buf.WriteString(fmt.Sprintf("<strong>Number of Cols:</strong> %d<br>\n", ncols))

	buf.WriteString("<ul>\n")
	for i := 0; i < ncols; i++ {
		var c = s.Column(i)

		buf.WriteString(fmt.Sprintf("\t<li>Name: %s; PhysTyp: %s</li>\n", c.Name(), c.PhysicalType()))
	}
	buf.WriteString("</ul>\n")
}

//export GetParquetSummary
func GetParquetSummary(cpath *C.char) (code C.int, outData unsafe.Pointer, outLen C.long) {
	path := C.GoString(cpath)

	var buf bytes.Buffer

	WriteFileInfoToBuf(path, &buf)

	return 0, C.CBytes(buf.Bytes()), C.long(buf.Len())
}

func main() {}
